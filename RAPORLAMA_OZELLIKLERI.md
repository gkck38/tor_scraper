# 🎯 Raporlama Özellikleri - Türkçe Rehberi

## Genel Bakış

Tor Scraper artık kapsamlı raporlama yeteneklerine sahiptir. Scan sonuçlarını çeşitli formatlarda inceleyebilirsiniz.

## 📊 Oluşturulan Rapor Dosyaları

### 1️⃣ **scan_report.json**
Bilgisayar tarafından okunabilir format. Diğer araçlarla entegrasyon için idealdir.

```json
{
  "total_urls": 5,
  "successful": 5,
  "failed": 0,
  "results": [...]
}
```

### 2️⃣ **scan_report.html** ⭐ YENİ
İnteraktif web raporu! Tarayıcıda açın ve görsel analiz yapın.

**Özellikler:**
- 📈 İstatistik kartları
- 📊 Başarı oranı çubuğu
- 📋 Detaylı sonuç tablosu
- 🎨 Profesyonel tasarım
- 📱 Mobil uyumlu

**Açmak için:**
```bash
# Windows
start output\scan_report.html

# Linux/Mac
open output/scan_report.html
```

### 3️⃣ **scan_report.txt** ✨ YENILENDI
Düzgün formatlı metin raporu detaylı analiz için.

```
╔════════════════════════════════════════════════════════════╗
║          TOR SCRAPER - DETAILED SCAN REPORT                ║
╚════════════════════════════════════════════════════════════╝

📈 STATISTICS:
├─ Total URLs: 5
├─ Successful: 5 (100%)
└─ Failed: 0 (0%)

⏱️ TIMING:
├─ Start: 2025-12-31 21:16:52
├─ End: 2025-12-31 21:16:57
└─ Duration: 5 seconds
```

### 4️⃣ **scan_report.csv** ⭐ YENİ
Excel/Google Sheets'e aktarmak için mükemmel!

```
URL,Status,HTTP_Code,Content_Size,Error
http://example.onion,SUCCESS,200,1024,"None"
...
```

**Kullanım:**
```
1. CSV dosyasını aç
2. Excel/Google Sheets'e aktar
3. Pivot tablolar oluştur
4. Grafikler çiz
```

### 5️⃣ **SCAN_SUMMARY.txt** ⭐ YENİ
Hızlı özet ve sonraki adımlar

```
✅ QUICK STATS:
   • Total Targets: 5
   • Successful: 5 (100.0%)
   • Failed: 0 (0.0%)

📁 OUTPUT FILES:
   • scan_report.json
   • scan_report.html
   • scan_report.txt
   • scan_report.csv
   • content/ (Bireysel dosyalar)
```

### 6️⃣ **content/** (Klasör)
Her başarılı tarama için ayrı HTML dosya kaydedilir.

## 📊 Sağlanan İstatistikler

✅ **Temel Metrikler:**
- Toplam URL sayısı
- Başarılı taramalar
- Başarısız taramalar
- Başarı oranı (%)

⏱️ **Zaman Bilgileri:**
- Başlangıç tarihi
- Bitiş tarihi
- Toplam süre
- Her request'in zamanı

🔗 **URL Başına Detaylar:**
- Tam URL
- Durum (SUCCESS/FAILED/ERROR)
- HTTP durum kodu
- Request zamanı
- İçerik boyutu (bytes)
- Hata mesajı (varsa)

## 🎯 Kullanım Senaryoları

### 1. Uyumluluk & Denetim
```bash
# Sunuma hazır HTML raporunu aç
start output\scan_report.html

# Veritabanına import et
import output/scan_report.csv into Excel
```

### 2. Veri Analizi
```bash
# CSV'yi spreadsheet'e aktar
# Pivot tablolar oluştur
# Grafikler çiz
```

### 3. Tehdit İstihbaratı
```bash
# JSON'u otomatik işleme al
# İçerik dosyalarını analiz et
# Zaman içinde değişimleri takip et
```

### 4. Takım Raporlaması
```bash
# HTML raporu e-postayla gönder
# SCAN_SUMMARY.txt ile sunum yap
# Tam raporları arşivle
```

## 🚀 Örnek İş Akışı

```bash
# 1. Taramayı çalıştır
.\tor-scraper.exe targets.yaml

# 2. Hızlı özeti gör
type output\SCAN_SUMMARY.txt

# 3. Detaylı analiz (Metin)
type output\scan_report.txt

# 4. İnteraktif rapor (HTML)
start output\scan_report.html

# 5. Excel'e aktar (Windows)
start output\scan_report.csv

# 6. İçerik dosyalarını kontrol et
dir output\content\
```

## 🎨 HTML Rapor Özellikleri

### Göz Alıcı Tasarım
- 🟣 Modern gradient arka plan
- 📊 İstatistik kartları
- 📈 Başarı oranı görselleştirmesi
- 🌈 Renkli durum göstergeleri

### Duyarlı Layout
- 🖥️ Masaüstü: 4 sütun grid
- 📱 Tablet: 2 sütun layout
- 📱 Mobil: 1 sütun layout

### Interaktif Tablo
- Sıralanabilir kolonlar
- Hover efektleri
- Tam URL görünümleri
- HTTP kodu renk kodlaması

## 📈 Sonuçlar

✅ **Başarılı (SUCCESS)**
- HTTP 200 yanıtı
- İçerik başarıyla alındı

❌ **Başarısız (FAILED)**
- Bağlantı hatası
- Timeout
- Ağ sorunu

⚠️ **Kısmi (PARTIAL)**
- İçerik kısmen alındı
- Dosya boyutu limiti aşıldı

🔴 **Hata (ERROR)**
- Request oluşturma hatası
- Yapılandırma sorunu

## 💡 İpuçları ve Püf Noktaları

### Farklı Taramaları Arşivle
```bash
mkdir reports/2025-12-31
.\tor-scraper.exe targets.yaml reports/2025-12-31
```

### Raporları Karşılaştır
```bash
# İki tarama sonucunu karşılaştır
fc output/scan_report.csv output2/scan_report.csv
```

### Raporu E-postayla Gönder
```bash
# HTML raporunu gönder
# CSV'yi import et
```

### Otomatik Raporlama
```bash
# Günlük tarama yap
# Raporu kaydet
# Ekibe gönder
```

## 🔐 Güvenlik Notları

⚠️ **Önemli:**
- Raporlar hassas veri içerir
- Güvenli yerlerde saklayın
- Güvenilmez kanallarda paylaşmayın
- Arşivleri şifreleyin
- İçerik dosyalarında kötü amaçlı kod olabilir!

## 📊 Dosya Yapısı

```
output/
├── scan_report.json        (JSON verisi)
├── scan_report.html        (Web raporu)
├── scan_report.txt         (Detaylı rapor)
├── scan_report.csv         (Tablo verisi)
├── SCAN_SUMMARY.txt        (Özet)
└── content/                (İçerik arşivi)
    ├── 3g2upl4pq3khfchc.onion.html
    ├── thehiddenwiki.onion.html
    └── ...
```

## 🔄 Sürüm Tarihi

### v1.0 - İlk Raporlama Sistemi
- ✅ JSON export
- ✅ Metin log dosyası
- ✅ İçerik dosyaları

### v1.1 - Gelişmiş Raporlama ⭐ GÜNCEL
- ✅ İnteraktif HTML raporlar
- ✅ CSV veri analizi
- ✅ Yönetim özeti
- ✅ Geliştirilmiş metin formatı
- ✅ İlerleme görselleştirmesi

---

**Sorularınız için README.md dosyasına bakın.**
