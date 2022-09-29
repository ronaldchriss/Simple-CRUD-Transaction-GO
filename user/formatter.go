package user

type UserFormater struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FormatUser(user User, token string) UserFormater {
	formatter := UserFormater{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Token:    token,
	}

	return formatter
}
