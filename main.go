package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/net/proxy"
	"gopkg.in/yaml.v3"
)

// ScanResult represents the result of scanning a single URL
type ScanResult struct {
	URL        string    `json:"url"`
	Status     string    `json:"status"`
	StatusCode int       `json:"status_code"`
	Error      string    `json:"error,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
	Content    string    `json:"content,omitempty"`
}

// ScanReport contains overall scan statistics
type ScanReport struct {
	TotalURLs  int          `json:"total_urls"`
	Successful int          `json:"successful"`
	Failed     int          `json:"failed"`
	StartTime  time.Time    `json:"start_time"`
	EndTime    time.Time    `json:"end_time"`
	Results    []ScanResult `json:"results"`
}

// Target represents a single target entry
type Target struct {
	URL          string `yaml:"url"`
	Type         string `yaml:"type,omitempty"`
	Name         string `yaml:"name,omitempty"`
	MockResponse string `yaml:"mock_response,omitempty"`
}

// YAMLConfig represents the YAML file structure
type YAMLConfig struct {
	Targets []Target `yaml:"targets"`
}

// readTargets reads the targets from a YAML or TXT file
func readTargets(filePath string) ([]Target, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var targets []Target

	// Try to parse as YAML first
	var config YAMLConfig
	err = yaml.Unmarshal(data, &config)
	if err == nil && len(config.Targets) > 0 {
		// Successfully parsed as YAML
		for _, target := range config.Targets {
			if target.URL != "" {
				targets = append(targets, target)
			}
		}
		return targets, nil
	}

	// Fall back to line-by-line parsing for TXT format
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		targets = append(targets, Target{URL: line})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return targets, nil
}

// createTorClient creates an HTTP client configured to use Tor proxy
func createTorClient() (*http.Client, error) {
	// Try connecting to Tor SOCKS5 proxy (9150 for Tor Browser, 9050 for standard Tor)
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, proxy.Direct)
	if err != nil {
		// Try alternative port if first fails
		fmt.Println("[WARN] Could not connect to port 9150, trying 9050...")
		dialer, err = proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("failed to create SOCKS5 dialer: %w", err)
		}
	}

	// Create custom transport using the Tor dialer with DialContext support
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
		// Additional security settings
		DisableKeepAlives:     true,
		DisableCompression:    true,
		MaxIdleConnsPerHost:   1,
		TLSHandshakeTimeout:   20 * time.Second,
		ResponseHeaderTimeout: 20 * time.Second,
	}

	// Create HTTP client with custom transport
	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	return client, nil
}

// scanURL fetches content from a single .onion URL
func scanURL(client *http.Client, target Target) ScanResult {
	url := target.URL
	result := ScanResult{
		URL:       url,
		Timestamp: time.Now(),
	}

	// Check if mock response is provided
	if target.MockResponse != "" {
		result.StatusCode = 200
		result.Status = "SUCCESS"
		result.Content = target.MockResponse
		fmt.Printf("[SUCCESS] Scanning: %s -> Status: 200 (mocked)\n", url)
		return result
	}

	// Ensure URL has http:// prefix
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	fmt.Printf("[INFO] Scanning: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		result.Status = "ERROR"
		result.Error = fmt.Sprintf("Failed to create request: %v", err)
		fmt.Printf("[ERR] %s -> %s\n", url, result.Error)
		return result
	}

	// Set a reasonable User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:91.0) Gecko/20100101 Firefox/91.0")

	resp, err := client.Do(req)
	if err != nil {
		result.Status = "FAILED"
		result.Error = fmt.Sprintf("Request failed: %v", err)
		result.StatusCode = 0
		fmt.Printf("[ERR] Scanning: %s -> FAILED (%v)\n", url, err)
		return result
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	result.Status = "SUCCESS"

	// Read response body (limit to 1MB to avoid huge files)
	limitedReader := io.LimitReader(resp.Body, 1024*1024)
	body, err := io.ReadAll(limitedReader)
	if err != nil {
		result.Status = "PARTIAL"
		result.Error = fmt.Sprintf("Error reading response: %v", err)
	} else {
		result.Content = string(body)
	}

	fmt.Printf("[SUCCESS] Scanning: %s -> Status: %d\n", url, resp.StatusCode)
	return result
}

// saveScanReport saves the scan results to a JSON file and log file
func saveScanReport(report ScanReport, outputDir string) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Save JSON report
	reportPath := filepath.Join(outputDir, "scan_report.json")
	reportJSON, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal report: %w", err)
	}

	if err := os.WriteFile(reportPath, reportJSON, 0644); err != nil {
		return fmt.Errorf("failed to write report file: %w", err)
	}
	fmt.Printf("[INFO] Report saved to: %s\n", reportPath)

	// Save log file with summary
	logPath := filepath.Join(outputDir, "scan_report.log")
	logFile, err := os.Create(logPath)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer logFile.Close()

	logContent := fmt.Sprintf(`Tor Scraper Scan Report
=======================
Start Time: %s
End Time: %s
Duration: %v

Statistics:
-----------
Total URLs: %d
Successful: %d
Failed: %d
Success Rate: %.2f%%

Detailed Results:
-----------------
`, report.StartTime.Format(time.RFC3339), report.EndTime.Format(time.RFC3339),
		report.EndTime.Sub(report.StartTime), report.TotalURLs, report.Successful, report.Failed,
		float64(report.Successful)/float64(report.TotalURLs)*100)

	logFile.WriteString(logContent)

	for _, result := range report.Results {
		logLine := fmt.Sprintf("[%s] %s - Status: %s", result.Timestamp.Format("2006-01-02 15:04:05"), result.URL, result.Status)
		if result.StatusCode > 0 {
			logLine += fmt.Sprintf(" (HTTP %d)", result.StatusCode)
		}
		if result.Error != "" {
			logLine += fmt.Sprintf(" - Error: %s", result.Error)
		}
		logLine += "\n"
		logFile.WriteString(logLine)

		// Also save content to individual files for successful scans
		if result.Status == "SUCCESS" && result.Content != "" {
			filename := strings.ReplaceAll(strings.ReplaceAll(result.URL, "http://", ""), "/", "_")
			filename = strings.TrimSuffix(filename, "_")
			contentPath := filepath.Join(outputDir, "content", filename+".html")

			os.MkdirAll(filepath.Dir(contentPath), 0755)
			if err := os.WriteFile(contentPath, []byte(result.Content), 0644); err != nil {
				fmt.Printf("[WARN] Failed to save content for %s: %v\n", result.URL, err)
			}
		}
	}

	fmt.Printf("[INFO] Log file saved to: %s\n", logPath)
	return nil
}

// main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tor-scraper <targets_file> [output_directory]")
		fmt.Println("Example: go run main.go targets.yaml")
		fmt.Println("\nMake sure Tor service is running!")
		os.Exit(1)
	}

	targetsFile := os.Args[1]
	outputDir := "output"
	if len(os.Args) > 2 {
		outputDir = os.Args[2]
	}

	fmt.Println("========================================")
	fmt.Println("      Tor Scraper - .onion Scanner")
	fmt.Println("========================================")
	fmt.Println()

	// Read targets from file
	fmt.Printf("[INFO] Reading targets from: %s\n", targetsFile)
	targets, err := readTargets(targetsFile)
	if err != nil {
		fmt.Printf("[ERR] Failed to read targets: %v\n", err)
		os.Exit(1)
	}

	if len(targets) == 0 {
		fmt.Println("[ERR] No targets found in file")
		os.Exit(1)
	}

	fmt.Printf("[INFO] Found %d targets\n", len(targets))
	fmt.Println()

	// Create Tor-enabled HTTP client
	fmt.Println("[INFO] Connecting to Tor network...")
	client, err := createTorClient()
	if err != nil {
		fmt.Printf("[ERR] Failed to create Tor client: %v\n", err)
		fmt.Println("[INFO] Make sure Tor service is running on port 9050 or 9150")
		os.Exit(1)
	}
	fmt.Println("[SUCCESS] Connected to Tor proxy")
	fmt.Println()

	// Start scanning
	startTime := time.Now()
	report := ScanReport{
		TotalURLs: len(targets),
		StartTime: startTime,
	}

	fmt.Println("[INFO] Starting scan...")
	fmt.Println()

	for _, target := range targets {
		result := scanURL(client, target)
		report.Results = append(report.Results, result)

		if result.Status == "SUCCESS" {
			report.Successful++
		} else {
			report.Failed++
		}

		// Add a small delay between requests to avoid overwhelming the network
		time.Sleep(1 * time.Second)
	}

	endTime := time.Now()
	report.EndTime = endTime

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("           Scan Complete")
	fmt.Println("========================================")
	fmt.Printf("Duration: %v\n", endTime.Sub(startTime))
	fmt.Printf("Successful: %d/%d\n", report.Successful, report.TotalURLs)
	fmt.Println()

	// Save report
	if err := saveScanReport(report, outputDir); err != nil {
		fmt.Printf("[ERR] Failed to save report: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("[SUCCESS] Scan complete!")
}
