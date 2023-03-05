package structs

type Student struct {
	ID       string `json:"id"`
	Name     string `json:"name"     validate:"required"`
	Password string `json:"password" validate:"required"`
}

type StudentResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
