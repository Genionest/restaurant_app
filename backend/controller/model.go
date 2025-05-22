package controller

type Dish struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Price    int
	Category string
	Img      string
}

type Record struct {
	ID     uint `gorm:"primaryKey"`
	DishID uint `gorm:"foreignKey:Dish.ID"`
	Time   string
	Count  int
}

type Bill struct {
	// no database
	DishID uint
	Count  int
}

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
	Role     string
}
