package user

type RegisterUserInput struct {
	FullName string `json:"fullname" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type FormCreateUserInput struct {
	FullName string `form:"fullname" binding:"required"`
	UserName string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
	Error    error
}

type FormUpdateUserInput struct {
	ID       int
	FullName string `json:"fullname" binding:"required"`
	UserName string `json:"username" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Error    error
}
