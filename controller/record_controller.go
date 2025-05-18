package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllRecords 函数用于获取所有记录信息
//
// 参数：
//   - ctx: *gin.Context - gin 框架的上下文对象，用于获取请求参数和返回响应
//
// 返回值：
//   - 无
func GetAllRecords(ctx *gin.Context) {
	var records []Record
	if err := GetAllData(ctx, &records, nil); err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusOK, records)
}

func GetRecentRecords(ctx *gin.Context) {
	var records []Record
	if err := GetAllData(ctx, &records, nil); err != nil {
		return
	}
}

func SubmitOrder(ctx *gin.Context) {
	var bills []Bill
	if err := BindJSON(ctx, &bills); err != nil {
		return
	}
	for _, bill := range bills {
		cnt := bill.Count
		// 防止篡改数量
		if cnt <= 0 {
			log.Printf("Count must be positive, dish_id: %d, count: %d\n", bill.DishID, cnt)
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": "Count must be positive",
			})
			return
		}
		record := Record{
			DishID: bill.DishID,
			Time:   time.Now().Format("2006-01-02 15:04:05"),
			Count:  bill.Count,
		}
		CreateDataWithoutBind(ctx, &record)
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"msg": "提交成功",
	})
}
