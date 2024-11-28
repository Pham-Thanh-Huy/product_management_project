package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-management-project/models"
	"product-management-project/services"
	"strconv"
	"time"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) FindAllProducts(ctx *gin.Context) {
	products, err := pc.productService.FindAllProduct()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Lỗi trong quá trình lấy sản phẩm",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (pc *ProductController) FindProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Id phải là số!",
		})
		return
	}

	product, err := pc.productService.FindProductById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Lỗi trong quá trình lấy sản phẩm",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User id phải là số",
		})
		return
	}

	var product models.ProductModel
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Định dạng json không hợp lệ, vui lòng truyền đúng các giá trị có trong product",
			"error":   err.Error(),
		})
		return
	}

	now := time.Now()
	product.CreatedAt = &now

	productAfterAdd, err := pc.productService.CreateProduct(product, userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Lỗi trong quá trình thêm sản phẩm",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Thêm sản phẩm thành công",
		"data":    productAfterAdd,
	})
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	var product models.ProductModel
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Id phải là số",
		})
		return
	}
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Định dạng json không hợp lệ, vui lòng truyền đúng các giá trị có trong product",
			"error":   err.Error(),
		})
		return
	}

	now := time.Now()
	product.UpdatedAt = &now

	// Update product using the service layer
	updatedProduct, err := pc.productService.UpdateProduct(product, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Lỗi trong quá trình cập nhật sản phẩm",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Cập nhật sản phẩm thành công",
		"data":    updatedProduct,
	})
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Id phải là số",
		})
		return
	}

	errs := pc.productService.DeleteProduct(id)
	if errs != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Lỗi trong quá trình xóa sản phẩm",
			"error":   errs.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Xóa sản phẩm thành công",
	})
}
