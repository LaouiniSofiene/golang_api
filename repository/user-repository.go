package repository

import (
	"github.com/LaouiniSofiene/golang_api/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(u entity.User) entity.User
	UpdateUser(u entity.User) entity.User
	DeleteUser(u entity.User)
	AllUser() []entity.User
	FindUserById(userID uint64) entity.User
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	GetAllUsers(u *entity.User, p *entity.Pagination) (*[]entity.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &userConnection{
		connection: dbConn,
	}
}

func (db *userConnection) InsertUser(u entity.User) entity.User {
	db.connection.Save(&u)
	db.connection.Find(&u)
	return u
}

func (db *userConnection) UpdateUser(u entity.User) entity.User {
	db.connection.Save(&u)
	db.connection.Find(&u)
	return u
}

func (db *userConnection) DeleteUser(u entity.User) {
	db.connection.Delete(&u)
}

func (db *userConnection) FindUserById(userID uint64) entity.User {
	var user entity.User
	db.connection.Find(&user, userID)
	return user
}

func (db *userConnection) AllUser() []entity.User {
	var users []entity.User
	db.connection.Find(&users)
	return users
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) GetAllUsers(user *entity.User, pagination *entity.Pagination) (*[]entity.User, error) {
	var users []entity.User
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.connection.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&entity.User{}).Where(user).Find(&users)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &users, nil
}
