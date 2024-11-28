package repositories

import (
	"gorm.io/gorm"
	"product-management-project/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() ([]models.UserModel, error) {
	var users []models.UserModel
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindById(id int) (models.UserModel, error) {
	var user models.UserModel
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepository) Create(user models.UserModel) (models.UserModel, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UserRepository) Update(user models.UserModel) error {
	return r.db.Save(&user).Error
}

func (r *UserRepository) Delete(id int) error {
	return r.db.Delete(&models.UserModel{}, id).Error
}
