package analysis

type Subject struct {
	ID             uint     `json:"id" gorm:"primaryKey"`
	Name           string   `json:"name" gorm:"type:text;not null"`
	TotalQuestions int      `json:"total_questions" gorm:"type:int;not null"`
	ExamType       ExamType `json:"exam_type" gorm:"not null"`

	// Relations
	Topics       []Topic       `json:"topics" gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE"`
	ExamSubjects []ExamSubject `json:"exam_subjects" gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE"`
}

type SubjectResponse struct {
	ID             uint     `json:"id"`
	Name           string   `json:"name"`
	TotalQuestions int      `json:"total_questions"`
	ExamType       ExamType `json:"exam_type"`
}
