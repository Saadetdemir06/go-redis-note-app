package main

import (
	"net/http"
	"testing"
)

// Sunucunun ayakta olup olmadığını test eder
func TestAnaSayfa(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/")
	if err != nil {
		t.Fatalf("Sunucuya bağlanılamadı: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Beklenen durum kodu 200, alınan: %d", resp.StatusCode)
	}
}