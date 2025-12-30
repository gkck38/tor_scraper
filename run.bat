@echo off
REM Tor Scraper Windows Helper Script
REM This script helps run the Tor Scraper on Windows

echo.
echo ========================================
echo    Tor Scraper - Windows Helper
echo ========================================
echo.

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Go is not installed or not in PATH!
    echo Please install Go from: https://golang.org/dl/
    pause
    exit /b 1
)

echo [OK] Go is installed
go version
echo.

REM Check if targets.yaml exists
if not exist targets.yaml (
    echo [ERROR] targets.yaml not found!
    echo Please create targets.yaml with .onion addresses
    pause
    exit /b 1
)

echo [OK] targets.yaml found
echo.

REM Check Tor connection
echo [INFO] Checking Tor connection on port 9050...
netstat -ano | findstr 9050 >nul
if %ERRORLEVEL% NEQ 0 (
    echo [WARN] Port 9050 not responding
    echo [INFO] Trying port 9150...
    netstat -ano | findstr 9150 >nul
    if %ERRORLEVEL% NEQ 0 (
        echo [ERROR] Neither port 9050 nor 9150 is open!
        echo [ERROR] Please start Tor and try again.
        echo.
        echo Instructions:
        echo 1. Download Tor: https://www.torproject.org/download/
        echo 2. Install and run Tor Browser
        echo 3. Run this script again
        pause
        exit /b 1
    ) else (
        echo [OK] Tor is running on port 9150
    )
) else (
    echo [OK] Tor is running on port 9050
)

echo.
echo [INFO] Downloading dependencies...
call go mod download
call go mod tidy

echo.
echo [INFO] Building project...
if exist tor-scraper.exe del tor-scraper.exe
call go build -o tor-scraper.exe main.go

if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Build failed!
    pause
    exit /b 1
)

echo [OK] Build successful: tor-scraper.exe created
echo.

REM Run the scraper
echo [INFO] Starting scan...
echo.
call tor-scraper.exe targets.yaml output

if %ERRORLEVEL% NEQ 0 (
    echo.
    echo [ERROR] Scan failed!
    pause
    exit /b 1
)

echo.
echo ========================================
echo [SUCCESS] Scan completed!
echo ========================================
echo.
echo Results saved to:
echo   - output\scan_report.json
echo   - output\scan_report.log
echo   - output\content\[website files]
echo.
echo Opening results folder...
explorer.exe output

echo.
pause
