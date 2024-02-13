package response

// Responde o User sem a senha

type UserResponse struct {
	//gorm.Model
	//ID       string `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username"`
	//Birthday time.Time `json:"birthday"`
}
