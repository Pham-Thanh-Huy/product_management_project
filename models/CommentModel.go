package models

import "time"

type CommentModel struct {
	Id        int          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID    int          `json:"user_id" gorm:"column:user_id;not null"`
	User      UserModel    `json:"user" gorm:"foreignKey:UserID"`
	ProductID int          `json:"product_id" gorm:"column:product_id;not null"`
	Product   ProductModel `json:"product" gorm:"foreignKey:ProductID"`
	Content   string       `json:"content" gorm:"column:content;type:text"`
	Rating    int          `json:"rating" gorm:"column:rating"` // 1-5 sao
	CreatedAt *time.Time   `json:"created_at" gorm:"column:created_at;type:datetime"`
}

func (CommentModel) TableName() string {
	return "comments"
}
