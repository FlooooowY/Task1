package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user Users) error
	GetAllUser() ([]Users, error)
	GetUserByID(id string) (Users, error)
	UpdateUser(user Users) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user Users) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetAllUser() ([]Users, error) {
	var users []Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserByID(id string) (Users, error) {
	var user Users
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) UpdateUser(user Users) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&Users{}, "id = ?", id).Error
}
