package responsemodel

type UserSignup struct {
	Error string
	AccessToken string
	RefreshToken string
	Name  string `json:"name"`
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Errors struct {
	Err string
}
