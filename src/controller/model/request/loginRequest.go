package request

// Requisita o Email e a sernha com restrições

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,containsany=!@#$%&*()_+"`
}
