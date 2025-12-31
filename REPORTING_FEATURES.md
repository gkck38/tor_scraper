# 📊 Advanced Reporting Features

## Overview
The Tor Scraper now includes comprehensive reporting capabilities that generate multiple output formats for detailed analysis and visualization of scan results.

## 📁 Generated Report Files

### 1. **scan_report.json**
- **Format:** Machine-readable JSON
- **Purpose:** Data integration with other tools and systems
- **Contents:**
  - Total URLs scanned
  - Success/failure statistics
  - Detailed results for each URL (status, HTTP code, timestamp, content)
  - Start and end times

**Example:**
```json
{
  "total_urls": 5,
  "successful": 5,
  "failed": 0,
  "start_time": "2025-12-31T21:16:52+03:00",
  "end_time": "2025-12-31T21:16:57+03:00",
  "results": [...]
}
```

### 2. **scan_report.html** ⭐ (NEW)
- **Format:** Interactive HTML webpage
- **Purpose:** Visual analysis and presentation
- **Features:**
  - Professional gradient design
  - Real-time statistics cards
  - Progress bar showing success rate
  - Detailed results table
  - Responsive layout (works on desktop and mobile)
  - Color-coded status indicators

**Usage:**
```bash
# Open the HTML report in your browser
open output/scan_report.html  # macOS
start output\scan_report.html # Windows
xdg-open output/scan_report.html # Linux
```

### 3. **scan_report.txt** ⭐ (ENHANCED)
- **Format:** Formatted text with ASCII art
- **Purpose:** Human-readable detailed analysis
- **Sections:**
  - Scan summary with borders
  - Comprehensive statistics
  - Detailed per-URL results
  - Success rate percentage
  - Timestamps in RFC3339 format

### 4. **scan_report.csv** ⭐ (NEW)
- **Format:** Comma-separated values
- **Purpose:** Excel/Google Sheets integration
- **Columns:**
  - URL
  - Status (SUCCESS/FAILED/ERROR/PARTIAL)
  - HTTP Status Code
  - Timestamp
  - Content Size (in bytes)
  - Error Message (if any)

**Usage:**
```bash
# Import into Excel or Google Sheets
# 1. Open your spreadsheet application
# 2. File > Import > Select scan_report.csv
# 3. Analyze with pivot tables and charts
```

### 5. **SCAN_SUMMARY.txt** ⭐ (NEW)
- **Format:** Quick reference guide
- **Purpose:** Executive summary
- **Contains:**
  - Quick statistics
  - Timing information
  - List of generated files
  - Next steps for analysis

### 6. **content/** (Directory)
- **Format:** Individual HTML files
- **Purpose:** Content archival
- **One file per successful scan:**
  - Automatically extracts and saves webpage content
  - Preserves original HTML structure
  - Useful for offline analysis

## 📊 Statistics Provided

### Basic Metrics
- ✅ Total URLs scanned
- ✅ Successful scans
- ✅ Failed scans
- ✅ Success rate (percentage)

### Timing Information
- 📅 Start timestamp
- 📅 End timestamp
- ⏱️ Total duration
- ⏱️ Per-request timing

### Per-URL Details
- 🔗 Full URL
- 📍 Status (SUCCESS/FAILED/ERROR/PARTIAL)
- 🌐 HTTP status code
- ⏰ Request timestamp
- 📦 Content size (bytes)
- ❌ Error messages (if applicable)

## 🎯 Use Cases

### 1. **Compliance & Auditing**
- Use HTML report for presentations
- CSV export for database integration
- JSON for API consumption

### 2. **Data Analysis**
- Import CSV into Excel/Google Sheets
- Create pivot tables and charts
- Compare success rates across different scans

### 3. **Threat Intelligence**
- Parse JSON for automated processing
- Extract content files for malware analysis
- Track changes over time with multiple scans

### 4. **Team Reporting**
- Share HTML reports via email
- Use summary file for quick briefings
- Archive complete scan reports

## 📈 Example Analysis Workflow

```bash
# 1. Run the scan
./tor-scraper targets.yaml

# 2. Review quick summary
cat output/SCAN_SUMMARY.txt

# 3. Detailed analysis in text format
cat output/scan_report.txt

# 4. Open interactive HTML report
open output/scan_report.html

# 5. Export to spreadsheet (Windows)
start output\scan_report.csv

# 6. Access individual content
ls output/content/
```

## 🔧 Customization Options

### Output Directory
By default, reports are saved to `output/`. To use a different directory:

```bash
./tor-scraper targets.yaml /path/to/custom/output
```

### File Naming
All report files follow the naming convention:
- `scan_report.[format]` - Main reports
- `SCAN_SUMMARY.txt` - Summary file
- `content/[url].html` - Individual content files

## 📋 File Structure Example

```
output/
├── scan_report.json        # Machine-readable data
├── scan_report.html        # Visual report
├── scan_report.txt         # Detailed text report
├── scan_report.csv         # Spreadsheet format
├── SCAN_SUMMARY.txt        # Executive summary
├── scan_report.log         # (Legacy log file)
└── content/
    ├── 3g2upl4pq3khfchc.onion.html
    ├── thehiddenwiki.onion.html
    ├── msydqstlz5daysqf.onion.html
    ├── kingpin5gzmk4zd3.onion.html
    └── nothiddenwiki.com.html
```

## 🚀 Advanced Features

### Color-Coded Status in HTML
- 🟢 SUCCESS - HTTP 200 response
- 🔴 FAILED - Connection/timeout error
- 🟠 PARTIAL - Content partially retrieved
- ⚪ ERROR - Request creation error

### Responsive HTML Design
- Desktop: Full statistics grid (4 columns)
- Tablet: 2-column layout
- Mobile: Single column layout

### Data Integrity
- All timestamps in RFC3339 format (ISO 8601)
- Content size in bytes
- Original error messages preserved
- No data loss or truncation

## 💡 Tips & Tricks

### Create Multiple Reports
```bash
# Archive different scans by date
mkdir -p reports/scan-2025-12-31
./tor-scraper targets.yaml reports/scan-2025-12-31
```

### Compare Scans Over Time
```bash
# Use CSV exports for historical analysis
diff <(tail -n +2 reports/scan1/scan_report.csv | cut -d, -f1,2) \
     <(tail -n +2 reports/scan2/scan_report.csv | cut -d, -f1,2)
```

### Automated Report Sharing
```bash
# Email the HTML report
mail -s "Tor Scan Report" recipient@example.com < output/scan_report.html
```

## 🔐 Security Notes

- Reports contain sensitive scan data
- Store in secure locations
- Avoid sharing raw reports in untrusted channels
- Consider encrypting archived reports
- Individual content files may contain malicious code - handle with care

## 📝 Version History

- **v1.0** - Initial reporting system
  - ✅ JSON export
  - ✅ Text log file
  - ✅ Individual content files

- **v1.1** - Advanced reporting
  - ✅ Interactive HTML reports
  - ✅ CSV for data analysis
  - ✅ Executive summary
  - ✅ Enhanced text formatting
  - ✅ Progress visualization

---

**For questions or issues, refer to the main README.md file.**
