package study_material

type AddStudyMaterialRequest struct {
	UserId   string `json:"userId"`
	LessonId string `json:"lessonId" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type DeleteStudyMaterialRequest struct {
	Id     string `json:"id" validate:"required"`
	UserId string `json:"userId" validate:"required"`
}

type GetStudyMaterialsRequest struct {
	LessonId string `json:"lessonId" validate:"required" form:"lessonId"`
	UserId   string `json:"userId"`
}
