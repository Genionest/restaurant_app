package controller

import (
	"log"
	"net/http"
	"sync"
	"time"

	"example.com/m/v2/global"
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
	if ok := GetAllDatas(ctx, &records, nil); !ok {
		return
	}
	ctx.IndentedJSON(http.StatusOK, records)
}

// SubmitOrder 处理订单提交请求，将订单中的账单信息转换为记录存入数据库，
// 并在订单提交成功后异步删除过期记录。
//
// 参数:
// ctx: *gin.Context - gin 框架的上下文对象，用于处理 HTTP 请求和响应。
//
// 返回值:
// 无返回值
func GetRecentRecords(ctx *gin.Context) {
	var records []Record
	if ok := GetAllDatas(ctx, &records, nil); !ok {
		return
	}
}

func SubmitOrder(ctx *gin.Context) {
	var bills []Bill
	if ok := BindJSON(ctx, &bills); !ok {
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

	// 清空过期记录
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		DeleteExpiredRecords(ctx)
	}()
	wg.Wait()
}

func DeleteExpiredRecords(ctx *gin.Context) {
	// 超过一个月的记录
	expiredTime := time.Now().AddDate(0, -1, 0).Format("2006-01-02 15:04:05")
	var records []Record
	if err := global.DB.Where("time < ?", expiredTime).Delete(&records).Error; err != nil {
		log.Println()
		log.Printf("Delete expired records error")
		log.Printf("expired_time: %s\n", expiredTime)
		log.Printf("%v\n", err.Error())
		log.Println()
		return
	}
}
