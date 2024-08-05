package model

import (
	"github.com/Sotatek-HungNgo3/be-practical-order/config"
	"github.com/Sotatek-HungNgo3/be-practical-order/enum"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID            uint             `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
	ProductName   string           `json:"productName" gorm:"column:product_name;not null;type:varchar"`
	UnitPrice     float64          `json:"unitPrice" gorm:"column:unit_price;not null;type:float"`
	Quantity      int64            `json:"quantity" gorm:"column:quantity;not null;type:int8"`
	Status        enum.OrderStatus `json:"status" gorm:"column:status;not null;type:varchar"`
	PaymentMethod string           `json:"paymentMethod" gorm:"column:payment_method;not null;type:varchar"`
	ErrorMessage  string           `json:"errorMessage" gorm:"column:error_message;type:varchar"`
}

func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetPostgresDsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	_ = db.AutoMigrate(&Order{})
	return db, nil
}
