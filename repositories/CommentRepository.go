package repositories

import (
	"gorm.io/gorm"
	"product-management-project/models"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(comment models.CommentModel) (models.CommentModel, error) {
	err := r.db.Create(&comment).Error
	return comment, err
}

func (r *CommentRepository) FindCommentsByProduct(productId int) ([]models.CommentModel, error) {
	var comments []models.CommentModel
	err := r.db.Preload("User").Where("product_id = ?", productId).Find(&comments).Error
	return comments, err
}
