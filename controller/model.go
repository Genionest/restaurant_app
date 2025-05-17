package controller

type Dish struct {
	ID       uint `gorm:"primaryKey`
	Name     string
	Price    int
	Category string
}

type Record struct {
	ID     uint `gorm:"primaryKey"`
	DishID uint `gorm:"foreignKey:Dish.ID"`
	Time   string
	Count  int
}

type Bill struct {
	DishID uint
	Count  int
}
