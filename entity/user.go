package entity

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	FirstName string `gorm:"type:varchar(255)" json:"firstname"`
	LastName  string `gorm:"type:varchar(255)" json:"lastname"`
	Email     string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Address   string `gorm:"type:varchar(255)" json:"address"`
	Age       uint8  `gorm:"type:int(5)" json:"age"`
}
