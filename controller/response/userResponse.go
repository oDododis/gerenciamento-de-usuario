package response

// Responde o Usuario sem a senha

type UserResponse struct {
	ID       uint   `json:"id"` // não sei se é necessario
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
