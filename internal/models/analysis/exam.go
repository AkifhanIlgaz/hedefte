package analysis

import "time"

type Exam struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UID       string    `json:"uid" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	ExamType  ExamType  `json:"exam_type"`
	Name      string    `json:"name" gorm:"type:text;not null"`
	TotalNet  float64   `json:"total_net" gorm:"type:decimal(5,2)"`

	// Relations
	ExamSubjects []ExamSubject `json:"exam_subjects" gorm:"foreignKey:ExamID;constraint:OnDelete:CASCADE"`
}
