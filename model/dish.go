package model

type Dish struct {
	ID       uint `gorm:"primaryKey`
	Name     string
	Price    int
	Category string
}
