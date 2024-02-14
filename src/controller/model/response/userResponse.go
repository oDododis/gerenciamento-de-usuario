package response

// Responde o User sem a senha

type UserResponse struct {
	//ID       string `json:"id"` // não sei se é necessario
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username"`
	//Birthday time.Time `json:"birthday"`
}
