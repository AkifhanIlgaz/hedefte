package services

import (
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"gorm.io/gorm"
)

type AnalysisService struct {
	db *gorm.DB
}

func NewAnalysisService(db *gorm.DB) AnalysisService {
	return AnalysisService{
		db: db,
	}
}
func (s AnalysisService) Add(uid string, req models.ExamAnalysisRequest) error {
	examSubjects := make([]models.ExamSubject, 0, len(req.Subjects))
	for _, examSub := range req.Subjects {
		topicMistakes := make([]models.TopicMistake, 0, len(examSub.TopicMistakes))
		for _, tm := range examSub.TopicMistakes {
			topicMistakes = append(topicMistakes, models.TopicMistake{
				TopicID: tm.Id, // Topic alanı string ise burada ID'ye dönüştürülmeli!
				Wrong:   tm.Mistakes,
			})
		}
		examSubjects = append(examSubjects, models.ExamSubject{
			SubjectID:     examSub.Id, // Name alanı string, burada subject ID'ye dönüştürülmeli!
			Correct:       examSub.Correct,
			Wrong:         examSub.Wrong,
			Empty:         examSub.Empty,
			Net:           float64(examSub.Correct) - float64(examSub.Wrong)*0.25,
			TopicMistakes: topicMistakes,
		})
	}

	totalNet := calculateTotalNet(req.Subjects)

	exam := models.Exam{
		UID:          uid,
		CreatedAt:    time.Now(),
		ExamType:     req.ExamType,
		Name:         req.Name,
		TotalNet:     totalNet,
		ExamSubjects: examSubjects,
	}

	if err := s.db.Create(&exam).Error; err != nil {
		return err
	}

	return nil
}

func (s AnalysisService) GetExam(uid string, id uint) (*models.Exam, error) {
	var exam models.Exam
	err := s.db.Preload("ExamSubjects").
		Preload("ExamSubjects.Subject").
		Preload("ExamSubjects.TopicMistakes").
		Preload("ExamSubjects.TopicMistakes.Topic").
		Where("uid = ? AND id = ?", uid, id).
		First(&exam).Error
	if err != nil {
		return nil, err
	}
	return &exam, nil
}

func (s AnalysisService) GetAllExams(uid string, exam models.ExamType) ([]models.Exam, error) {
	var exams []models.Exam
	err := s.db.Preload("ExamSubjects").
		Preload("ExamSubjects.Subject").
		Preload("ExamSubjects.TopicMistakes").
		Preload("ExamSubjects.TopicMistakes.Topic").
		Where("uid = ? AND exam_type = ?", uid, exam).
		Find(&exams).Error
	if err != nil {
		return nil, err
	}
	return exams, nil
}

func calculateTotalNet(subjects []models.SubjectRequest) float64 {
	var totalNet float64
	for _, subject := range subjects {
		totalNet += float64(subject.Correct) + (float64(subject.Wrong) * 0.25)
	}
	return totalNet
}

type ExamSubjectResponse struct {
	SubjectName   string  `json:"subjectName"`
	Correct       int     `json:"correct"`
	Wrong         int     `json:"wrong"`
	Empty         int     `json:"empty"`
	Net           float64 `json:"net"`
	TopicMistakes []struct {
		TopicName string `json:"topicName"`
		Wrong     int    `json:"wrong"`
	} `json:"topicMistakes"`
}

type ExamDetailsResponse struct {
	ID           uint                  `json:"id"`
	UID          string                `json:"uid"`
	Name         string                `json:"name"`
	TotalNet     float64               `json:"totalNet"`
	CreatedAt    string                `json:"createdAt"`
	ExamSubjects []ExamSubjectResponse `json:"examSubjects"`
}

func ToExamDetailsResponse(exam *models.Exam) ExamDetailsResponse {
	resp := ExamDetailsResponse{
		ID:        exam.ID,
		UID:       exam.UID,
		Name:      exam.Name,
		TotalNet:  exam.TotalNet,
		CreatedAt: exam.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	for _, es := range exam.ExamSubjects {
		esResp := ExamSubjectResponse{
			SubjectName: es.Subject.Name,
			Correct:     es.Correct,
			Wrong:       es.Wrong,
			Empty:       es.Empty,
			Net:         es.Net,
		}
		for _, tm := range es.TopicMistakes {
			esResp.TopicMistakes = append(esResp.TopicMistakes, struct {
				TopicName string `json:"topicName"`
				Wrong     int    `json:"wrong"`
			}{
				TopicName: tm.Topic.Name,
				Wrong:     tm.Wrong,
			})
		}
		resp.ExamSubjects = append(resp.ExamSubjects, esResp)
	}
	return resp
}

func (s AnalysisService) GetExamResponse(uid string, id uint) (*ExamDetailsResponse, error) {
	exam, err := s.GetExam(uid, id)
	if err != nil {
		return nil, err
	}
	resp := ToExamDetailsResponse(exam)
	return &resp, nil
}
