package dto

type UserUpdateDTO struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required" validate:"email"`
	Address   string `json:"address" form:"address" binding:"required"`
	Age       uint8  `json:"age" form:"address" binding:"required"`
}

type UserCreateDTO struct {
	FirstName string `json:"firstname" form:"firstname" binding:"required"`
	LastName  string `json:"lastname" form:"lastname" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required" validate:"email"`
	Address   string `json:"address" form:"address" binding:"required"`
	Age       uint8  `json:"age" form:"address" binding:"required"`
}
