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

	database := client.Database("hedefte")
	collection := database.Collection("ayt_analysis")

	uid := "996609ad-5046-45dc-8797-f7fd8663d3e6"

	// EA (Eşit Ağırlık) Test Verileri
	eaTestData := []models.ExamAnalysisDB{
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2025, 1, 6, 10, 0, 0, 0, time.UTC),
				Name: "EA Deneme 1",
				Subjects: []models.Subject{
					{
						Name: "Türk Dili ve Edebiyatı", Correct: 20, Wrong: 3, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Divan Edebiyatı", Mistakes: 2},
							{Topic: "Metin Analizi", Mistakes: 1},
						},
					},
					{
						Name: "Tarih", Correct: 7, Wrong: 2, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Osmanlı Tarihi", Mistakes: 1},
							{Topic: "Cumhuriyet Dönemi", Mistakes: 1},
						},
					},
					{
						Name: "Coğrafya", Correct: 4, Wrong: 1, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Türkiye Coğrafyası", Mistakes: 1},
						},
					},
					{
						Name: "Matematik", Correct: 32, Wrong: 6, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "İntegral", Mistakes: 3},
							{Topic: "Logaritma", Mistakes: 2},
							{Topic: "Trigonometri", Mistakes: 1},
						},
					},
				},
			},
			TotalNet: 75.25,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2025, 1, 4, 14, 30, 0, 0, time.UTC),
				Name: "EA Deneme 2",
				Subjects: []models.Subject{
					{
						Name: "Türk Dili ve Edebiyatı", Correct: 18, Wrong: 5, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Modern Edebiyat", Mistakes: 3},
							{Topic: "Dil Bilgisi", Mistakes: 2},
						},
					},
					{
						Name: "Tarih", Correct: 5, Wrong: 4, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "İnkilap Tarihi", Mistakes: 2},
							{Topic: "Osmanlı Tarihi", Mistakes: 2},
						},
					},
					{
						Name: "Coğrafya", Correct: 3, Wrong: 2, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Beşeri Coğrafya", Mistakes: 2},
						},
					},
					{
						Name: "Matematik", Correct: 28, Wrong: 10, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Limit", Mistakes: 4},
							{Topic: "İntegral", Mistakes: 3},
							{Topic: "Trigonometri", Mistakes: 3},
						},
					},
				},
			},
			TotalNet: 65.0,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2025, 1, 2, 11, 15, 0, 0, time.UTC),
				Name: "EA Deneme 3",
				Subjects: []models.Subject{
					{
						Name: "Türk Dili ve Edebiyatı", Correct: 22, Wrong: 2, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Halk Edebiyatı", Mistakes: 1},
							{Topic: "Metin Analizi", Mistakes: 1},
						},
					},
					{
						Name: "Tarih", Correct: 8, Wrong: 1, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Cumhuriyet Dönemi", Mistakes: 1},
						},
					},
					{
						Name: "Coğrafya", Correct: 5, Wrong: 1, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "İklim", Mistakes: 1},
						},
					},
					{
						Name: "Matematik", Correct: 35, Wrong: 4, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Logaritma", Mistakes: 2},
							{Topic: "Türev", Mistakes: 2},
						},
					},
				},
			},
			TotalNet: 82.25,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2024, 12, 30, 9, 45, 0, 0, time.UTC),
				Name: "EA Deneme 4",
				Subjects: []models.Subject{
					{
						Name: "Türk Dili ve Edebiyatı", Correct: 15, Wrong: 8, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Dil Bilgisi", Mistakes: 4},
							{Topic: "Divan Edebiyatı", Mistakes: 3},
							{Topic: "Modern Edebiyat", Mistakes: 1},
						},
					},
					{
						Name: "Tarih", Correct: 4, Wrong: 5, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Osmanlı Tarihi", Mistakes: 3},
							{Topic: "İnkilap Tarihi", Mistakes: 2},
						},
					},
					{
						Name: "Coğrafya", Correct: 2, Wrong: 3, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Türkiye Coğrafyası", Mistakes: 2},
							{Topic: "Beşeri Coğrafya", Mistakes: 1},
						},
					},
					{
						Name: "Matematik", Correct: 20, Wrong: 18, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "İntegral", Mistakes: 7},
							{Topic: "Limit", Mistakes: 6},
							{Topic: "Türev", Mistakes: 5},
						},
					},
				},
			},
			TotalNet: 50.25,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2024, 12, 27, 16, 20, 0, 0, time.UTC),
				Name: "EA Deneme 5",
				Subjects: []models.Subject{
					{
						Name: "Türk Dili ve Edebiyatı", Correct: 21, Wrong: 2, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Halk Edebiyatı", Mistakes: 1},
							{Topic: "Metin Analizi", Mistakes: 1},
						},
					},
					{
						Name: "Tarih", Correct: 7, Wrong: 2, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Cumhuriyet Dönemi", Mistakes: 2},
						},
					},
					{
						Name: "Coğrafya", Correct: 4, Wrong: 2, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "İklim", Mistakes: 1},
							{Topic: "Beşeri Coğrafya", Mistakes: 1},
						},
					},
					{
						Name: "Matematik", Correct: 30, Wrong: 8, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Trigonometri", Mistakes: 3},
							{Topic: "Logaritma", Mistakes: 3},
							{Topic: "Türev", Mistakes: 2},
						},
					},
				},
			},
			TotalNet: 73.5,
		},
	}

	// MF (Matematik-Fen) Test Verileri
	mfTestData := []models.ExamAnalysisDB{
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2025, 1, 7, 13, 30, 0, 0, time.UTC),
				Name: "MF Deneme 1",
				Subjects: []models.Subject{
					{
						Name: "Matematik", Correct: 35, Wrong: 4, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "İntegral", Mistakes: 2},
							{Topic: "Türev", Mistakes: 2},
						},
					},
					{
						Name: "Fizik", Correct: 10, Wrong: 3, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Elektromanyetik", Mistakes: 2},
							{Topic: "Modern Fizik", Mistakes: 1},
						},
					},
					{
						Name: "Kimya", Correct: 11, Wrong: 2, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Organik Kimya", Mistakes: 1},
							{Topic: "Elektrokimya", Mistakes: 1},
						},
					},
					{
						Name: "Biyoloji", Correct: 9, Wrong: 3, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Genetik", Mistakes: 2},
							{Topic: "Sistem Fizyolojisi", Mistakes: 1},
						},
					},
				},
			},
			TotalNet: 78.25,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2025, 1, 5, 15, 0, 0, 0, time.UTC),
				Name: "MF Deneme 2",
				Subjects: []models.Subject{
					{
						Name: "Matematik", Correct: 30, Wrong: 8, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Limit", Mistakes: 4},
							{Topic: "İntegral", Mistakes: 3},
							{Topic: "Trigonometri", Mistakes: 1},
						},
					},
					{
						Name: "Fizik", Correct: 8, Wrong: 5, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Dalgalar", Mistakes: 3},
							{Topic: "Elektrik", Mistakes: 2},
						},
					},
					{
						Name: "Kimya", Correct: 9, Wrong: 4, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Kimyasal Denge", Mistakes: 2},
							{Topic: "Organik Kimya", Mistakes: 2},
						},
					},
					{
						Name: "Biyoloji", Correct: 7, Wrong: 5, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Ekoloji", Mistakes: 3},
							{Topic: "Genetik", Mistakes: 2},
						},
					},
				},
			},
			TotalNet: 65.5,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2025, 1, 3, 10, 15, 0, 0, time.UTC),
				Name: "MF Deneme 3",
				Subjects: []models.Subject{
					{
						Name: "Matematik", Correct: 38, Wrong: 2, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Logaritma", Mistakes: 1},
							{Topic: "Türev", Mistakes: 1},
						},
					},
					{
						Name: "Fizik", Correct: 12, Wrong: 2, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Modern Fizik", Mistakes: 1},
							{Topic: "Elektromanyetik", Mistakes: 1},
						},
					},
					{
						Name: "Kimya", Correct: 12, Wrong: 2, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Elektrokimya", Mistakes: 2},
						},
					},
					{
						Name: "Biyoloji", Correct: 11, Wrong: 2, Empty: 0,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Sistem Fizyolojisi", Mistakes: 2},
						},
					},
				},
			},
			TotalNet: 88.0,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2024, 12, 29, 8, 45, 0, 0, time.UTC),
				Name: "MF Deneme 4",
				Subjects: []models.Subject{
					{
						Name: "Matematik", Correct: 22, Wrong: 16, Empty: 2,
						TopicMistakes: []models.TopicMistake{
							{Topic: "İntegral", Mistakes: 6},
							{Topic: "Limit", Mistakes: 5},
							{Topic: "Türev", Mistakes: 3},
							{Topic: "Trigonometri", Mistakes: 2},
						},
					},
					{
						Name: "Fizik", Correct: 5, Wrong: 8, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Dalgalar", Mistakes: 4},
							{Topic: "Elektrik", Mistakes: 3},
							{Topic: "Modern Fizik", Mistakes: 1},
						},
					},
					{
						Name: "Kimya", Correct: 6, Wrong: 7, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Organik Kimya", Mistakes: 4},
							{Topic: "Kimyasal Denge", Mistakes: 3},
						},
					},
					{
						Name: "Biyoloji", Correct: 4, Wrong: 8, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Ekoloji", Mistakes: 4},
							{Topic: "Genetik", Mistakes: 3},
							{Topic: "Sistem Fizyolojisi", Mistakes: 1},
						},
					},
				},
			},
			TotalNet: 46.25,
		},
		{
			Uid: uid,
			ExamAnalysisRequest: models.ExamAnalysisRequest{
				Date: time.Date(2024, 12, 26, 14, 10, 0, 0, time.UTC),
				Name: "MF Deneme 5",
				Subjects: []models.Subject{
					{
						Name: "Matematik", Correct: 33, Wrong: 6, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Trigonometri", Mistakes: 3},
							{Topic: "Logaritma", Mistakes: 2},
							{Topic: "İntegral", Mistakes: 1},
						},
					},
					{
						Name: "Fizik", Correct: 10, Wrong: 3, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Elektromanyetik", Mistakes: 2},
							{Topic: "Dalgalar", Mistakes: 1},
						},
					},
					{
						Name: "Kimya", Correct: 10, Wrong: 3, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Elektrokimya", Mistakes: 2},
							{Topic: "Organik Kimya", Mistakes: 1},
						},
					},
					{
						Name: "Biyoloji", Correct: 10, Wrong: 2, Empty: 1,
						TopicMistakes: []models.TopicMistake{
							{Topic: "Genetik", Mistakes: 1},
							{Topic: "Sistem Fizyolojisi", Mistakes: 1},
						},
					},
				},
			},
			TotalNet: 76.5,
		},
	}

	// Tüm test verilerini birleştir
	allTestData := append(eaTestData, mfTestData...)

	// Verileri ekle
	for i, data := range allTestData {
		_, err := collection.InsertOne(context.Background(), data)
		if err != nil {
			log.Printf("Veri %d eklenemedi: %v", i+1, err)
			continue
		}
		fmt.Printf("Veri %d başarıyla eklendi: %s (Net: %.2f)\n", i+1, data.Name, data.TotalNet)
	}

	fmt.Printf("\n✅ Toplam %d AYT test verisi başarıyla eklendi!\n", len(allTestData))
	fmt.Printf("📊 EA Denemeleri: %d adet\n", len(eaTestData))
	fmt.Printf("📊 MF Denemeleri: %d adet\n", len(mfTestData))
	fmt.Println("\n📚 EA Subject'ler:")
	fmt.Println("   • Türk Dili ve Edebiyatı (24 soru)")
	fmt.Println("   • Tarih (10 soru)")
	fmt.Println("   • Coğrafya (6 soru)")
	fmt.Println("   • Matematik (40 soru)")
	fmt.Println("\n🔬 MF Subject'ler:")
	fmt.Println("   • Matematik (40 soru)")
	fmt.Println("   • Fizik (14 soru)")
	fmt.Println("   • Kimya (14 soru)")
	fmt.Println("   • Biyoloji (13 soru)")
}
