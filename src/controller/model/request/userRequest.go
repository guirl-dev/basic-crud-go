package request

type UserRequest struct {
	Username string `json:"username" binding:"required,min=4,max=32"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,containsany=!@#$%*&"`
	Age      int8   `json:"age"      binding:"required,min=14,max=140"`
}
