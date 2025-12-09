package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type SoruAnalizi struct {
	Ders string `json:"ders"`
	Konu string `json:"konu"`
	Hata string `json:"hata,omitempty"` // Eğer soru değilse dolu gelir
}

func main() {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 2. Modeli Seç ve Konfigüre Et (JSON Modu Önemli!)
	model := client.GenerativeModel("gemini-2.5-flash")
	model.ResponseMIMEType = "application/json" // Gemini'yi JSON dönmeye zorlar

	// 3. Resmi Yükle (Gerçek senaryoda bu HTTP request'ten gelen []byte olacak)
	imgURL := "https://example.com/path/to/turkce.jpeg"
	imgResp, err := http.Get(imgURL)
	if err != nil {
		log.Fatal("Resim indirilemedi:", err)
	}
	defer imgResp.Body.Close()

	imgData, err := io.ReadAll(imgResp.Body)
	if err != nil {
		log.Fatal("Resim okunamadı:", err)
	}

	configTalimati := `
		Sen bir sınav sorusu analiz motorusun.
		Görevin: Sana verilen görseli analiz edip JSON formatında yanıt döndürmek.

		Kurallar:
		1. Eğer görsel bir ders çalışma sorusu (test, klasik vb.) İÇERİYORSA:
		- Hangi ders oldugunu analiz et ve ders alanını doldur
  - Hangi konuya ait olduğunu analiz et.
		   - "konu" alanını doldur, "hata" alanını boş bırak.

		2. Eğer görsel bir soru İÇERMİYORSA (örneğin manzara, selfie, anlamsız çizim, yemek fotoğrafı vb.):
		   - "konu" alanını boş bırak ("").
		   - "hata" alanına "Görselde bir soru tespit edilemedi." yaz.

		3. Çıktın SADECE şu JSON formatında olmalı:
		   {"ders": "...", "konu": "...", "hata": "..."}
		`
	// ---------------------

	model.SystemInstruction = genai.NewUserContent(genai.Text(configTalimati))
	model.ResponseMIMEType = "application/json"

	// 5. İsteği Gönder (Text + Image)
	prompt := genai.Text("bu bir turkce sorusu. Bu görseldeki soruyu analiz et.")
	resp, err := model.GenerateContent(ctx,
		genai.Text(prompt),
		genai.ImageData("jpeg", imgData)) // Resim formatı (png, jpeg, webp)
	if err != nil {
		log.Fatal(err)
	}

	// 6. Sonucu İşle
	if len(resp.Candidates) > 0 {
		part := resp.Candidates[0].Content.Parts[0]
		jsonString := fmt.Sprintf("%v", part)

		var analiz SoruAnalizi

		// Gelen JSON'u senin struct'ına çeviriyoruz (Unmarshal)
		if err := json.Unmarshal([]byte(jsonString), &analiz); err != nil {
			fmt.Println("JSON Parse Hatası:", err)
			fmt.Println("Gelen Ham Veri:", jsonString)
		} else {
			// SONUÇ BAŞARILI
			// Eğer hata yoksa sadece konuyu, hata varsa hatayı basabiliriz.
			fmt.Printf("--- Analiz Sonucu ---\n")
			fmt.Printf("Konu: %s\n", analiz.Konu)
			fmt.Printf("Hata: %s\n", analiz.Hata)

			// Tüm struct'ı JSON olarak görelim
			finalJson, _ := json.MarshalIndent(analiz, "", "  ")
			fmt.Println("\nJSON Çıktısı:\n", string(finalJson))
		}
	}
}

/*
 * 	filename := "index.png"
	err := orcgen.ConvertHTML(orcgen.PNG, getHTML(), getName(filename))
	if err == nil {
		fmt.Printf("%s generated succesfully\n", filename)
	}
 *
*/

func getHTML() []byte {
	file := filepath.Join(getBasepath(), "./index.html")
	html, _ := os.ReadFile(file)

	return html
}

func getName(name string) string {
	return filepath.Join(getBasepath(), "testdata", name)
}

func getBasepath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
