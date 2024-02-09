package response

import "time"

type UserResponse struct {
	ID       string    `json:"id"`
	FullName string    `json:"fullName"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Birthday time.Time `json:"birthday"`
}
