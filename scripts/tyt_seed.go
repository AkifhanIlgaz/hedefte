package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/pkg/db"
)

func main() {
	// Config yükle
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config yüklenemedi: %v", err)
	}

	// MongoDB bağlantısı
	client, err := db.ConnectMongo(cfg.Mongo)
	if err != nil {
		log.Fatalf("MongoDB bağlantısı kurulamadı: %v", err)
	}
	defer client.Disconnect(context.Background())

	database := client.Database("hedefte") // Database adını kendi projenize göre ayarlayın
	collection := database.Collection("tyt_analysis")

	uid := "996609ad-5046-45dc-8797-f7fd8663d3e6"

	// Test verileri - 9 subject ile
	testData := []models.ExamAnalysisDB{
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2025, 1, 5, 10, 30, 0, 0, time.UTC),
				Name: "Haftalık Deneme 1",
				Subjects: []models.Subject{
					{
						Name: "Türkçe", Correct: 30, Wrong: 8, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Sözcükte Anlam", Mistakes: 3},
							{Topic: "Paragraf", Mistakes: 3},
							{Topic: "Edebiyat", Mistakes: 2},
						},
					},
					{
						Name: "Tarih", Correct: 4, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Yakın Çağ", Mistakes: 1},
						},
					},
					{
						Name: "Coğrafya", Correct: 3, Wrong: 2, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Fiziki Coğrafya", Mistakes: 2},
						},
					},
					{
						Name: "Felsefe", Correct: 4, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Bilgi Felsefesi", Mistakes: 1},
						},
					},
					{
						Name: "Din", Correct: 4, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Tefsir", Mistakes: 1},
						},
					},
					{
						Name: "Matematik", Correct: 28, Wrong: 10, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Fonksiyonlar", Mistakes: 4},
							{Topic: "Geometri", Mistakes: 3},
							{Topic: "Sayılar", Mistakes: 3},
						},
					},
					{
						Name: "Fizik", Correct: 6, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Mekanik", Mistakes: 1},
						},
					},
					{
						Name: "Kimya", Correct: 6, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Organik Kimya", Mistakes: 1},
						},
					},
					{
						Name: "Biyoloji", Correct: 5, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Hücre", Mistakes: 1},
						},
					},
				},
			},
			TotalNet: 102.5,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2025, 1, 3, 14, 15, 0, 0, time.UTC),
				Name: "Aralık Ayı Sonu",
				Subjects: []models.Subject{
					{
						Name: "Türkçe", Correct: 27, Wrong: 12, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Cümlede Anlam", Mistakes: 5},
							{Topic: "Dil Bilgisi", Mistakes: 4},
							{Topic: "Paragraf", Mistakes: 3},
						},
					},
					{
						Name: "Tarih", Correct: 2, Wrong: 3, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Orta Çağ", Mistakes: 2},
							{Topic: "İlk Çağ", Mistakes: 1},
						},
					},
					{
						Name: "Coğrafya", Correct: 1, Wrong: 4, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Beşeri Coğrafya", Mistakes: 3},
							{Topic: "Harita Bilgisi", Mistakes: 1},
						},
					},
					{
						Name: "Felsefe", Correct: 2, Wrong: 3, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Modern Felsefe", Mistakes: 2},
							{Topic: "Bilgi Felsefesi", Mistakes: 1},
						},
					},
					{
						Name: "Din", Correct: 3, Wrong: 2, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Dinler Tarihi", Mistakes: 2},
						},
					},
					{
						Name: "Matematik", Correct: 22, Wrong: 15, Empty: 3,
						TopicMistakes: []models.TopicMistake{
							{Topic: "İstatistik", Mistakes: 6},
							{Topic: "Fonksiyonlar", Mistakes: 5},
							{Topic: "Geometri", Mistakes: 4},
						},
					},
					{
						Name: "Fizik", Correct: 4, Wrong: 3, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Optik", Mistakes: 2},
							{Topic: "Elektrik", Mistakes: 1},
						},
					},
					{
						Name: "Kimya", Correct: 4, Wrong: 3, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Asit-Baz", Mistakes: 2},
							{Topic: "Kimyasal Tepkime", Mistakes: 1},
						},
					},
					{
						Name: "Biyoloji", Correct: 3, Wrong: 3, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Ekosistem", Mistakes: 2},
							{Topic: "Genetik", Mistakes: 1},
						},
					},
				},
			},
			TotalNet: 78.5,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC),
				Name: "Yılbaşı Denemesi",
				Subjects: []models.Subject{
					{
						Name: "Türkçe", Correct: 35, Wrong: 4, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Edebiyat", Mistakes: 2},
							{Topic: "Sözcükte Anlam", Mistakes: 2},
						},
					},
					{
						Name: "Tarih", Correct: 5, Wrong: 0, Empty: 0,
						TopicMistakes: []models.TopicMistake{},
					},
					{
						Name: "Coğrafya", Correct: 4, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Harita Bilgisi", Mistakes: 1},
						},
					},
					{
						Name: "Felsefe", Correct: 5, Wrong: 0, Empty: 0,
						TopicMistakes: []models.TopicMistake{},
					},
					{
						Name: "Din", Correct: 5, Wrong: 0, Empty: 0,
						TopicMistakes: []models.TopicMistake{},
					},
					{
						Name: "Matematik", Correct: 32, Wrong: 6, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Çember", Mistakes: 3},
							{Topic: "Sayılar", Mistakes: 3},
						},
					},
					{
						Name: "Fizik", Correct: 7, Wrong: 0, Empty: 0,
						TopicMistakes: []models.TopicMistake{},
					},
					{
						Name: "Kimya", Correct: 6, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Asit-Baz", Mistakes: 1},
						},
					},
					{
						Name: "Biyoloji", Correct: 6, Wrong: 0, Empty: 0,
						TopicMistakes: []models.TopicMistake{},
					},
				},
			},
			TotalNet: 118.5,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2024, 12, 28, 16, 45, 0, 0, time.UTC),
				Name: "Kış Kampı Deneme",
				Subjects: []models.Subject{
					{
						Name: "Türkçe", Correct: 20, Wrong: 18, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Paragraf", Mistakes: 8},
							{Topic: "Cümlede Anlam", Mistakes: 6},
							{Topic: "Dil Bilgisi", Mistakes: 4},
						},
					},
					{
						Name: "Tarih", Correct: 1, Wrong: 4, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Orta Çağ", Mistakes: 2},
							{Topic: "Yakın Çağ", Mistakes: 2},
						},
					},
					{
						Name: "Coğrafya", Correct: 0, Wrong: 5, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Beşeri Coğrafya", Mistakes: 3},
							{Topic: "Fiziki Coğrafya", Mistakes: 2},
						},
					},
					{
						Name: "Felsefe", Correct: 1, Wrong: 4, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Bilgi Felsefesi", Mistakes: 2},
							{Topic: "Modern Felsefe", Mistakes: 2},
						},
					},
					{
						Name: "Din", Correct: 2, Wrong: 3, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Tefsir", Mistakes: 2},
							{Topic: "Dinler Tarihi", Mistakes: 1},
						},
					},
					{
						Name: "Matematik", Correct: 18, Wrong: 20, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Geometri", Mistakes: 8},
							{Topic: "İstatistik", Mistakes: 7},
							{Topic: "Fonksiyonlar", Mistakes: 5},
						},
					},
					{
						Name: "Fizik", Correct: 2, Wrong: 5, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Mekanik", Mistakes: 3},
							{Topic: "Elektrik", Mistakes: 2},
						},
					},
					{
						Name: "Kimya", Correct: 2, Wrong: 5, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Organik Kimya", Mistakes: 3},
							{Topic: "Asit-Baz", Mistakes: 2},
						},
					},
					{
						Name: "Biyoloji", Correct: 1, Wrong: 5, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Ekosistem", Mistakes: 3},
							{Topic: "Hücre", Mistakes: 2},
						},
					},
				},
			},
			TotalNet: 58.75,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2024, 12, 25, 11, 20, 0, 0, time.UTC),
				Name: "Noel Tatili Çalışması",
				Subjects: []models.Subject{
					{
						Name: "Türkçe", Correct: 33, Wrong: 6, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Edebiyat", Mistakes: 3},
							{Topic: "Sözcükte Anlam", Mistakes: 3},
						},
					},
					{
						Name: "Tarih", Correct: 4, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "İlk Çağ", Mistakes: 1},
						},
					},
					{
						Name: "Coğrafya", Correct: 3, Wrong: 2, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Harita Bilgisi", Mistakes: 2},
						},
					},
					{
						Name: "Felsefe", Correct: 4, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Modern Felsefe", Mistakes: 1},
						},
					},
					{
						Name: "Din", Correct: 4, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Dinler Tarihi", Mistakes: 1},
						},
					},
					{
						Name: "Matematik", Correct: 30, Wrong: 8, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Sayılar", Mistakes: 4},
							{Topic: "Çember", Mistakes: 4},
						},
					},
					{
						Name: "Fizik", Correct: 6, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Optik", Mistakes: 1},
						},
					},
					{
						Name: "Kimya", Correct: 6, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Organik Kimya", Mistakes: 1},
						},
					},
					{
						Name: "Biyoloji", Correct: 5, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Ekosistem", Mistakes: 1},
						},
					},
				},
			},
			TotalNet: 112.75,
		},
	}

	// Verileri ekle
	for i, data := range testData {
		_, err := collection.InsertOne(context.Background(), data)
		if err != nil {
			log.Printf("Veri %d eklenemedi: %v", i+1, err)
			continue
		}
		fmt.Printf("Veri %d başarıyla eklendi: %s (Net: %.2f)\n", i+1, data.Name, data.TotalNet)
	}

	fmt.Printf("\n✅ Toplam %d test verisi başarıyla eklendi!\n", len(testData))
	fmt.Printf("📊 Tüm verilerde 9 subject bulunmaktadır:\n")
	fmt.Println("   • Türkçe, Tarih, Coğrafya, Felsefe, Din")
	fmt.Println("   • Matematik, Fizik, Kimya, Biyoloji")
}
