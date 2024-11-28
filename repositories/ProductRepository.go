package repositories

import (
	"errors"
	"gorm.io/gorm"
	"product-management-project/models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindAllProducts() ([]models.ProductModel, error) {
	var models []models.ProductModel

	err := r.db.Preload("User").Find(&models).Error

	return models, err
}

func (r *ProductRepository) FindProductById(id int) (models.ProductModel, error) {
	var product models.ProductModel
	err := r.db.Preload("User").First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.ProductModel{}, errors.New("Không tìm thấy sản phẩm theo id này")
		}
	}
	return product, err
}

func (r *ProductRepository) CreateProduct(product models.ProductModel, userId int) (models.ProductModel, error) {
	// Kiểm tra xem user có tồn tại không
	var user models.UserModel
	err := r.db.First(&user, userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Trả về lỗi nếu không tìm thấy user
			return product, errors.New("Người dùng không tồn tại với id này")
		}
		// Trả về lỗi khác (lỗi cơ sở dữ liệu)
		return product, err
	}

	// Gán userId vào sản phẩm trước khi lưu
	product.UserID = uint(userId)

	// Tạo sản phẩm
	err = r.db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *ProductRepository) UpdateProduct(product models.ProductModel, id int) (models.ProductModel, error) {
	var existingProduct models.ProductModel
	err := r.db.First(&existingProduct, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Product not found
			return product, errors.New("product not found")
		}
		// Some other error (database error)
		return product, err
	}

	// Update the product details
	err = r.db.Model(&existingProduct).Updates(product).Error
	if err != nil {
		return product, err
	}
	return existingProduct, nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
	var product models.ProductModel
	err := r.db.First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Product not found
			return errors.New("product not found")
		}
		// Some other error (database error)
		return err
	}

	// Delete the product
	err = r.db.Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
