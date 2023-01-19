package userdto

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
}

type UserUpdateResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	PosCode  string `json:"pos_code" form:"pos_code" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
}

type UserDeleteResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" validate:"required"`
}
