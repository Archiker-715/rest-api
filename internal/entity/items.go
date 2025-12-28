package entity

type Item struct {
	ID          uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null,check:price >=0"`
	Inserted    string  `json:"inserted"`
	Updated     string  `json:"updated"`
}

type ItemRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"min=0"`
}
