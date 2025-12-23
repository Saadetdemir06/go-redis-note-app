# 📝 go-redis-note-app

**High-Performance Note-Taking System** | Built with **Go** & **Redis** | **CI/CD** via GitHub Actions | **Performance** validated with k6.

Bu proje, modern yazılım mimarisi prensipleri ve yüksek performanslı veri saklama çözümleri (in-memory) kullanılarak geliştirilmiş bir not alma uygulamasıdır.

## 🚀 Öne Çıkan Teknik Özellikler
- **Backend:** Go (Golang) ile geliştirilmiş yüksek performanslı REST API.
- **Veritabanı:** Redis (In-memory data structure store) ile ultra düşük gecikme süreli veri erişimi.
- **Konteynerleştirme:** Docker ve Docker Compose ile izole çalışma ortamı.
- **Otomasyon (CI/CD):** GitHub Actions ile her push işleminde otomatik birim testler.
- **Yük Testi:** k6 aracı ile sistemin eşzamanlı kullanıcı altındaki kararlılığının ölçümü.

## 📊 Performans Ölçüm Sonuçları (k6 Bulguları)
Sistem üzerinde yapılan k6 performans testlerinde elde edilen veriler şunlardır:
- **Ortalama Yanıt Süresi (Avg Latency):** 1.84ms (Redis sayesinde hedeflenen 200ms'nin çok altında).
- **Başarı Oranı (Success Rate):** %100 (10 eşzamanlı sanal kullanıcı altında sıfır hata).
- **İşlem Kapasitesi:** Saniyede ortalama 10 başarılı istek (10 VUs).

## 🛠️ Kurulum ve Çalıştırma
Sistemi kendi yerel ortamınızda ayağa kaldırmak için:

1. Projeyi klonlayın:
   ```bash
   git clone [https://github.com/Saadetdemir06/go-redis-note-app.git](https://github.com/Saadetdemir06/go-redis-note-app.git)
