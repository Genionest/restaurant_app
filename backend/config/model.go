package config

import (
	"log"

	"example.com/m/v2/controller"
	"example.com/m/v2/global"
)

// migrate 是一个泛型函数，用于对指定的数据模型进行数据库自动迁移操作。
// 自动迁移会根据传入的模型结构体定义，在数据库中创建或更新对应的表结构。
// 参数 data 是一个指向泛型类型 T 的指针，表示要进行迁移的模型。
// 返回值为 error 类型，如果迁移过程中出现错误，则返回具体的错误信息；否则返回 nil。
func migrate[T any](data *T) error {
	if err := global.DB.AutoMigrate(data); err != nil {
		log.Println()
		log.Printf("AutoMigrate error(CreateData)\n")
		log.Printf("error: %s\n", err.Error())
		log.Printf("struct: %+v\n", *data)
		log.Println()
		return err
	}
	return nil
}

// InitModel 初始化数据库模型，通过自动迁移的方式确保数据库表结构与 Go 结构体定义一致。
// 该函数会调用 migrate 函数对指定的模型进行自动迁移操作。
func InitModel() {
	// 自动迁移数据库
	migrate(&controller.Dish{})
	migrate(&controller.Record{})
	migrate(&controller.User{})
}
