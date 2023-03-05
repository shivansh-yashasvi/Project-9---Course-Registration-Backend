package structs

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentSignInBody struct {
	ID       string `json:"ID"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
type StudentToken struct {
	Role        string `json:"role"`
	ID          string `json:"ID"`
	TokenString string `json:"token"`
}
