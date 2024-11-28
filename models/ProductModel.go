package models

import "time"

type ProductModel struct {
	Id                 int        `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Code               string     `json:"code" gorm:"column:code"`
	Name               string     `json:"name" gorm:"column:name"`
	ListedPrice        float64    `json:"listed_price" gorm:"column:listed_price"`
	ProductPrice       float64    `json:"product_price" gorm:"column:product_price"`
	ProductDetail      string     `json:"product_detail" gorm:"column:product_detail"`
	ProductDescription string     `json:"product_description" gorm:"column:product_description"`
	Outstanding        bool       `json:"outstanding" gorm:"column:outstanding"`
	CreatedAt          *time.Time `json:"created_at" gorm:"column:created_at;type:datetime"`
	UpdatedAt          *time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	UserID             uint       `json:"user_id" gorm:"column:user_id;not null"`
	User               UserModel  `json:"user" gorm:"foreignKey:UserID"`
}

func (ProductModel) TableName() string {
	return "product"
}
