# Tor Scraper - .onion Scanner Tool

A Go-based automation tool for scanning .onion addresses through the Tor network to gather Cyber Threat Intelligence (CTI) data.

## Features

✅ **Tor Integration**: Routes all traffic through Tor SOCKS5 proxy  
✅ **Bulk Scanning**: Processes multiple .onion targets from a file  
✅ **Error Handling**: Continues scanning despite failures  
✅ **Content Collection**: Saves HTTP responses and metadata  
✅ **Detailed Reporting**: JSON report and log files with statistics  
✅ **Concurrency-Ready**: Framework supports goroutines for parallel scanning

## Prerequisites

- Go 1.21 or higher
- Tor service installed and running
- Tor SOCKS5 proxy listening on `127.0.0.1:9050` or `127.0.0.1:9150`

## Installation

### 1. Install Tor

**Windows:**
- Download from https://www.torproject.org/download/
- Install Tor Browser or Tor standalone

**Linux:**
```bash
sudo apt-get install tor
sudo systemctl start tor
```

**macOS:**
```bash
brew install tor
brew services start tor
```

### 2. Verify Tor is Running

```bash
# Test Tor connection
curl -x socks5://127.0.0.1:9050 https://check.torproject.org/

# Or install netstat to check if port 9050 is listening
# On Windows: netstat -ano | findstr 9050
# On Linux/Mac: netstat -tuln | grep 9050
```

## Usage

### Basic Usage

```bash
# Build the project
go build -o tor-scraper main.go

# Run with default targets file
./tor-scraper targets.yaml

# Or use custom output directory
./tor-scraper targets.yaml my_output_folder

# With go run directly
go run main.go targets.yaml
```

### Prepare Target File

Create a file (e.g., `targets.yaml`) with .onion addresses:

```yaml
# targets.yaml
3g2upl4pq6kufc4m.onion
thehiddenwiki.onion
http://example1.onion
http://example2.onion
```

Lines starting with `#` are ignored, as are empty lines.

## Output Structure

The tool generates the following output:

```
output/
├── scan_report.json      # Full scan results in JSON format
├── scan_report.log       # Human-readable scan summary
└── content/              # Individual HTML files from successful scans
    ├── example1.onion.html
    ├── example2.onion.html
    └── ...
```

### scan_report.json Format

```json
{
  "total_urls": 3,
  "successful": 2,
  "failed": 1,
  "start_time": "2024-01-15T10:30:00Z",
  "end_time": "2024-01-15T10:35:00Z",
  "results": [
    {
      "url": "http://example.onion",
      "status": "SUCCESS",
      "status_code": 200,
      "timestamp": "2024-01-15T10:30:05Z",
      "content": "..."
    }
  ]
}
```

### scan_report.log Format

```
Tor Scraper Scan Report
=======================
Start Time: 2024-01-15T10:30:00Z
End Time: 2024-01-15T10:35:00Z
Duration: 5m0s

Statistics:
-----------
Total URLs: 3
Successful: 2
Failed: 1
Success Rate: 66.67%

Detailed Results:
-----------------
[2024-01-15 10:30:05] http://example1.onion - Status: SUCCESS (HTTP 200)
[2024-01-15 10:30:15] http://example2.onion - Status: FAILED - Error: ...
```

## Project Structure

```
ikinci_gorev/
├── main.go            # Main application code
├── go.mod             # Go module definition
├── go.sum             # Go dependencies (auto-generated)
├── targets.yaml       # Target .onion addresses
├── README.md          # This file
└── output/            # Generated reports and content
```

## Code Architecture

### Main Components

1. **readTargets()** - Parses targets from YAML/text file
2. **createTorClient()** - Sets up HTTP client with SOCKS5 proxy
3. **scanURL()** - Fetches content from individual .onion address
4. **saveScanReport()** - Writes results to JSON and log files
5. **main()** - Orchestrates the entire scanning process

### Request Flow

```
User Command
    ↓
Read Targets File
    ↓
Create Tor HTTP Client
    ↓
For Each Target:
    ├─ Make HTTP Request (via Tor)
    ├─ Capture Response
    └─ Record Result
    ↓
Generate Reports
    ↓
Complete
```

## Configuration

### Tor Proxy Ports

The tool tries these ports in order:
1. `127.0.0.1:9050` (default Tor port)
2. `127.0.0.1:9150` (Tor Browser port)

If using a different port, modify this line in `main.go`:
```go
dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:YOUR_PORT", nil, proxy.Direct)
```

### Request Timeouts

- Connection timeout: 10 seconds
- Response timeout: 10 seconds
- Total request timeout: 15 seconds
- Response body limit: 1 MB

Modify in `main.go`:
```go
TLSHandshakeTimeout:   10 * time.Second,
ResponseHeaderTimeout: 10 * time.Second,
Timeout:               15 * time.Second,
```

### Delay Between Requests

Default: 1 second between requests

Modify in `main.go`:
```go
time.Sleep(1 * time.Second)
```

## Troubleshooting

### Error: "failed to create SOCKS5 dialer"

**Solution:**
- Verify Tor service is running
- Check if port 9050 or 9150 is listening: `netstat -ano | findstr 9050`
- Restart Tor service

### Error: "Request failed: context deadline exceeded"

**Solution:**
- Increase timeouts in `main.go`
- .onion site is slow or offline
- Network connection issues

### Error: "Failed to read targets"

**Solution:**
- Ensure targets file exists and path is correct
- File should contain valid URLs or domain names
- Check file permissions

## Advanced Usage

### Parallel Scanning with Goroutines

Modify the scanning loop for concurrent requests:

```go
// Create a buffered channel
resultChan := make(chan ScanResult, len(targets))

// Launch goroutines
for _, target := range targets {
    go func(t string) {
        resultChan <- scanURL(client, t)
    }(target)
}

// Collect results
for i := 0; i < len(targets); i++ {
    report.Results = append(report.Results, <-resultChan)
}
```

### Custom Headers

Modify request headers in `scanURL()`:
```go
req.Header.Set("User-Agent", "Custom User Agent")
req.Header.Set("Accept-Language", "en-US")
```

## Legal Notice

⚠️ This tool is designed for:
- Educational purposes
- Authorized security research
- Cyber threat intelligence gathering
- Authorized network testing

Unauthorized access to computer networks is illegal. Ensure you have proper authorization before scanning any targets.

## License

This project is provided as-is for educational purposes.

## Support

For issues or questions:
1. Check the troubleshooting section above
2. Verify Tor is running correctly
3. Review the generated log files for detailed errors
4. Check your targets file format

---

**Version:** 1.0  
**Last Updated:** 2024-01-15  
**Language:** Go 1.21+
