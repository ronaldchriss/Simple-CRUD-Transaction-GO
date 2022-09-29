package user

type RegisterUserInput struct {
	Name         string `json:"name" binding:"required"`
	PasswordHash string `json:"password" binding:"required"`
	Username     string `json:"username" binding:"required"`
}

type LoginInput struct {
	Username string `json: "username" binding: "required"`
	Password string `json: "password" binding:"required"`
}

type CheckUsername struct {
	Username string `json: "email" binding: "required"`
}
