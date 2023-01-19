package userdto

type UserUpdateRequest struct {
	Name     string `json:"name" gorm:"type: varchar(255)" validate:"required"`
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Image    string `json:"image" gorm:"type: varchar(255)" validate:"required"`
	Phone    string `json:"phone" gorm:"type: varchar(255)" validate:"required"`
	PosCode  string `json:"pos_code" gorm:"type: varchar(255)" validate:"required"`
	Address  string `json:"address" gorm:"type: varchar(255)" validate:"required"`
}
