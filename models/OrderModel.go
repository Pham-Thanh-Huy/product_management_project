package models

import "time"

type OrderModel struct {
	Id          int         `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID      int         `json:"user_id" gorm:"column:user_id;not null"`
	User        UserModel   `json:"user" gorm:"foreignKey:UserID"`
	TotalAmount float64     `json:"total_amount" gorm:"column:total_amount"`
	Status      string      `json:"status" gorm:"column:status"` // pending, completed, canceled
	CreatedAt   *time.Time  `json:"created_at" gorm:"column:created_at;type:datetime"`
	UpdatedAt   *time.Time  `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	Items       []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

func (OrderModel) TableName() string {
	return "orders"
}

// Bảng chi tiết đơn hàng (OrderItem)
type OrderItem struct {
	Id        int          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	OrderID   int          `json:"order_id" gorm:"column:order_id;not null"`
	Order     OrderModel   `json:"order" gorm:"foreignKey:OrderID"`
	ProductID int          `json:"product_id" gorm:"column:product_id;not null"`
	Product   ProductModel `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int          `json:"quantity" gorm:"column:quantity;not null"`
	Price     float64      `json:"price" gorm:"column:price"`
	CreatedAt *time.Time   `json:"created_at" gorm:"column:created_at;type:datetime"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
