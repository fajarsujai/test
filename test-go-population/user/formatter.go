package user

type UserFormatter struct {
	ID          int    `json:"id"`
	FullName    string `json:"fullname"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	ProfilePict string `json:"profile_pict"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		FullName:    user.FullName,
		UserName:    user.UserName,
		Email:       user.Email,
		Token:       token,
		ProfilePict: user.ProfilePict,
	}

	return formatter
}
