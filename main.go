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

// generateHTMLReport generates a detailed HTML report
func generateHTMLReport(report ScanReport) string {
	duration := report.EndTime.Sub(report.StartTime)
	successRate := 0.0
	if report.TotalURLs > 0 {
		successRate = float64(report.Successful) / float64(report.TotalURLs) * 100
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tor Scraper - Scan Report</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            padding: 20px;
            min-height: 100vh;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background: white;
            border-radius: 10px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            overflow: hidden;
        }
        .header {
            background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
            color: white;
            padding: 40px 20px;
            text-align: center;
        }
        .header h1 {
            font-size: 2.5em;
            margin-bottom: 10px;
        }
        .header p {
            font-size: 1.1em;
            opacity: 0.9;
        }
        .content {
            padding: 40px 20px;
        }
        .stats-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 20px;
            margin-bottom: 40px;
        }
        .stat-card {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 20px;
            border-radius: 8px;
            text-align: center;
            box-shadow: 0 5px 15px rgba(0,0,0,0.1);
        }
        .stat-card h3 {
            font-size: 1.2em;
            margin-bottom: 10px;
            opacity: 0.9;
        }
        .stat-card .value {
            font-size: 2.5em;
            font-weight: bold;
        }
        .timeline-info {
            background: #f5f5f5;
            padding: 20px;
            border-radius: 8px;
            margin-bottom: 30px;
            border-left: 4px solid #667eea;
        }
        .timeline-info p {
            margin: 8px 0;
            color: #333;
        }
        .results-table {
            width: 100%%;
            border-collapse: collapse;
            margin-top: 30px;
        }
        .results-table th {
            background: #333;
            color: white;
            padding: 15px;
            text-align: left;
            font-weight: 600;
        }
        .results-table td {
            padding: 12px 15px;
            border-bottom: 1px solid #ddd;
        }
        .results-table tr:hover {
            background: #f9f9f9;
        }
        .status-success {
            color: #27ae60;
            font-weight: 600;
        }
        .status-failed {
            color: #e74c3c;
            font-weight: 600;
        }
        .status-error {
            color: #e67e22;
            font-weight: 600;
        }
        .footer {
            background: #f5f5f5;
            padding: 20px;
            text-align: center;
            color: #666;
            border-top: 1px solid #ddd;
        }
        .progress-bar {
            width: 100%%;
            height: 30px;
            background: #ddd;
            border-radius: 5px;
            overflow: hidden;
            margin-top: 10px;
        }
        .progress-fill {
            height: 100%%;
            background: linear-gradient(90deg, #27ae60 0%, #2ecc71 100%);
            display: flex;
            align-items: center;
            justify-content: center;
            color: white;
            font-weight: bold;
            transition: width 0.3s;
        }
        h2 {
            color: #333;
            margin-top: 30px;
            margin-bottom: 15px;
            border-bottom: 2px solid #667eea;
            padding-bottom: 10px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🎯 Tor Scraper Scan Report</h1>
            <p>Comprehensive Web Scanning Analysis</p>
        </div>

        <div class="content">
            <div class="stats-grid">
                <div class="stat-card">
                    <h3>Total URLs</h3>
                    <div class="value">%d</div>
                </div>
                <div class="stat-card">
                    <h3>Successful</h3>
                    <div class="value" style="color: #2ecc71;">%d</div>
                </div>
                <div class="stat-card">
                    <h3>Failed</h3>
                    <div class="value" style="color: #e74c3c;">%d</div>
                </div>
                <div class="stat-card">
                    <h3>Success Rate</h3>
                    <div class="value">%.1f%%</div>
                </div>
            </div>

            <div class="stat-card" style="grid-column: span 2;">
                <h3>Progress</h3>
                <div class="progress-bar">
                    <div class="progress-fill" style="width: %.1f%%">
                        %.1f%%
                    </div>
                </div>
            </div>

            <div class="timeline-info">
                <h2>📊 Scan Timeline</h2>
                <p><strong>Start Time:</strong> %s</p>
                <p><strong>End Time:</strong> %s</p>
                <p><strong>Duration:</strong> %v</p>
            </div>

            <h2>📋 Detailed Results</h2>
            <table class="results-table">
                <thead>
                    <tr>
                        <th>URL</th>
                        <th>Status</th>
                        <th>HTTP Code</th>
                        <th>Timestamp</th>
                        <th>Details</th>
                    </tr>
                </thead>
                <tbody>`,
		report.TotalURLs, report.Successful, report.Failed, successRate, successRate, successRate,
		report.StartTime.Format("2006-01-02 15:04:05"),
		report.EndTime.Format("2006-01-02 15:04:05"),
		duration)

	// Add result rows
	for _, result := range report.Results {
		statusClass := "status-success"
		if result.Status == "FAILED" || result.Status == "ERROR" {
			statusClass = "status-failed"
		} else if result.Status == "PARTIAL" {
			statusClass = "status-error"
		}

		details := ""
		if result.Error != "" {
			details = result.Error
		}
		if result.Content != "" && result.Status == "SUCCESS" {
			details = fmt.Sprintf("Content length: %d bytes", len(result.Content))
		}

		html += fmt.Sprintf(`
                    <tr>
                        <td><strong>%s</strong></td>
                        <td class="%s">%s</td>
                        <td>%d</td>
                        <td>%s</td>
                        <td>%s</td>
                    </tr>`,
			result.URL, statusClass, result.Status, result.StatusCode,
			result.Timestamp.Format("15:04:05"), details)
	}

	html += `
                </tbody>
            </table>

            <div class="footer">
                <p>Generated: ` + time.Now().Format("2006-01-02 15:04:05") + `</p>
                <p>Tor Scraper v1.0 | Cyber Threat Intelligence Tool</p>
            </div>
        </div>
    </div>
</body>
</html>`

	return html
}

// saveScanReport saves the scan results to multiple formats (JSON, HTML, TXT, CSV)
func saveScanReport(report ScanReport, outputDir string) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// 1. Save JSON report
	reportPath := filepath.Join(outputDir, "scan_report.json")
	reportJSON, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal report: %w", err)
	}

	if err := os.WriteFile(reportPath, reportJSON, 0644); err != nil {
		return fmt.Errorf("failed to write report file: %w", err)
	}
	fmt.Printf("[INFO] 📄 JSON report saved to: %s\n", reportPath)

	// 2. Save HTML report
	htmlPath := filepath.Join(outputDir, "scan_report.html")
	htmlContent := generateHTMLReport(report)
	if err := os.WriteFile(htmlPath, []byte(htmlContent), 0644); err != nil {
		return fmt.Errorf("failed to write HTML report: %w", err)
	}
	fmt.Printf("[INFO] 🌐 HTML report saved to: %s\n", htmlPath)

	// 3. Save detailed TXT/LOG report
	logPath := filepath.Join(outputDir, "scan_report.txt")
	logFile, err := os.Create(logPath)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer logFile.Close()

	successRate := 0.0
	if report.TotalURLs > 0 {
		successRate = float64(report.Successful) / float64(report.TotalURLs) * 100
	}

	logContent := fmt.Sprintf(`
╔════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║                    TOR SCRAPER - DETAILED SCAN REPORT                      ║
║                                                                            ║
╚════════════════════════════════════════════════════════════════════════════╝

📊 SCAN SUMMARY
═══════════════════════════════════════════════════════════════════════════

Start Time:       %s
End Time:         %s
Duration:         %v
Timestamp:        %s

📈 STATISTICS
═══════════════════════════════════════════════════════════════════════════

Total URLs:       %d
Successful:       %d
Failed:           %d
Success Rate:     %.2f%%

🔍 DETAILED RESULTS
═══════════════════════════════════════════════════════════════════════════

`, report.StartTime.Format(time.RFC3339),
		report.EndTime.Format(time.RFC3339),
		report.EndTime.Sub(report.StartTime),
		time.Now().Format(time.RFC3339),
		report.TotalURLs, report.Successful, report.Failed, successRate)

	logFile.WriteString(logContent)

	for i, result := range report.Results {
		logLine := fmt.Sprintf(`
[%d] URL: %s
    Status:       %s
    HTTP Code:    %d
    Timestamp:    %s
`, i+1, result.URL, result.Status, result.StatusCode, result.Timestamp.Format("2006-01-02 15:04:05"))

		if result.Error != "" {
			logLine += fmt.Sprintf("    Error:        %s\n", result.Error)
		}
		if result.Content != "" {
			logLine += fmt.Sprintf("    Content Size: %d bytes\n", len(result.Content))
		}
		logLine += "\n"
		logFile.WriteString(logLine)
	}

	logFile.WriteString("\n═══════════════════════════════════════════════════════════════════════════\n")
	logFile.WriteString(fmt.Sprintf("Report generated: %s\n", time.Now().Format(time.RFC3339)))
	logFile.WriteString("═══════════════════════════════════════════════════════════════════════════\n")

	fmt.Printf("[INFO] 📝 Text report saved to: %s\n", logPath)

	// 4. Save CSV report for data analysis
	csvPath := filepath.Join(outputDir, "scan_report.csv")
	csvFile, err := os.Create(csvPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer csvFile.Close()

	csvFile.WriteString("URL,Status,HTTP_Code,Timestamp,Content_Size,Error\n")
	for _, result := range report.Results {
		contentSize := 0
		if result.Content != "" {
			contentSize = len(result.Content)
		}
		errorMsg := result.Error
		if errorMsg == "" {
			errorMsg = "None"
		}
		csvLine := fmt.Sprintf("%s,%s,%d,%s,%d,\"%s\"\n",
			result.URL, result.Status, result.StatusCode,
			result.Timestamp.Format(time.RFC3339), contentSize, errorMsg)
		csvFile.WriteString(csvLine)
	}
	fmt.Printf("[INFO] 📊 CSV report saved to: %s\n", csvPath)

	// 5. Save individual HTML content files
	for _, result := range report.Results {
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

	// 6. Save summary report
	summaryPath := filepath.Join(outputDir, "SCAN_SUMMARY.txt")
	summaryFile, err := os.Create(summaryPath)
	if err != nil {
		return fmt.Errorf("failed to create summary file: %w", err)
	}
	defer summaryFile.Close()

	summaryContent := fmt.Sprintf(`TOR SCRAPER - SCAN SUMMARY
═══════════════════════════════════════════════════════════════════════════

✅ QUICK STATS:
   • Total Targets Scanned: %d
   • Successful: %d (%.1f%%)
   • Failed: %d (%.1f%%)

⏱️  TIMING:
   • Started: %s
   • Completed: %s
   • Duration: %v

📁 OUTPUT FILES GENERATED:
   • scan_report.json  - Machine-readable JSON format
   • scan_report.html  - Interactive HTML visualization
   • scan_report.txt   - Detailed text report
   • scan_report.csv   - CSV format for spreadsheets
   • content/          - Individual HTML files for each successful scan
   • SCAN_SUMMARY.txt  - This summary file

💡 NEXT STEPS:
   1. Open scan_report.html in a web browser for visualization
   2. Review scan_report.txt for detailed analysis
   3. Import scan_report.csv to Excel/Google Sheets
   4. Check individual content files in content/ directory

═══════════════════════════════════════════════════════════════════════════
Generated: %s
`,
		report.TotalURLs, report.Successful, successRate, report.Failed, 100-successRate,
		report.StartTime.Format("2006-01-02 15:04:05"),
		report.EndTime.Format("2006-01-02 15:04:05"),
		report.EndTime.Sub(report.StartTime),
		time.Now().Format("2006-01-02 15:04:05"))

	summaryFile.WriteString(summaryContent)
	fmt.Printf("[INFO] 📋 Summary saved to: %s\n", summaryPath)

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
