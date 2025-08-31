package db

import (
	"log"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/models/analysis"
	"gorm.io/gorm"
)

func SeedTYTSubjectsAndTopics(db *gorm.DB) error {
	type subjectWithTopics struct {
		Subject models.Subject
		Topics  []string
	}

	// ExamType olarak TYT kullanıyoruz
	const examType = "TYT"

	subjects := []subjectWithTopics{
		{
			Subject: models.Subject{Name: "Türkçe", TotalQuestions: 40, ExamType: examType},
			Topics: []string{
				"Sözcükte Anlam", "Söz Yorumu", "Deyim ve Atasözü", "Cümlede Anlam", "Paragraf",
				"Paragrafta Anlatım Teknikleri", "Paragrafta Düşünceyi Geliştirme Yolları", "Paragrafta Yapı",
				"Paragrafta Konu-Ana Düşünce", "Paragrafta Yardımcı Düşünce", "Ses Bilgisi", "Yazım Kuralları",
				"Noktalama İşaretleri", "Sözcükte Yapı/Ekler", "Sözcük Türleri", "İsimler", "Zamirler", "Sıfatlar",
				"Zarflar", "Edat – Bağlaç – Ünlem", "Fiiller", "Fiilde Anlam (Kip-Kişi-Yapı)", "Ek Fiil", "Fiilimsi",
				"Fiilde Çatı", "Sözcük Grupları", "Cümlenin Ögeleri", "Cümle Türleri", "Anlatım Bozukluğu",
			},
		},
		{
			Subject: models.Subject{Name: "Tarih", TotalQuestions: 5, ExamType: examType},
			Topics: []string{
				"Tarih ve Zaman", "İnsanlığın İlk Dönemleri", "Orta Çağ’da Dünya", "İlk ve Orta Çağlarda Türk Dünyası",
				"İslam Medeniyetinin Doğuşu", "Türklerin İslamiyet’i Kabulü ve İlk Türk İslam Devletleri",
				"Yerleşme ve Devletleşme Sürecinde Selçuklu Türkiyesi", "Beylikten Devlete Osmanlı Siyaseti",
				"Devletleşme Sürecinde Savaşçılar ve Askerler", "Beylikten Devlete Osmanlı Medeniyeti",
				"Dünya Gücü Osmanlı", "Sultan ve Osmanlı Merkez Teşkilatı", "Klasik Çağda Osmanlı Toplum Düzeni",
				"Değişen Dünya Dengeleri Karşısında Osmanlı Siyaseti", "Değişim Çağında Avrupa ve Osmanlı",
				"Uluslararası İlişkilerde Denge Stratejisi (1774-1914)", "Devrimler Çağında Değişen Devlet-Toplum İlişkileri",
				"Sermaye ve Emek", "XIX. ve XX. Yüzyılda Değişen Gündelik Hayat", "XX. Yüzyıl Başlarında Osmanlı Devleti ve Dünya",
				"Milli Mücadele", "Atatürkçülük ve Türk İnkılabı",
			},
		},
		{
			Subject: models.Subject{Name: "Coğrafya", TotalQuestions: 5, ExamType: examType},
			Topics: []string{
				"Doğa ve İnsan", "Dünya’nın Şekli ve Hareketleri", "Coğrafi Konum", "Harita Bilgisi", "Atmosfer ve Sıcaklık",
				"İklimler", "Basınç ve Rüzgarlar", "Nem, Yağış ve Buharlaşma", "İç Kuvvetler / Dış Kuvvetler",
				"Su – Toprak ve Bitkiler", "Nüfus", "Göç", "Yerleşme", "Türkiye’nin Yer Şekilleri", "Ekonomik Faaliyetler",
				"Bölgeler", "Uluslararası Ulaşım Hatları", "Çevre ve Toplum", "Doğal Afetler",
			},
		},
		{
			Subject: models.Subject{Name: "Felsefe", TotalQuestions: 5, ExamType: examType},
			Topics: []string{
				"Felsefe’nin Konusu", "Bilgi Felsefesi", "Varlık Felsefesi", "Ahlak Felsefesi", "Sanat Felsefesi",
				"Din Felsefesi", "Siyaset Felsefesi", "Bilim Felsefesi", "İlk Çağ Felsefesi",
				"2. Yüzyıl ve 15. Yüzyıl Felsefeleri", "15. Yüzyıl ve 17. Yüzyıl Felsefeleri",
				"18. Yüzyıl ve 19. Yüzyıl Felsefeleri", "20. Yüzyıl Felsefesi",
			},
		},
		{
			Subject: models.Subject{Name: "Din Kültürü", TotalQuestions: 5, ExamType: examType},
			Topics: []string{
				"Bilgi ve İnanç", "İslam ve İbadet", "Ahlak ve Değerler", "Allah İnsan İlişkisi", "Hz. Muhammed (S.A.V.)",
				"Vahiy ve Akıl", "İslam Düşüncesinde Yorumlar, Mezhepler", "Din, Kültür ve Medeniyet",
				"İslam ve Bilim, Estetik, Barış", "Yaşayan Dinler",
			},
		},
		{
			Subject: models.Subject{Name: "Matematik", TotalQuestions: 40, ExamType: examType},
			Topics: []string{
				"Temel Kavramlar", "Sayı Basamakları", "Bölme ve Bölünebilme", "EBOB – EKOK", "Rasyonel Sayılar",
				"Basit Eşitsizlikler", "Mutlak Değer", "Üslü Sayılar", "Köklü Sayılar", "Çarpanlara Ayırma",
				"Oran Orantı", "Denklem Çözme", "Problemler", "Sayı Problemleri", "Kesir Problemleri",
				"Yaş Problemleri", "Hareket Hız Problemleri", "İşçi Emek Problemleri", "Yüzde Problemleri",
				"Kar Zarar Problemleri", "Karışım Problemleri", "Grafik Problemleri", "Rutin Olmayan Problemler",
				"Kümeler – Kartezyen Çarpım", "Mantık", "Fonksiyonlar", "Polinomlar", "2. Dereceden Denklemler",
				"Permütasyon ve Kombinasyon", "Olasılık", "Veri – İstatistik", "Temel Kavramlar", "Doğruda Açılar",
				"Üçgende Açılar", "Özel Üçgenler", "Dik Üçgen", "İkizkenar Üçgen", "Eşkenar Üçgen", "Açıortay",
				"Kenarortay", "Eşlik ve Benzerlik", "Üçgende Alan", "Üçgende Benzerlik", "Açı Kenar Bağıntıları",
				"Çokgenler", "Özel Dörtgenler", "Dörtgenler", "Deltoid", "Paralelkenar", "Eşkenar Dörtgen",
				"Dikdörtgen", "Kare", "Yamuk", "Çember ve Daire", "Çemberde Açı", "Çemberde Uzunluk",
				"Dairede Çevre ve Alan", "Analitik Geometri", "Noktanın Analitiği", "Doğrunun Analitiği",
				"Dönüşüm Geometrisi", "Katı Cisimler", "Prizmalar", "Küp", "Silindir", "Piramit", "Koni",
				"Küre", "Çemberin Analitiği",
			},
		},
		{
			Subject: models.Subject{Name: "Fizik", TotalQuestions: 7, ExamType: examType},
			Topics: []string{
				"Fizik Bilimine Giriş", "Madde ve Özellikleri", "Sıvıların Kaldırma Kuvveti", "Basınç",
				"Isı, Sıcaklık ve Genleşme", "Hareket ve Kuvvet", "Dinamik", "İş, Güç ve Enerji", "Elektrik",
				"Manyetizma", "Dalgalar", "Optik",
			},
		},
		{
			Subject: models.Subject{Name: "Kimya", TotalQuestions: 7, ExamType: examType},
			Topics: []string{
				"Kimya Bilimi", "Atom ve Periyodik Sistem", "Kimyasal Türler Arası Etkileşimler", "Maddenin Halleri",
				"Doğa ve Kimya", "Kimyanın Temel Kanunları", "Kimyasal Hesaplamalar", "Karışımlar",
				"Asit, Baz ve Tuz", "Kimya Her Yerde",
			},
		},
		{
			Subject: models.Subject{Name: "Biyoloji", TotalQuestions: 6, ExamType: examType},
			Topics: []string{
				"Canlıların Ortak Özellikleri", "Canlıların Temel Bileşenleri", "Hücre ve Organelleri",
				"Hücre Zarından Madde Geçişi", "Canlıların Sınıflandırılması", "Mitoz ve Eşeysiz Üreme",
				"Mayoz ve Eşeyli Üreme", "Kalıtım", "Ekosistem Ekolojisi", "Güncel Çevre Sorunları",
			},
		},
	}

	for _, s := range subjects {
		var subject models.Subject
		// Önce subject var mı kontrol et, yoksa ekle
		result := db.Where("name = ? AND exam_type = ?", s.Subject.Name, examType).First(&subject)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if err := db.Create(&s.Subject).Error; err != nil {
					log.Printf("Subject eklenemedi: %s, hata: %v", s.Subject.Name, err)
					return err
				}
				subject = s.Subject
				log.Printf("Subject eklendi: %s", subject.Name)
			} else {
				return result.Error
			}
		}

		// Her topic için var mı kontrol et, yoksa ekle
		for _, topicName := range s.Topics {
			var topic models.Topic
			topicResult := db.Where("name = ? AND subject_id = ?", topicName, subject.ID).First(&topic)
			if topicResult.Error != nil {
				if topicResult.Error == gorm.ErrRecordNotFound {
					newTopic := models.Topic{Name: topicName, SubjectID: subject.ID, ExamType: analysis.TYT}
					if err := db.Create(&newTopic).Error; err != nil {
						log.Printf("Topic eklenemedi: %s, hata: %v", topicName, err)
						return err
					}
					log.Printf("Topic eklendi: %s", topicName)
				} else {
					return topicResult.Error
				}
			}
		}
	}

	return nil
}
