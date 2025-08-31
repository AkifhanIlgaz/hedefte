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

// Sadece AYT subjectlerini ekler
func SeedAYTSubjects(db *gorm.DB) error {
	aytSubjects := []models.Subject{
		{Name: "Edebiyat", TotalQuestions: 24, ExamType: "AYT"},
		{Name: "Tarih", TotalQuestions: 10, ExamType: "AYT"},
		{Name: "Coğrafya", TotalQuestions: 6, ExamType: "AYT"},
		{Name: "Matematik", TotalQuestions: 40, ExamType: "AYT"},
		{Name: "Fizik", TotalQuestions: 14, ExamType: "AYT"},
		{Name: "Kimya", TotalQuestions: 13, ExamType: "AYT"},
		{Name: "Biyoloji", TotalQuestions: 13, ExamType: "AYT"},
	}

	for _, s := range aytSubjects {
		var subject models.Subject
		result := db.Where("name = ? AND exam_type = ?", s.Name, "AYT").First(&subject)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if err := db.Create(&s).Error; err != nil {
					log.Printf("AYT subject eklenemedi: %s, hata: %v", s.Name, err)
					return err
				}
				log.Printf("AYT subject eklendi: %s", s.Name)
			} else {
				return result.Error
			}
		}
	}
	return nil
}

// Subjectlere ait topicleri ekler
func SeedAYTTopics(db *gorm.DB) error {
	aytTopics := map[string][]string{
		"Edebiyat": {
			"Anlam Bilgisi", "Dil Bilgisi", "Güzel Sanatlar ve Edebiyat", "Metinlerin Sınıflandırılması",
			"Şiir Bilgisi", "Edebi Sanatlar", "Türk Edebiyatı Dönemleri", "İslamiyet Öncesi Türk Edebiyatı ve Geçiş Dönemi",
			"Halk Edebiyatı", "Divan Edebiyatı", "Tanzimat Edebiyatı", "Servet-i Fünun Edebiyatı",
			"Fecr-i Ati Edebiyatı", "Milli Edebiyat", "Cumhuriyet Dönemi Edebiyatı", "Edebiyat Akımları", "Dünya Edebiyatı",
		},
		"Tarih": {
			"Tarih ve Zaman", "İnsanlığın İlk Dönemleri", "Orta Çağ’da Dünya", "İlk ve Orta Çağlarda Türk Dünyası",
			"İslam Medeniyetinin Doğuşu", "Türklerin İslamiyet’i Kabulü ve İlk Türk İslam Devletleri",
			"Yerleşme ve Devletleşme Sürecinde Selçuklu Türkiyesi", "Beylikten Devlete Osmanlı Siyaseti",
			"Devletleşme Sürecinde Savaşçılar ve Askerler", "Beylikten Devlete Osmanlı Medeniyeti",
			"Dünya Gücü Osmanlı", "Sultan ve Osmanlı Merkez Teşkilatı", "Klasik Çağda Osmanlı Toplum Düzeni",
			"Değişen Dünya Dengeleri Karşısında Osmanlı Siyaseti", "Değişim Çağında Avrupa ve Osmanlı",
			"Uluslararası İlişkilerde Denge Stratejisi (1774-1914)", "Devrimler Çağında Değişen Devlet-Toplum İlişkileri",
			"Sermaye ve Emek", "XIX. ve XX. Yüzyılda Değişen Gündelik Hayat", "XX. Yüzyıl Başlarında Osmanlı Devleti ve Dünya",
			"Milli Mücadele", "Atatürkçülük ve Türk İnkılabı", "İki Savaş Arasındaki Dönemde Türkiye ve Dünya",
			"II. Dünya Savaşı Sürecinde Türkiye ve Dünya", "II. Dünya Savaşı Sonrasında Türkiye ve Dünya",
			"Toplumsal Devrim Çağında Dünya ve Türkiye", "XXI. Yüzyılın Eşiğinde Türkiye ve Dünya",
		},
		"Coğrafya": {
			"Biyoçeşitlilik", "Biyomlar", "Ekosistemin Unsurları", "Enerji Akışı ve Madde Döngüsü", "Ekstrem Doğa Olayları",
			"Küresel İklim Değişimi", "Nüfus Politikaları", "Türkiye’de Nüfus ve Yerleşme", "Ekonomik Faaliyetler ve Doğal Kaynaklar",
			"Göç ve Şehirleşme", "Türkiye’nin Ekonomi Politikaları", "Türkiye Ekonomisinin Sektörel Dağılımı", "Türkiye’de Tarım",
			"Türkiye’de Hayvancılık", "Türkiye’de Madenler ve Enerji Kaynakları", "Türkiye’de Sanayi", "Türkiye’de Ulaşım",
			"Türkiye’de Ticaret ve Turizm", "Geçmişten Geleceğe Şehir ve Ekonomi", "Türkiye’nin İşlevsel Bölgeleri ve Kalkınma Projeleri",
			"Hizmet Sektörünün Ekonomideki Yeri", "Küresel Ticaret", "İlk Uygarlıklar", "Kültür Bölgeleri ve Türk Kültürü",
			"Sanayileşme Süreci: Almanya", "Tarım ve Ekonomi İlişkisi Fransa – Somali", "Ülkeler Arası Etkileşim", "Jeopolitik Konum",
			"Çatışma Bölgeleri", "Küresel ve Bölgesel Örgütler", "Çevre Sorunları ve Türleri", "Madenler ve Enerji Kaynaklarının Çevreye Etkisi",
			"Doğal Kaynakların Sürdürülebilir Kullanımı", "Ekolojik Ayak İzi", "Doğal Çevrenin Sınırlılığı", "Çevre Politikaları",
			"Çevresel Örgütler", "Çevre Anlaşmaları", "Doğal Afetler",
		},
		"Matematik": {
			"Kümelerde Temel Kavramlar", "Kümelerde İşlemler", "Gerçek Sayılar", "Birinci Dereceden Denklem ve Eşitsizlikler",
			"Üslü İfadeler ve Denklemler", "Denklem ve Eşitsizliklerle İlgili Uygulamalar", "Fonksiyon Kavramı ve Gösterimi",
			"Üçgenlerin Eşliği", "Üçgenlerin Benzerliği", "Üçgenin Yardımcı Elemanları", "Dik Üçgen ve Trigonometri",
			"Üçgenin Alanı", "Merkezi Eğilim ve Yayılım Ölçüleri", "Verilerin Grafikle Gösterilmesi", "Basit Olayların Olasılıkları",
			"Sıralama ve Seçme", "Fonksiyonların Simetrileri ve Cebirsel Özellikleri", "İki Fonksiyonun Bileşkesi ve Bir Fonksiyonun Tersi",
			"Dörtgenler ve Özellikleri", "Özel Dörtgenler", "İkinci Dereceden Bir Bilinmeyenli Denklemler",
			"Polinom Kavramı ve Polinomlarla İşlemler", "Polinomlarda Çarpanlara Ayırma", "Polinom ve Rasyonel Denklemlerin Çözüm Kümeleri",
			"Katı Cisimlerin Yüzey Alanları ve Hacimleri", "Yönlü Açılar", "Trigonometrik Fonksiyonlar", "Doğrunun Analitik İncelenmesi",
			"Fonksiyonlarla İlgili Uygulamalar", "İkinci Dereceden Fonksiyonlar ve Grafikleri", "Fonksiyonların Dönüşümleri",
			"İkinci Dereceden İki Bilinmeyenli Denklem Sistemleri", "İkinci Dereceden Bir Bilinmeyenli Eşitsizlikler ve Eşitsizlik Sistemleri",
			"Çemberin Temel Elemanları", "Çemberde Açılar", "Çemberde Teğet", "Dairenin Çevresi ve Alanı", "Katı Cisimler",
			"Koşullu Olasılık", "Deneysel ve Teorik Olasılık", "Üstel Fonksiyonlar", "Logaritma Fonksiyonu",
			"Üstel ve Logaritmik Denklem ve Eşitsizlikler", "Gerçek Sayı Dizileri", "Limit ve Süreklilik", "Anlık Değişim Oranı ve Türev",
			"Türevin Uygulamaları", "Belirsiz İntegral", "Belirli İntegral ve Uygulamaları", "Toplam – Fark ve İki Kat Açı Formülleri",
			"Trigonometrik Denklemler", "Analitik Düzlemde Temel Dönüşümler", "Çemberin Analitik İncelenmesi",
			"Temel Kavramlar", "Doğruda Açılar", "Üçgende Açılar", "Özel Üçgenler", "Dik Üçgen", "İkizkenar Üçgen",
			"Eşkenar Üçgen", "Açıortay", "Kenarortay", "Üçgende Alan", "Üçgende Benzerlik", "Açı Kenar Bağıntıları",
			"Çokgenler", "Özel Dörtgenler", "Dörtgenler", "Deltoid", "Paralelkenar", "Eşkenar Dörtgen", "Dikdörtgen",
			"Kare", "İkizkenar", "Yamuk", "Çember ve Daire", "Analitik Geometri", "Noktanın Analitiği", "Doğrunun Analitiği",
			"Dönüşüm Geometrisi", "Katı Cisimler (Uzay Geometri)", "Dikdörtgenler Prizması", "Küp", "Silindir", "Piramit",
			"Koni", "Küre", "Çemberin Analitiği",
		},
		"Fizik": {
			"Vektörler", "Kuvvet, Tork ve Denge", "Kütle Merkezi", "Basit Makineler", "Hareket", "Newton’un Hareket Yasaları",
			"İş, Güç ve Enerji II", "Atışlar", "İtme ve Momentum", "Elektrik Alan ve Potansiyel", "Paralel Levhalar ve Sığa",
			"Manyetik Alan ve Manyetik Kuvvet", "İndüksiyon, Alternatif Akım ve Transformatörler", "Çembersel Hareket",
			"Dönme, Yuvarlanma ve Açısal Momentum", "Kütle Çekim ve Kepler Yasaları", "Basit Harmonik Hareket",
			"Dalga Mekaniği ve Elektromanyetik Dalgalar", "Atom Modelleri", "Büyük Patlama ve Parçacık Fiziği",
			"Radyoaktivite", "Özel Görelilik", "Kara Cisim Işıması", "Fotoelektrik Olay ve Compton Olayı",
			"Modern Fiziğin Teknolojideki Uygulamaları",
		},
		"Kimya": {
			"Kimya Bilimi", "Atom ve Periyodik Sistem", "Kimyasal Türler Arası Etkileşimler", "Kimyasal Hesaplamalar",
			"Kimyanın Temel Kanunları", "Asit, Baz ve Tuz", "Maddenin Halleri", "Karışımlar", "Doğa ve Kimya",
			"Kimya Her Yerde", "Modern Atom Teorisi", "Gazlar", "Sıvı Çözeltiler", "Kimyasal Tepkimelerde Enerji",
			"Kimyasal Tepkimelerde Hız", "Kimyasal Tepkimelerde Denge", "Asit-Baz Dengesi", "Çözünürlük Dengesi",
			"Kimya ve Elektrik", "Organik Kimyaya Giriş", "Organik Kimya", "Enerji Kaynakları ve Bilimsel Gelişmeler",
		},
		"Biyoloji": {
			"Sinir Sistemi", "Endokrin Sistem ve Hormonlar", "Duyu Organları", "Destek ve Hareket Sistemi",
			"Sindirim Sistemi", "Dolaşım ve Bağışıklık Sistemi", "Solunum Sistemi", "Üriner Sistem (Boşaltım Sistemi)",
			"Üreme Sistemi ve Embriyonik Gelişim", "Komünite Ekolojisi", "Popülasyon Ekolojisi", "Genden Proteine",
			"Nükleik Asitler", "Genetik Şifre ve Protein Sentezi", "Canlılarda Enerji Dönüşümleri", "Canlılık ve Enerji",
			"Fotosentez", "Kemosentez", "Hücresel Solunum", "Bitki Biyolojisi", "Canlılar ve Çevre",
		},
	}

	for subjectName, topics := range aytTopics {
		var subject models.Subject
		if err := db.Where("name = ? AND exam_type = ?", subjectName, "AYT").First(&subject).Error; err != nil {
			log.Printf("AYT subject bulunamadı: %s, hata: %v", subjectName, err)
			continue
		}
		for _, topicName := range topics {
			var topic models.Topic
			topicResult := db.Where("name = ? AND subject_id = ?", topicName, subject.ID).First(&topic)
			if topicResult.Error != nil {
				if topicResult.Error == gorm.ErrRecordNotFound {
					newTopic := models.Topic{Name: topicName, SubjectID: subject.ID, ExamType: analysis.AYT}
					if err := db.Create(&newTopic).Error; err != nil {
						log.Printf("AYT topic eklenemedi: %s, hata: %v", topicName, err)
						return err
					}
					log.Printf("AYT topic eklendi: %s", topicName)
				} else {
					return topicResult.Error
				}
			}
		}
	}
	return nil
}
