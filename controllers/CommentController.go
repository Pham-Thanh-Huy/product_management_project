// controllers/comment_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-management-project/models"
	"product-management-project/services"
	"strconv"
	"time"
)

type CommentController struct {
	commentService *services.CommentService
}

func NewCommentController(commentService *services.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

// CreateComment: Thêm bình luận mới
func (cc *CommentController) CreateComment(ctx *gin.Context) {
	var comment models.CommentModel
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Định dạng json không hợp lệ, vui lòng truyền đúng các giá trị có trong comment",
			"error":   err.Error(),
		})
		return
	}

	// Thời gian tạo bình luận
	now := time.Now()
	comment.CreatedAt = &now

	// Tạo bình luận mới
	createdComment, err := cc.commentService.CreateComment(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Lỗi trong quá trình thêm bình luận",
			"error":   err.Error(),
		})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Thêm bình luận thành công",
		"data":    createdComment,
	})
}

// FindCommentsByProduct: Lấy tất cả bình luận của một sản phẩm
func (cc *CommentController) FindCommentsByProduct(ctx *gin.Context) {
	// Lấy productId từ URL params
	productId, err := strconv.Atoi(ctx.Param("productId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Product ID phải là số",
		})
		return
	}

	// Lấy danh sách bình luận
	comments, err := cc.commentService.FindCommentsByProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Lỗi trong quá trình lấy bình luận",
			"error":   err.Error(),
		})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}

// FindCommentById: Lấy chi tiết bình luận theo ID
func (cc *CommentController) FindCommentById(ctx *gin.Context) {
	// Lấy id bình luận từ URL params
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID phải là số",
		})
		return
	}

	// Lấy bình luận
	comment, err := cc.commentService.FindCommentsByProduct(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Lỗi trong quá trình lấy bình luận",
			"error":   err.Error(),
		})
		return
	}

	// Trả về kết quả
	ctx.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}
