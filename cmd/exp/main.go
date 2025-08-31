package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/models/analysis"
	"github.com/AkifhanIlgaz/hedefte/pkg/db"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pg, err := db.ConnectPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	getTytSubjectsWithTopics(pg)

}

type TopicResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type SubjectWithTopicsResponse struct {
	ID             uint              `json:"id"`
	Name           string            `json:"name"`
	TotalQuestions int               `json:"total_questions"`
	ExamType       analysis.ExamType `json:"exam_type"`
	Topics         []TopicResponse   `json:"topics"`
}

func getTytSubjectsWithTopics(pg *gorm.DB) {
	var tytSubjects []analysis.Subject
	if err := pg.Preload("Topics").Where("exam_type = ?", analysis.TYT).Find(&tytSubjects).Error; err != nil {
		log.Fatal(err)
	}

	var resp []SubjectWithTopicsResponse
	for _, s := range tytSubjects {
		var topics []TopicResponse
		for _, t := range s.Topics {
			topics = append(topics, TopicResponse{
				ID:   t.ID,
				Name: t.Name,
			})
		}
		resp = append(resp, SubjectWithTopicsResponse{
			ID:             s.ID,
			Name:           s.Name,
			TotalQuestions: s.TotalQuestions,
			ExamType:       s.ExamType,
			Topics:         topics,
		})
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func getTytSubjects(pg *gorm.DB) {
	var tytSubjects []models.Subject
	if err := pg.Where("exam_type = ?", analysis.TYT).Find(&tytSubjects).Error; err != nil {
		log.Fatal(err)
	}

	var resp []models.SubjectResponse
	for _, s := range tytSubjects {
		resp = append(resp, models.SubjectResponse{
			ID:             s.ID,
			Name:           s.Name,
			TotalQuestions: s.TotalQuestions,
			ExamType:       s.ExamType,
		})
	}
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
