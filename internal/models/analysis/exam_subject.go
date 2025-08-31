package analysis

type ExamSubject struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	ExamID    uint    `json:"exam_id" gorm:"not null"`
	SubjectID uint    `json:"subject_id" gorm:"not null"`
	Correct   int     `json:"correct" gorm:"type:int;default:0"`
	Wrong     int     `json:"wrong" gorm:"type:int;default:0"`
	Empty     int     `json:"empty" gorm:"type:int;default:0"`
	Net       float64 `json:"net" gorm:"type:decimal(5,2);default:0"`

	// Relations
	Exam          Exam           `json:"exam" gorm:"foreignKey:ExamID;constraint:OnDelete:CASCADE"`
	Subject       Subject        `json:"subject" gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE"`
	TopicMistakes []TopicMistake `json:"topic_mistakes" gorm:"foreignKey:ExamSubjectID;constraint:OnDelete:CASCADE"`
}
