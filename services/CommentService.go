package services

import (
	"product-management-project/models"
	"product-management-project/repositories"
)

type CommentService struct {
	commentRepo *repositories.CommentRepository
}

func NewCommentService(commentRepo *repositories.CommentRepository) *CommentService {
	return &CommentService{commentRepo: commentRepo}
}

func (s *CommentService) CreateComment(comment models.CommentModel) (models.CommentModel, error) {
	return s.commentRepo.CreateComment(comment)
}

func (s *CommentService) FindCommentsByProduct(productId int) ([]models.CommentModel, error) {
	return s.commentRepo.FindCommentsByProduct(productId)
}
