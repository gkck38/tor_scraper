# 📋 Tor Scraper Project - Complete Index

Welcome! This is your **Tor Scraper** - a complete Cyber Threat Intelligence tool written in Go.

## 🚀 Start Here

### For First-Time Users (Choose One)

**⏱️ 5-Minute Quick Start:**
→ Read: [QUICKSTART.md](QUICKSTART.md)

**📚 Complete Documentation:**
→ Read: [README.md](README.md)

**🛠️ Advanced Customization:**
→ Read: [ADVANCED.md](ADVANCED.md)

**📊 Project Overview:**
→ Read: [PROJECT_SUMMARY.md](PROJECT_SUMMARY.md)

**📁 File Descriptions:**
→ Read: [FILES_OVERVIEW.md](FILES_OVERVIEW.md)

---

## 📂 What's Included

### Core Application
- **main.go** - Complete Tor scraper (430 lines)
  - Reads .onion target lists
  - Routes traffic through Tor proxy
  - Scans URLs and collects data
  - Generates comprehensive reports

### Configuration
- **targets.yaml** - Sample .onion addresses
- **go.mod** - Go dependencies

### Automation
- **run.bat** - Windows auto-runner (one-click)
- **run.sh** - Linux/Mac auto-runner

### Documentation (5 Complete Guides)
1. README.md - Full documentation (3,500 words)
2. QUICKSTART.md - 5-minute setup (2,000 words)
3. ADVANCED.md - Advanced features (3,500 words)
4. PROJECT_SUMMARY.md - Project overview (2,000 words)
5. FILES_OVERVIEW.md - File descriptions (1,500 words)

---

## ✅ What This Tool Does

```
Input File (targets.yaml)
    ↓
Read .onion addresses
    ↓
Connect to Tor Network
    ↓
For each address:
  - Send HTTP request (via Tor)
  - Capture response
  - Log result (success/failure)
    ↓
Generate Reports:
  - JSON format (structured data)
  - Log file (readable summary)
  - HTML files (content backup)
    ↓
Output Folder
```

**Key Features:**
- ✅ Bulk scanning of multiple .onion sites
- ✅ Error-tolerant (continues on failures)
- ✅ Detailed reporting (JSON + logs)
- ✅ Content preservation (saves HTML)
- ✅ Easy to use (automated runners)
- ✅ Production-ready (robust error handling)

---

## 🎯 Quick Start (30 Seconds)

### Step 1: Install Tor
- Download from: https://www.torproject.org/
- Install and run
- Verify port 9050 is open

### Step 2: Run Scraper
```bash
# Windows: Double-click
run.bat

# Linux/Mac: Run
./run.sh

# Or manually:
go run main.go targets.yaml
```

### Step 3: Check Results
```bash
cd output
# See: scan_report.json, scan_report.log, content/
```

---

## 📖 Reading Guide

### Path 1: Just Want to Use It (15 min)
1. [QUICKSTART.md](QUICKSTART.md) - Sections "Quick Start (5 Minutes)"
2. Create/edit targets.yaml
3. Run run.bat or run.sh
4. Done!

### Path 2: Understand How It Works (1 hour)
1. [README.md](README.md) - Sections "Features" and "Architecture"
2. [FILES_OVERVIEW.md](FILES_OVERVIEW.md) - Understand each file
3. Skim through main.go
4. Run the tool
5. Review output files

### Path 3: Learn Everything & Customize (2-4 hours)
1. [README.md](README.md) - Read completely
2. [QUICKSTART.md](QUICKSTART.md) - Read completely
3. [ADVANCED.md](ADVANCED.md) - Study customization options
4. Study main.go code thoroughly
5. Experiment with modifications

### Path 4: Troubleshooting (As needed)
1. [QUICKSTART.md](QUICKSTART.md) - Section "Common Issues"
2. [README.md](README.md) - Section "Troubleshooting"
3. [ADVANCED.md](ADVANCED.md) - Section "Troubleshooting Deep Dive"

---

## 🔧 System Requirements

**Required:**
- Go 1.21+
- Tor service (running on 9050 or 9150)
- Internet connection

**Supported OS:**
- Windows 10+
- Linux (Ubuntu, Debian, Fedora, etc.)
- macOS
- Any OS with Go support

**Disk Space:**
- 100 MB for source code and dependencies
- 50-500 MB for scan results (varies)

**Network:**
- Tor SOCKS5 proxy access
- Internet access to .onion sites

---

## 📊 Project Structure

```
ikinci_gorev/
│
├─ MAIN APPLICATION
│  ├─ main.go           ← Core scraper (START HERE for code)
│  ├─ go.mod            ← Dependencies
│  └─ go.sum            ← Lock file (auto-generated)
│
├─ CONFIGURATION
│  └─ targets.yaml      ← Your target list (EDIT THIS)
│
├─ RUNNERS (Pick one based on your OS)
│  ├─ run.bat           ← Windows automated runner
│  └─ run.sh            ← Linux/Mac automated runner
│
├─ DOCUMENTATION (Choose based on your need)
│  ├─ QUICKSTART.md     ← 5-minute setup guide
│  ├─ README.md         ← Complete documentation
│  ├─ ADVANCED.md       ← Advanced customization
│  ├─ PROJECT_SUMMARY.md ← Project overview
│  ├─ FILES_OVERVIEW.md ← File descriptions
│  └─ INDEX.md          ← This file
│
└─ GENERATED ON FIRST RUN
   └─ output/
      ├─ scan_report.json   ← Results (JSON)
      ├─ scan_report.log    ← Results (readable)
      └─ content/          ← Downloaded HTML files
```

---

## 🚀 Getting Started (Choose Your Path)

### Path A: Automated (Easiest)
```bash
# Windows
double-click run.bat

# Linux/Mac
chmod +x run.sh
./run.sh
```
✅ Does everything automatically

### Path B: Manual (Standard)
```bash
# Install dependencies
go mod download
go mod tidy

# Build
go build -o tor-scraper main.go

# Run
./tor-scraper targets.yaml
```

### Path C: Direct (Fastest)
```bash
# Just run it
go run main.go targets.yaml
```

---

## 📝 What to Edit

### Only File You MUST Edit: targets.yaml

Add your .onion addresses here:
```yaml
# targets.yaml

# Add one per line
example1.onion
example2.onion
example3.onion
http://example4.onion
http://example5.onion

# Comments start with #
```

### Optional to Edit: main.go

For customization only:
- Adjust timeouts
- Change Tor port
- Add features
- See ADVANCED.md for examples

### Should NOT Edit
- go.mod (unless adding dependencies)
- Documentation files
- Run scripts (unless changing defaults)

---

## 📊 Example Output

### Console
```
[INFO] Reading targets from: targets.yaml
[INFO] Found 5 targets
[SUCCESS] Connected to Tor proxy
[INFO] Scanning: http://example1.onion
[SUCCESS] Scanning: http://example1.onion -> Status: 200
[ERR] Scanning: http://example2.onion -> FAILED (timeout)
Successful: 4/5
```

### Files
```
output/
├── scan_report.json (detailed results)
├── scan_report.log (human readable)
└── content/
    ├── example1.onion.html
    ├── example3.onion.html
    └── example4.onion.html
```

---

## 🆘 Help & Support

### Problem: Program won't start
→ [QUICKSTART.md](QUICKSTART.md#common-issues--solutions)

### Problem: Can't connect to Tor
→ [README.md](README.md#troubleshooting)

### Problem: Want to customize
→ [ADVANCED.md](ADVANCED.md)

### Problem: Want to understand code
→ [FILES_OVERVIEW.md](FILES_OVERVIEW.md#file-descriptions)

### Problem: Need quick reference
→ [PROJECT_SUMMARY.md](PROJECT_SUMMARY.md)

---

## 📚 Documentation Map

```
YOU ARE HERE: INDEX.md (This file)
│
├─→ First Time?      Read: QUICKSTART.md
├─→ Need Details?    Read: README.md
├─→ Want Fancy?      Read: ADVANCED.md
├─→ Need Context?    Read: PROJECT_SUMMARY.md
├─→ Which File?      Read: FILES_OVERVIEW.md
│
└─→ Source Code:     main.go (430 lines, well-commented)
```

---

## ⚡ Quick Commands

```bash
# Run with defaults
go run main.go targets.yaml

# Run with custom output
go run main.go targets.yaml my_reports

# Build executable
go build -o tor-scraper main.go

# Run built executable
./tor-scraper targets.yaml

# Windows automation
run.bat

# Linux/Mac automation
./run.sh

# Check Tor connection
netstat -ano | findstr 9050

# View results
cat output/scan_report.log
```

---

## 🎓 Learning Path

**Level 1: User (15 min)**
- Run the tool
- Review output
- Done!

**Level 2: Operator (1 hour)**
- Understand how it works
- Customize targets.yaml
- Troubleshoot issues

**Level 3: Developer (2-4 hours)**
- Study the code
- Modify and extend
- Add features
- Optimize performance

---

## ✨ Features Checklist

- ✅ Read .onion addresses from file
- ✅ Connect to Tor network
- ✅ Send HTTP requests anonymously
- ✅ Handle errors gracefully
- ✅ Save results to JSON
- ✅ Generate readable logs
- ✅ Preserve HTML content
- ✅ Calculate statistics
- ✅ Windows automation
- ✅ Linux/Mac automation
- ✅ Complete documentation
- ✅ Production-ready

---

## 🔐 Security Notes

⚠️ This tool is for:
- ✅ Educational purposes
- ✅ Authorized security research
- ✅ CTI gathering on authorized targets
- ✅ Authorized network testing

❌ Not for:
- Unauthorized network access
- Illegal surveillance
- Hacking or cracking
- Bypassing security intentionally

---

## 📞 Next Steps

1. **Right Now:** Read [QUICKSTART.md](QUICKSTART.md)
2. **In 5 Min:** Install Tor if needed
3. **In 10 Min:** Create targets.yaml
4. **In 15 Min:** Run the tool
5. **In 20 Min:** Review results

---

## 📊 File Summary

| File | Purpose | Read? | Edit? |
|------|---------|-------|-------|
| main.go | Application code | Optional | Advanced users |
| targets.yaml | Target list | No | **YES, required** |
| go.mod | Dependencies | No | No |
| run.bat | Windows runner | No | Optional |
| run.sh | Unix runner | No | Optional |
| QUICKSTART.md | Fast setup | **YES, first** | No |
| README.md | Full docs | After QUICKSTART | No |
| ADVANCED.md | Advanced | For customization | No |
| PROJECT_SUMMARY.md | Overview | Optional | No |
| FILES_OVERVIEW.md | File guide | Optional | No |

---

## 🎯 Success Metrics

✅ **Setup Success:**
- Go installed
- Tor running
- targets.yaml created

✅ **Run Success:**
- No build errors
- Connected to Tor
- Scans completed

✅ **Output Success:**
- scan_report.json created
- scan_report.log readable
- Content files saved

---

## 📌 Remember

- 📖 **Start with QUICKSTART.md** (not this file!)
- 🎯 **Edit only targets.yaml** (add your .onion addresses)
- ▶️ **Run: `go run main.go targets.yaml`**
- 📁 **Check results in output/ folder**
- 📚 **Read README.md for full details**
- 🆘 **See ADVANCED.md if you need help**

---

## 🚀 Ready?

Pick your starting point:

- **I want to use it NOW** → [QUICKSTART.md](QUICKSTART.md)
- **I want to understand it** → [README.md](README.md)
- **I want to customize it** → [ADVANCED.md](ADVANCED.md)
- **I want overview** → [PROJECT_SUMMARY.md](PROJECT_SUMMARY.md)
- **I want file details** → [FILES_OVERVIEW.md](FILES_OVERVIEW.md)

---

**Version:** 1.0  
**Status:** ✅ Ready to Use  
**Language:** Go 1.21+  
**Time to Setup:** < 5 minutes  
**Documentation:** Complete (12,000+ words)

Enjoy your Tor Scraper! 🎉

---

*Questions? See the appropriate documentation file above.*

*Still stuck? Check QUICKSTART.md → Common Issues & Solutions*
