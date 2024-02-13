package request

// Requisita o User com restriçoes

type UserRequest struct {
	ID       string `json:"id" binding:""`
	FullName string `json:"fullName" binding:"required,min=3,max=150"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=150"`
	Password string `json:"password" binding:"required,min=8,containsany=!@#$%&*()_+"`
	//Birthday time.Time `json:"birthday" binding:""`
}
