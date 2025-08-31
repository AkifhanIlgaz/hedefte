package analysis

type TopicMistake struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	TopicID       uint `json:"topic_id" gorm:"not null"`
	ExamSubjectID uint `json:"exam_subject_id" gorm:"not null"`
	Wrong         int  `json:"wrong" gorm:"type:int;default:0"`

	// Relations
	Topic       Topic       `json:"topic" gorm:"foreignKey:TopicID;constraint:OnDelete:CASCADE"`
	ExamSubject ExamSubject `json:"exam_subject" gorm:"foreignKey:ExamSubjectID;constraint:OnDelete:CASCADE"`
}
