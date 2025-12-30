#!/bin/bash

# Tor Scraper Linux/Mac Helper Script
# This script helps run the Tor Scraper on Linux and macOS

echo ""
echo "========================================"
echo "   Tor Scraper - Unix Helper Script"
echo "========================================"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "[ERROR] Go is not installed or not in PATH!"
    echo "Please install Go from: https://golang.org/dl/"
    read -p "Press enter to exit..."
    exit 1
fi

echo "[OK] Go is installed"
go version
echo ""

# Check if targets.yaml exists
if [ ! -f targets.yaml ]; then
    echo "[ERROR] targets.yaml not found!"
    echo "Please create targets.yaml with .onion addresses"
    read -p "Press enter to exit..."
    exit 1
fi

echo "[OK] targets.yaml found"
echo ""

# Check Tor connection
echo "[INFO] Checking Tor connection on port 9050..."
if nc -z 127.0.0.1 9050 2>/dev/null; then
    echo "[OK] Tor is running on port 9050"
elif nc -z 127.0.0.1 9150 2>/dev/null; then
    echo "[OK] Tor is running on port 9150"
else
    echo "[ERROR] Cannot connect to Tor!"
    echo ""
    echo "Instructions:"
    echo "1. Linux: sudo systemctl start tor"
    echo "2. macOS: brew services start tor"
    echo "3. Or download Tor Browser from: https://www.torproject.org/download/"
    read -p "Press enter to exit..."
    exit 1
fi

echo ""
echo "[INFO] Downloading dependencies..."
go mod download
go mod tidy

echo ""
echo "[INFO] Building project..."
rm -f tor-scraper
go build -o tor-scraper main.go

if [ $? -ne 0 ]; then
    echo "[ERROR] Build failed!"
    read -p "Press enter to exit..."
    exit 1
fi

echo "[OK] Build successful: tor-scraper created"
echo ""

# Run the scraper
echo "[INFO] Starting scan..."
echo ""
./tor-scraper targets.yaml output

if [ $? -ne 0 ]; then
    echo ""
    echo "[ERROR] Scan failed!"
    read -p "Press enter to exit..."
    exit 1
fi

echo ""
echo "========================================"
echo "[SUCCESS] Scan completed!"
echo "========================================"
echo ""
echo "Results saved to:"
echo "  - output/scan_report.json"
echo "  - output/scan_report.log"
echo "  - output/content/[website files]"
echo ""

read -p "Press enter to view results..."
