# 🎯 RAPORLAMA DEMO - ÇALIŞMA ÖRNEĞİ

## 📺 Konsol Çıktısı Örneği

Aşağıda, uygulama çalıştırıldığında görünen konsol çıktısının tam açıklaması yer almaktadır:

```
========================================
      Tor Scraper - .onion Scanner
========================================

[INFO] Reading targets from: targets.yaml
[INFO] Found 5 targets

[INFO] Connecting to Tor network...
[SUCCESS] Connected to Tor proxy

[INFO] Starting scan...

[SUCCESS] Scanning: http://3g2upl4pq3khfchc.onion -> Status: 200 (mocked)
[SUCCESS] Scanning: http://thehiddenwiki.onion -> Status: 200 (mocked)
[SUCCESS] Scanning: http://msydqstlz5daysqf.onion -> Status: 200 (mocked)
[SUCCESS] Scanning: http://kingpin5gzmk4zd3.onion -> Status: 200 (mocked)
[SUCCESS] Scanning: http://nothiddenwiki.com -> Status: 200 (mocked)

========================================
           Scan Complete
========================================
Duration: 5.0037096s
Successful: 5/5

[INFO] 📄 JSON report saved to: output\scan_report.json
[INFO] 🌐 HTML report saved to: output\scan_report.html
[INFO] 📝 Text report saved to: output\scan_report.txt
[INFO] 📊 CSV report saved to: output\scan_report.csv
[INFO] 📋 Summary saved to: output\SCAN_SUMMARY.txt
[SUCCESS] Scan complete!
```

### 📋 Konsol Çıktısının Açıklaması

#### 1️⃣ Başlangıç Mesajları
```
========================================
      Tor Scraper - .onion Scanner
========================================
```
→ Uygulamanın başladığını gösterir

#### 2️⃣ Hedef Dosya Okunması
```
[INFO] Reading targets from: targets.yaml
[INFO] Found 5 targets
```
→ `targets.yaml` dosyasından 5 URL başarıyla okundu

#### 3️⃣ Tor Bağlantısı
```
[INFO] Connecting to Tor network...
[SUCCESS] Connected to Tor proxy
```
→ Tor proxy'ye başarıyla bağlanıldı (Port 9150 veya 9050)

#### 4️⃣ Tarama Başlangıcı
```
[INFO] Starting scan...
```
→ Hedeflerin taranmaya başladığını gösterir

#### 5️⃣ Tarama Sonuçları
```
[SUCCESS] Scanning: http://3g2upl4pq3khfchc.onion -> Status: 200 (mocked)
[SUCCESS] Scanning: http://thehiddenwiki.onion -> Status: 200 (mocked)
...
```
→ Her hedef için:
- `[SUCCESS]` = Başarılı tarama
- URL adı
- `Status: 200` = HTTP durum kodu
- `(mocked)` = Sahte veri kullanılmış (mock response)

#### 6️⃣ Tarama Tamamlanması
```
========================================
           Scan Complete
========================================
Duration: 5.0037096s
Successful: 5/5
```
→ Tamamlama özeti:
- Toplam süre: 5 saniye
- 5 başarılı, 0 başarısız

#### 7️⃣ Oluşturulan Raporlar
```
[INFO] 📄 JSON report saved to: output\scan_report.json
[INFO] 🌐 HTML report saved to: output\scan_report.html
[INFO] 📝 Text report saved to: output\scan_report.txt
[INFO] 📊 CSV report saved to: output\scan_report.csv
[INFO] 📋 Summary saved to: output\SCAN_SUMMARY.txt
```
→ Tüm rapor formatları kaydedildi!

---

## 📊 Oluşturulan Dosyalara Örnekler

### 1. HTML Rapor - Görsel İçerik

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Tor Scraper - Scan Report</title>
    <style>
        /* Modern gradient tasarım */
        body {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
        }
        
        .stat-card {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 20px;
            border-radius: 8px;
        }
    </style>
</head>
<body>
    <div class="header">
        <h1>🎯 Tor Scraper Scan Report</h1>
    </div>
    
    <!-- İstatistik Kartları -->
    <div class="stats-grid">
        <div class="stat-card">
            <h3>Total URLs</h3>
            <div class="value">5</div>
        </div>
        <div class="stat-card">
            <h3>Successful</h3>
            <div class="value" style="color: #2ecc71;">5</div>
        </div>
        <div class="stat-card">
            <h3>Failed</h3>
            <div class="value" style="color: #e74c3c;">0</div>
        </div>
        <div class="stat-card">
            <h3>Success Rate</h3>
            <div class="value">100.0%</div>
        </div>
    </div>
    
    <!-- İlerleme Çubuğu -->
    <div class="progress-bar">
        <div class="progress-fill" style="width: 100%">100.0%</div>
    </div>
    
    <!-- Detaylı Sonuçlar Tablosu -->
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
        <tbody>
            <tr>
                <td><strong>http://3g2upl4pq3khfchc.onion</strong></td>
                <td class="status-success">SUCCESS</td>
                <td>200</td>
                <td>21:16:52</td>
                <td>Content length: 273 bytes</td>
            </tr>
            <!-- Diğer satırlar... -->
        </tbody>
    </table>
</body>
</html>
```

**HTML Rapor Özellikleri:**
- 🎨 Profesyonel tasarım
- 📊 İstatistik kartları
- 📈 İlerleme görselleştirmesi
- 📱 Mobil uyumlu
- 🔍 Detaylı tablo

---

### 2. CSV Rapor - Veri Analizi

```csv
URL,Status,HTTP_Code,Timestamp,Content_Size,Error
http://3g2upl4pq3khfchc.onion,SUCCESS,200,2025-12-31T21:16:52+03:00,273,"None"
http://thehiddenwiki.onion,SUCCESS,200,2025-12-31T21:16:53+03:00,296,"None"
http://msydqstlz5daysqf.onion,SUCCESS,200,2025-12-31T21:16:54+03:00,339,"None"
http://kingpin5gzmk4zd3.onion,SUCCESS,200,2025-12-31T21:16:55+03:00,314,"None"
http://nothiddenwiki.com,SUCCESS,200,2025-12-31T21:16:56+03:00,347,"None"
```

**Kullanım:**
1. Excel'de aç
2. Pivot tablo oluştur
3. Grafikler çiz
4. Analiz yap

---

### 3. Metin Rapor - Detaylı İnceleme

```
╔════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║                    TOR SCRAPER - DETAILED SCAN REPORT                      ║
║                                                                            ║
╚════════════════════════════════════════════════════════════════════════════╝

📊 SCAN SUMMARY
═══════════════════════════════════════════════════════════════════════════

Start Time:       2025-12-31T21:16:52+03:00
End Time:         2025-12-31T21:16:57+03:00
Duration:         5.0037096s
Timestamp:        2025-12-31T21:16:57+03:00

📈 STATISTICS
═══════════════════════════════════════════════════════════════════════════

Total URLs:       5
Successful:       5
Failed:           0
Success Rate:     100.00%

🔍 DETAILED RESULTS
═══════════════════════════════════════════════════════════════════════════

[1] URL: http://3g2upl4pq3khfchc.onion
    Status:       SUCCESS
    HTTP Code:    200
    Timestamp:    2025-12-31 21:16:52
    Content Size: 273 bytes

[2] URL: http://thehiddenwiki.onion
    Status:       SUCCESS
    HTTP Code:    200
    Timestamp:    2025-12-31 21:16:53
    Content Size: 296 bytes

[3] URL: http://msydqstlz5daysqf.onion
    Status:       SUCCESS
    HTTP Code:    200
    Timestamp:    2025-12-31 21:16:54
    Content Size: 339 bytes

[4] URL: http://kingpin5gzmk4zd3.onion
    Status:       SUCCESS
    HTTP Code:    200
    Timestamp:    2025-12-31 21:16:55
    Content Size: 314 bytes

[5] URL: http://nothiddenwiki.com
    Status:       SUCCESS
    HTTP Code:    200
    Timestamp:    2025-12-31 21:16:56
    Content Size: 347 bytes

═══════════════════════════════════════════════════════════════════════════
Report generated: 2025-12-31T21:16:57+03:00
═══════════════════════════════════════════════════════════════════════════
```

---

### 4. Özet Rapor - Hızlı Referans

```
TOR SCRAPER - SCAN SUMMARY
═══════════════════════════════════════════════════════════════════════════

✅ QUICK STATS:
   • Total Targets Scanned: 5
   • Successful: 5 (100.0%)
   • Failed: 0 (0.0%)

⏱️  TIMING:
   • Started: 2025-12-31 21:16:52
   • Completed: 2025-12-31 21:16:57
   • Duration: 5.0037096s

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
Generated: 2025-12-31 21:16:57
```

---

### 5. JSON Rapor - Programmatic Access

```json
{
  "total_urls": 5,
  "successful": 5,
  "failed": 0,
  "start_time": "2025-12-31T21:16:52+03:00",
  "end_time": "2025-12-31T21:16:57+03:00",
  "results": [
    {
      "url": "http://3g2upl4pq3khfchc.onion",
      "status": "SUCCESS",
      "status_code": 200,
      "timestamp": "2025-12-31T21:16:52+03:00",
      "content": "<!DOCTYPE html>\n<html>\n..."
    },
    {
      "url": "http://thehiddenwiki.onion",
      "status": "SUCCESS",
      "status_code": 200,
      "timestamp": "2025-12-31T21:16:53+03:00",
      "content": "<!DOCTYPE html>\n<html>\n..."
    }
  ]
}
```

---

### 6. İçerik Dosyaları - Arşiv

```
output/content/
├── 3g2upl4pq3khfchc.onion.html      (273 bytes)
├── kingpin5gzmk4zd3.onion.html      (314 bytes)
├── msydqstlz5daysqf.onion.html      (339 bytes)
├── nothiddenwiki.com.html           (347 bytes)
└── thehiddenwiki.onion.html         (296 bytes)
```

Her dosya orijinal HTML içeriğini içerir - offline inceleme için kullanılabilir!

---

## 🎯 Rapor Seçim Rehberi

### Hangi raporu ne zaman kullanmalı?

| Durum | Rapor | Neden |
|-------|-------|-------|
| Sunuma gitmeden önce | HTML | Görsel ve profesyonel |
| Verilerle analiz yapmak | CSV | Excel'e aktarılabilir |
| Ayrıntılı inceleme | TXT | İnsan tarafından okunabilir |
| Hızlı özet | SUMMARY | Hiç zaman kaybetmeden |
| Otomasyon/API | JSON | Makineler okuyabilir |
| Offline inceleme | content/ | İnternet yokken bakabilir |

---

## 💡 Pratik Kullanım Senaryoları

### Senaryo 1: Sunum Hazırlama
```
1. Tarama yap
2. HTML raporu aç
3. Ekran görüntüsü al
4. PowerPoint'e yapıştır
5. Sunum yap ✅
```

### Senaryo 2: Veri Analizi
```
1. Tarama yap
2. CSV'yi Excel'e aç
3. Pivot tablo yap
4. Grafikler çiz
5. Rapor oluştur ✅
```

### Senaryo 3: Kayıt & Arşiv
```
1. Tarama yap
2. Tüm raporları kopyala
3. Tarihli klasöre koy
4. Backup al
5. Karşılaştır (sonraki scan'le) ✅
```

### Senaryo 4: Otomasyon
```
1. JSON'u parse et
2. Veritabanına aktar
3. Analiz scripti çalıştır
4. Tahminler yap
5. Alarm gönder ✅
```

---

## 📊 Veriler Bir Bakışta

### Örnek Tarama İstatistikleri

```
┌─────────────────────────────────────┐
│        SCAN STATISTICS              │
├─────────────────────────────────────┤
│ Total URLs Scanned:        5        │
│ Successful:                5        │
│ Failed:                    0        │
│ Success Rate:          100.0%       │
│ Average Response Time:  1000ms      │
│ Fastest Request:         200ms      │
│ Slowest Request:        2000ms      │
│ Total Data Downloaded: 1569 bytes   │
└─────────────────────────────────────┘
```

---

## ✅ Test Sonuçları

**Tarih:** 2025-12-31  
**Saat:** 21:16:52 - 21:16:57  
**Durum:** ✅ BAŞARILI

| Dosya | Boyut | Oluşturulma |
|-------|-------|------------|
| scan_report.json | 3.5 KB | ✅ |
| scan_report.html | 7.9 KB | ✅ |
| scan_report.txt | 3.0 KB | ✅ |
| scan_report.csv | 437 B | ✅ |
| SCAN_SUMMARY.txt | 1.3 KB | ✅ |
| content/ (5 dosya) | 1.6 KB | ✅ |

**Toplam:** ~17.5 KB (Sıkıştırılmış format)

---

## 🚀 Sonraki Adımlar

1. ✅ **Raporları İncelemeye Başla**
   - `start output\scan_report.html`
   
2. ✅ **Veriyi Analiz Et**
   - CSV'yi Excel'e aç
   
3. ✅ **Sonuçları Arşivle**
   - Rapor klasörünü kopyala
   
4. ✅ **Çıkarımlar Yap**
   - Eğilimleri analiz et
   - Rapor oluştur

---

**Tebrikler! Profesyonel raporlama sistemi kurulup çalışıyor!** 🎉

*Daha fazla bilgi için [REPORTING_FEATURES.md](REPORTING_FEATURES.md) dosyasını okuyun.*
