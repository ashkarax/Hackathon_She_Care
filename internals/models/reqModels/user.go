package requestmodel

type UserSignup struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type UserLogin struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}
