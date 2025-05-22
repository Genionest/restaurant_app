package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddDish 函数用于处理添加菜品的请求
//
// 参数:
//
//	ctx *gin.Context: Gin框架的上下文对象，用于获取请求参数和返回响应
//
// 返回值:
//
//	无
func AddDish(ctx *gin.Context) {
	var dish Dish
	if ok := CreateData(ctx, &dish); !ok {
		return
	}

	ctx.IndentedJSON(http.StatusOK, dish)
}

// GetDish 是处理HTTP GET请求的处理器函数，用于根据给定的ID获取菜品信息
// 参数：
//
//	ctx *gin.Context: Gin框架的上下文对象，包含了请求和响应信息
//
// 返回值：
//
//	无
func GetDish(ctx *gin.Context) {
	id := ctx.Param("id")
	var dish Dish
	query := map[string]interface{}{"id": id}
	if ok := GetData(ctx, &dish, query); !ok {
		return
	}
	ctx.IndentedJSON(http.StatusOK, dish)
}

// UpdateDish 更新菜品信息
//
// 参数:
//
//	ctx: gin.Context上下文对象
//	dish: Dish结构体类型的菜品信息
//
// 返回值:
//
//	无
func UpdateDish(ctx *gin.Context) {
	var dish Dish
	if ok := UpdateData(ctx, &dish); !ok {
		return
	}
	// var new_dish Dish
	// GetData(ctx, &new_dish, DataQuery{"id": dish.ID})
	// ctx.IndentedJSON(http.StatusOK, new_dish)
	// ctx.IndentedJSON(http.StatusOK, dish)
	log.Println()
	log.Printf("Update dish: %v\n", dish)
	ctx.IndentedJSON(http.StatusNoContent, nil)
}

// DeleteDish 函数用于删除菜品
// 参数:
//
//	ctx: *gin.Context - gin框架的上下文对象
//
// 返回值:
//
//	无返回值
func DeleteDish(ctx *gin.Context) {
	id := ctx.Param("id")
	iid, _ := strconv.Atoi(id)
	dish := Dish{ID: uint(iid)}
	if ok := DeleteData(ctx, &dish); !ok {
		return
	}
	// ctx.IndentedJSON(http.StatusOK, gin.H{
	// 	"id":  id,
	// 	"msg": "删除成功",
	// })
	log.Println()
	log.Printf("Delete dish: %s\n", id)
	ctx.IndentedJSON(http.StatusNoContent, nil)
}

// GetAllDishes 函数用于获取所有菜品信息
//
// 参数：
//   - ctx *gin.Context：Gin框架的上下文对象
//
// 返回值：
//
//	无返回值
func GetAllDishes(ctx *gin.Context) {
	var dishes []Dish
	if ok := GetAllDatas(ctx, &dishes, nil); !ok {
		return
	}
	ctx.IndentedJSON(http.StatusOK, dishes)
}

// GetDishesByCategory 函数根据传入的分类获取对应的菜品列表
// 参数：
//
//	ctx: *gin.Context - gin框架的上下文对象
//
// 返回值：
//
//	无
func GetDishesByCategory(ctx *gin.Context) {
	category := ctx.Param("category")
	var dishes []Dish
	query := map[string]interface{}{"category": category}
	if ok := GetAllDatas(ctx, &dishes, query); !ok {
		return
	}
	ctx.IndentedJSON(http.StatusOK, dishes)
}

func GetHotDishes(ctx *gin.Context) {
	// 获取所有record
	var records []Record
	if ok := GetAllDatas(ctx, &records, nil); !ok {
		return
	}
	// 统计每个dish的count
	dish_count := make(map[uint]int)
	for _, record := range records {
		dish_count[record.DishID] += record.Count
	}
	// 实现算法，获取map中value最大的前6个key
	// 用一个数组动态维护前6个key
	max_ids := make([]uint, 6)
	for id, cnt := range dish_count {
		for i := 0; i < 6; i++ {
			if dish_count[max_ids[i]] < cnt {
				// 其他元素往后移动
				for j := 5; j > i; j-- {
					max_ids[j] = max_ids[j-1]
				}
				// 插入到第i个位置
				max_ids[i] = id
				break // 插入成功，跳出循环
			}
		}
	}
	// 获取对应的dish
	var dishes []Dish
	for _, id := range max_ids {
		dishes = append(dishes, Dish{ID: id})
	}
	if ok := GetManyDatas(ctx, &dishes, "id in ?", max_ids); !ok {
		return
	}
	// fmt.Println(dishes)
	ctx.IndentedJSON(http.StatusOK, dishes)
}

// GetTotalPrice 函数计算并返回账单总金额
// 参数:
//
//	ctx: gin的上下文对象，用于处理HTTP请求和响应
//
// 返回值:
//
//	无返回值，通过gin的上下文对象返回总金额
func GetTotalPrice(ctx *gin.Context) {
	totalPrice := 0
	var bills []Bill
	if ok := BindJSON(ctx, &bills); !ok {
		return
	}
	query := map[string]interface{}{"id": 0}
	for _, bill := range bills {
		var dish Dish
		query["id"] = bill.DishID
		if ok := GetData(ctx, &dish, query); !ok {
			return
		}
		// 防止篡改价格
		price := dish.Price
		// 防止篡改数量
		cnt := bill.Count
		if cnt <= 0 {
			log.Printf("Count must be positive, dish_id: %d, count: %d\n", bill.DishID, cnt)
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": "Count must be positive",
			})
			return
		}
		totalPrice += price * cnt
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"total_price": totalPrice,
	})
}
