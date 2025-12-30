# Tor Scraper - Quick Start Guide

## What is This?

This is a **Tor Scraper** - a Go program that automatically scans multiple .onion (dark web) addresses through the Tor network to collect information. It's used for Cyber Threat Intelligence (CTI) gathering.

## What You Get

✅ **4 Complete Modules:**
1. File Reader - Reads list of .onion addresses
2. Tor Proxy Client - Routes traffic through Tor
3. Request Handler - Sends HTTP requests and manages errors
4. Output Writer - Saves results and creates reports

✅ **Deliverables:**
- Compiled binary (ready to run)
- Bulk dataset of accessible .onion content
- Status report with active/inactive sites

## Before You Start

### 1. Install Tor Service

**Windows:**
- Download from: https://www.torproject.org/download/
- Install and run (Tor Browser or standalone Tor)
- Default port: 9050 or 9150

**Linux (Ubuntu/Debian):**
```bash
sudo apt-get install tor
sudo systemctl start tor
sudo systemctl enable tor
```

**Verify Tor is Running:**
```bash
# Windows: Check if port 9050 is listening
netstat -ano | findstr 9050

# Linux/Mac:
sudo netstat -tlnp | grep 9050
```

### 2. Get Go

Download Go 1.21+ from: https://golang.org/dl/

Verify installation:
```bash
go version
```

## Quick Start (5 Minutes)

### Step 1: Navigate to Project

```bash
cd c:\Users\meg\Desktop\ctı+i\ikinci_gorev
```

### Step 2: Download Dependencies

```bash
go mod download
go mod tidy
```

### Step 3: Create Your Target List

Edit `targets.yaml` and add .onion addresses:

```yaml
# targets.yaml
3g2upl4pq6kufc4m.onion
thehiddenwiki.onion
example1.onion
example2.onion
example3.onion
```

### Step 4: Run the Scanner

```bash
# Option 1: With go run
go run main.go targets.yaml

# Option 2: Compile first, then run
go build -o tor-scraper.exe main.go
./tor-scraper.exe targets.yaml
```

### Step 5: Check Results

```bash
# Reports are saved in 'output' folder
output/scan_report.json       # All data in JSON
output/scan_report.log        # Human-readable summary
output/content/               # Individual HTML files
```

## Expected Output

### Console Output:
```
========================================
      Tor Scraper - .onion Scanner
========================================

[INFO] Reading targets from: targets.yaml
[INFO] Found 3 targets

[INFO] Connecting to Tor network...
[SUCCESS] Connected to Tor proxy

[INFO] Starting scan...

[INFO] Scanning: http://example1.onion
[SUCCESS] Scanning: http://example1.onion -> Status: 200

[INFO] Scanning: http://example2.onion
[ERR] Scanning: http://example2.onion -> FAILED (timeout)

[INFO] Scanning: http://example3.onion
[SUCCESS] Scanning: http://example3.onion -> Status: 200

========================================
           Scan Complete
========================================
Duration: 3m45s
Successful: 2/3

[INFO] Report saved to: output\scan_report.json
[INFO] Log file saved to: output\scan_report.log
[SUCCESS] Scan complete!
```

### Report Output (scan_report.log):
```
Tor Scraper Scan Report
=======================
Start Time: 2024-01-15T10:30:00Z
End Time: 2024-01-15T10:33:45Z
Duration: 3m45s

Statistics:
-----------
Total URLs: 3
Successful: 2
Failed: 1
Success Rate: 66.67%

Detailed Results:
-----------------
[2024-01-15 10:30:05] http://example1.onion - Status: SUCCESS (HTTP 200)
[2024-01-15 10:30:45] http://example2.onion - Status: FAILED - Error: context deadline exceeded
[2024-01-15 10:31:15] http://example3.onion - Status: SUCCESS (HTTP 200)
```

## Files in This Project

```
ikinci_gorev/
├── main.go           ← Main Go program (4 modules inside)
├── go.mod            ← Project dependencies
├── targets.yaml      ← Your .onion target list (EDIT THIS)
├── README.md         ← Detailed documentation
├── QUICKSTART.md     ← This file
└── output/           ← Generated results (created after running)
    ├── scan_report.json
    ├── scan_report.log
    └── content/
        └── [website files]
```

## Common Issues & Solutions

### "Failed to create Tor client"
**Problem:** Tor not running or wrong port  
**Solution:**
```bash
# Start Tor service
# Windows: Start Tor Browser
# Linux: sudo systemctl start tor
# Then check port 9050 is open
```

### "Request failed: context deadline exceeded"
**Problem:** .onion site is slow or offline  
**Solution:** Increase timeout in main.go or check if site is active

### "Failed to read targets"
**Problem:** targets.yaml not found or wrong path  
**Solution:**
```bash
# Make sure targets.yaml is in same folder as main.go
# Check file exists: dir targets.yaml
```

### "Connection refused on port 9050"
**Problem:** Tor service not listening  
**Solution:**
```bash
# Check if Tor is running
netstat -ano | findstr 9050

# If not, start Tor and wait 30 seconds
```

## Command Examples

```bash
# Basic usage
go run main.go targets.yaml

# Save to custom output directory
go run main.go targets.yaml my_reports

# Build as executable
go build -o tor_scraper.exe main.go
./tor_scraper.exe targets.yaml

# Run with different target file
go run main.go /path/to/urls.txt output
```

## How It Works (Simple Explanation)

1. **Reads File**: Gets list of .onion addresses from `targets.yaml`
2. **Connects to Tor**: Routes all internet traffic through Tor proxy
3. **Scans Each URL**:
   - Sends HTTP request via Tor
   - If successful: saves website content
   - If fails: records error, continues to next
4. **Saves Results**:
   - Detailed JSON report
   - Human-readable log file
   - Individual HTML files from each site

## Advanced Features (Optional)

### Parallel Scanning (Goroutines)
The code is designed for concurrent requests. Uncomment the goroutine section in main.go for faster scanning of many targets.

### Custom Headers
Modify User-Agent and other headers in scanURL() function:
```go
req.Header.Set("User-Agent", "Your Custom Agent")
```

### Request Timeouts
Adjust in main.go:
```go
timeout:   15 * time.Second,  // Change this
```

## Proof of Working Tool

To prove the tool works with Tor:

1. **Take Screenshot** of console showing:
   - "[SUCCESS] Connected to Tor proxy"
   - URLs being scanned
   - Results (SUCCESS/FAILED)

2. **Check Output Files**:
   - `output/scan_report.json` exists
   - `output/scan_report.log` has your results
   - `output/content/` has website files

3. **Verify IP Changes**:
   - Before running: `curl https://check.torproject.org/` (shows real IP)
   - During scan: Same command via Tor (shows Tor IP)

## Next Steps

1. ✅ Install Tor
2. ✅ Verify Go is installed
3. ✅ Edit targets.yaml with real .onion addresses
4. ✅ Run: `go run main.go targets.yaml`
5. ✅ Check results in output/ folder
6. ✅ Take screenshots for submission

## Support Resources

- **Tor Help**: https://www.torproject.org/
- **Go Docs**: https://golang.org/doc/
- **Go Proxy Package**: https://pkg.go.dev/golang.org/x/net/proxy

---

**Ready?** Start with Step 1 of "Quick Start" above! 🚀
