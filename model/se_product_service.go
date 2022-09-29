package model

import "time"

type Product struct {
	Model
	ID          int64     `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Variants    []Variant `json:"variants"`
}

func (Product) TableName() string {
	return "product"
}

type Variant struct {
	Model
	ID          int64  `json:"id"`
	VariantName string `json:"variant_name"`
	ProductId   int64  `json:"product_id"`
}

func (Variant) TableName() string {
	return "variant"
}

type Model struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	CreatedDate time.Time `gorm:"autoUpdateTime"`
	UpdatedDate time.Time `gorm:"autoCreateTime"`
}
