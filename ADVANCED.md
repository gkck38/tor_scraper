# Tor Scraper - Advanced Examples & Customization

## Table of Contents
1. [Advanced Configuration](#advanced-configuration)
2. [Custom Target Formats](#custom-target-formats)
3. [Performance Optimization](#performance-optimization)
4. [Extending Functionality](#extending-functionality)
5. [Troubleshooting Deep Dive](#troubleshooting-deep-dive)

---

## Advanced Configuration

### Custom Tor Port Configuration

If you're running Tor on a non-standard port, modify `main.go`:

```go
// Find this section in createTorClient():
dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
// Change 9050 to your port:
dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:CUSTOM_PORT", nil, proxy.Direct)
```

### Proxy with Authentication

If your Tor proxy requires authentication:

```go
// In createTorClient():
auth := &proxy.Auth{
    User:     "username",
    Password: "password",
}
dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", auth, proxy.Direct)
```

### Adjust Request Timeouts

For very slow .onion sites:

```go
// In createTorClient(), modify the transport:
transport := &http.Transport{
    Dial: dialer.Dial,
    TLSHandshakeTimeout:   30 * time.Second,  // Increase this
    ResponseHeaderTimeout: 30 * time.Second,  // Increase this
}

// And the client timeout:
client := &http.Client{
    Transport: transport,
    Timeout:   60 * time.Second,  // Increase this
}
```

### Custom User-Agent

Modify in `scanURL()` function:

```go
// Set custom User-Agent
userAgents := []string{
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36",
    "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36",
}
req.Header.Set("User-Agent", userAgents[rand.Intn(len(userAgents))])
```

---

## Custom Target Formats

### Format 1: Simple Text File

```
# targets.txt
example1.onion
example2.onion
example3.onion
```

### Format 2: YAML with Additional Metadata

```yaml
# targets.yaml
targets:
  - url: example1.onion
    priority: high
  - url: example2.onion
    priority: low
```

To parse this format, modify `readTargets()`:

```go
func readTargets(filepath string) ([]string, error) {
    // ... existing code ...
    
    // For YAML parsing, add after line trimming:
    if strings.Contains(line, "url:") {
        // Extract URL from "    url: example.onion"
        parts := strings.Split(line, ":")
        if len(parts) > 1 {
            url := strings.TrimSpace(parts[1])
            targets = append(targets, url)
        }
        continue
    }
    // ... rest of function ...
}
```

### Format 3: JSON File

```json
{
  "targets": [
    "example1.onion",
    "example2.onion"
  ]
}
```

To support JSON:

```go
import "encoding/json"

type TargetsJSON struct {
    Targets []string `json:"targets"`
}

// Add this function:
func readTargetsJSON(filepath string) ([]string, error) {
    data, err := os.ReadFile(filepath)
    if err != nil {
        return nil, err
    }
    
    var targets TargetsJSON
    err = json.Unmarshal(data, &targets)
    if err != nil {
        return nil, err
    }
    
    return targets.Targets, nil
}
```

---

## Performance Optimization

### 1. Parallel Scanning with Goroutines

Replace the serial scanning loop with concurrent version:

```go
// In main(), replace the for loop with:
resultChan := make(chan ScanResult, len(targets))
var wg sync.WaitGroup

fmt.Printf("[INFO] Starting %d concurrent scans...\n", len(targets))

for _, target := range targets {
    wg.Add(1)
    go func(t string) {
        defer wg.Done()
        result := scanURL(client, t)
        report.Results = append(report.Results, result)
        
        if result.Status == "SUCCESS" {
            atomic.AddInt64(&successCount, 1)
        } else {
            atomic.AddInt64(&failureCount, 1)
        }
    }(target)
    
    // Limit concurrent requests
    if len(targets) > 10 {
        time.Sleep(100 * time.Millisecond)
    }
}

wg.Wait()
```

At the top of main.go, add:

```go
import (
    // ... existing imports ...
    "sync"
    "sync/atomic"
)

var (
    successCount int64
    failureCount int64
)
```

### 2. Connection Pooling

Modify transport to reuse connections:

```go
transport := &http.Transport{
    Dial: dialer.Dial,
    MaxIdleConns:        100,
    MaxIdleConnsPerHost: 10,
    IdleConnTimeout:     90 * time.Second,
}
```

### 3. Batch Processing

Process multiple files in sequence:

```go
// In main():
targetFiles := []string{"targets1.yaml", "targets2.yaml", "targets3.yaml"}

for _, file := range targetFiles {
    fmt.Printf("[INFO] Processing batch: %s\n", file)
    targets, _ := readTargets(file)
    
    for _, target := range targets {
        result := scanURL(client, target)
        report.Results = append(report.Results, result)
    }
}
```

---

## Extending Functionality

### Add Response Filtering

```go
// Add after response headers are read:
func filterResponse(resp *http.Response) bool {
    // Only save responses with specific status codes
    acceptedCodes := []int{200, 301, 302, 401, 403}
    for _, code := range acceptedCodes {
        if resp.StatusCode == code {
            return true
        }
    }
    return false
}

// Use in scanURL():
if !filterResponse(resp) {
    result.Status = "FILTERED"
    return result
}
```

### Add Response Deduplication

```go
import "crypto/md5"

func deduplicateContent(content string) string {
    hash := md5.Sum([]byte(content))
    return fmt.Sprintf("%x", hash)
}

// In main(), track seen content:
seenContent := make(map[string]bool)

for _, result := range report.Results {
    if result.Status == "SUCCESS" {
        hash := deduplicateContent(result.Content)
        if seenContent[hash] {
            result.Status = "DUPLICATE"
        } else {
            seenContent[hash] = true
        }
    }
}
```

### Add Response Search/Regex

```go
import "regexp"

func searchInResponse(content, pattern string) bool {
    matched, _ := regexp.MatchString(pattern, content)
    return matched
}

// Use:
if searchInResponse(result.Content, "admin|login|password") {
    fmt.Printf("[INTERESTING] Found keyword in: %s\n", result.URL)
}
```

### Add Database Support

```go
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

func saveToDatabase(db *sql.DB, result ScanResult) error {
    _, err := db.Exec(`
        INSERT INTO scan_results (url, status, status_code, content, timestamp)
        VALUES (?, ?, ?, ?, ?)
    `, result.URL, result.Status, result.StatusCode, result.Content, result.Timestamp)
    return err
}

// Usage in main():
db, _ := sql.Open("sqlite3", "scans.db")
defer db.Close()

db.Exec(`CREATE TABLE IF NOT EXISTS scan_results (
    id INTEGER PRIMARY KEY,
    url TEXT,
    status TEXT,
    status_code INTEGER,
    content TEXT,
    timestamp DATETIME
)`)

for _, result := range report.Results {
    saveToDatabase(db, result)
}
```

---

## Troubleshooting Deep Dive

### Enable Debug Logging

```go
// Add at start of main():
log.SetFlags(log.LstdFlags | log.Lshortfile)
log.SetOutput(os.Stdout)

// Use in functions:
log.Printf("[DEBUG] Connecting to %s\n", url)
```

### Monitor Network Activity

```bash
# Linux/Mac - Watch network connections
sudo netstat -tlnp | grep go

# Windows - Check connections
netstat -ano | findstr tor-scraper
```

### Check Tor Exit Node

```bash
# Add this test function:
func checkTorIP() {
    resp, _ := http.Get("https://check.torproject.org/api/ip")
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    log.Printf("[DEBUG] Tor IP: %s\n", string(body))
}

// Call in main() before scanning:
checkTorIP()
```

### Memory Profiling

```bash
# Add to main():
go run -memprofile=mem.prof main.go targets.yaml

# Analyze:
go tool pprof mem.prof
```

### CPU Profiling

```bash
# Add to main():
go run -cpuprofile=cpu.prof main.go targets.yaml

# Analyze:
go tool pprof cpu.prof
```

---

## Example: Advanced Setup

Here's a complete example combining several advanced features:

```go
package main

import (
    "net/http"
    "sync"
    "time"
)

func advancedScan() {
    // Setup
    client, _ := createTorClient()
    targets, _ := readTargets("targets.yaml")
    
    // Configure
    maxConcurrency := 5
    semaphore := make(chan struct{}, maxConcurrency)
    var wg sync.WaitGroup
    
    // Parallel execution with concurrency limit
    for _, target := range targets {
        semaphore <- struct{}{}        // Acquire
        wg.Add(1)
        
        go func(t string) {
            defer wg.Done()
            defer func() { <-semaphore }() // Release
            
            result := scanURL(client, t)
            
            // Advanced processing
            if result.Status == "SUCCESS" {
                // Search for keywords
                if searchInResponse(result.Content, "admin") {
                    fmt.Printf("[ALERT] Admin panel found: %s\n", t)
                }
            }
        }(target)
        
        time.Sleep(100 * time.Millisecond)
    }
    
    wg.Wait()
}
```

---

## Performance Benchmarks

Typical performance on standard hardware:

- **Serial Scanning**: ~1-2 requests/second
- **5 Goroutines**: ~5-10 requests/second  
- **10 Goroutines**: ~10-15 requests/second
- **Network Limited**: ~0.5 requests/second (for slow .onion sites)

---

**Need help?** Refer to the README.md or QUICKSTART.md files!
