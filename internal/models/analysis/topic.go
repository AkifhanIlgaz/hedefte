package analysis

type Topic struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	Name      string   `json:"name" gorm:"type:text;not null"`
	SubjectID uint     `json:"subject_id" gorm:"not null"`
	ExamType  ExamType `json:"exam_type" gorm:"not null"`

	// Relations
	Subject       Subject        `json:"subject" gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE"`
	TopicMistakes []TopicMistake `json:"topic_mistakes" gorm:"foreignKey:TopicID;constraint:OnDelete:CASCADE"`
}
