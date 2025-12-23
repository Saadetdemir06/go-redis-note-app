package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAnaSayfa(t *testing.T) {
	// 1. Sanal bir GET isteği oluşturuyoruz
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 2. Yanıtı kaydetmek için sanal bir kaydedici (recorder) oluşturuyoruz
	rr := httptest.NewRecorder()
	
	// 3. Test edilecek handler fonksiyonu (Gerekirse main.go'daki isme göre güncelle)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Not Defteri Backend Çalışıyor 🚀"))
	})

	// 4. İsteği handler'a gönderiyoruz
	handler.ServeHTTP(rr, req)

	// 5. Durum kodunun 200 (OK) olup olmadığını kontrol ediyoruz
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Yanlış durum kodu: alınan %v beklenen %v", status, http.StatusOK)
	}

	// 6. Yanıt içeriğini kontrol ediyoruz
	expected := "Not Defteri Backend Çalışıyor 🚀"
	if rr.Body.String() != expected {
		t.Errorf("Beklenmeyen yanıt: alınan %v beklenen %v", rr.Body.String(), expected)
	}
}