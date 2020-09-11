package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// 执行go源码文件 go run hello_world.go
func main() {
	fmt.Println("hello, world")
	fmt.Printf("%d", zero(0))
	startStr := "1998-07-01 08:49:50" //"0001-01-01 00:00:00"//"1998-07-01 08:49:50"
	if len(startStr) == 10 {
		startStr = startStr + " 00:00:00"
	}
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", startStr, time.Local)
	fmt.Println("startTime: ", startTime)
	str2 := startTime.Format("2006-01-02 15:04:05")
	fmt.Println("time to string: ", str2)
	fmt.Println("date to string: ", str2[0:10])

	fmt.Println(fmt.Sprintf("G%06s", "2565689"))
	var str = "ab,cd,ef"
	list := strings.Split(str, ",")
	fmt.Println(len(list))
	fmt.Println(list)
	fmt.Println(str[1:])
	fmt.Println(strings.Join(list, "-"))
	var supplierId = "G151638"
	id, _ := strconv.Atoi(supplierId[1:])
	fmt.Println(supplierId[1:])
	fmt.Printf("%T\n", supplierId[1:])
	fmt.Printf("%T\n", id)
	supplierId = "G" + strconv.Itoa((id + 1))
	fmt.Println(supplierId)
	fmt.Printf("%T\n", supplierId)

	strList := []string{}
	for _, tmp := range strList {
		fmt.Println("aa:", tmp)
	}
}

// 生成区间[-m, n]的安全随机数
func RangeRand(min, max int64) string {
	if min > max {
		panic("the min is greater than max!")
	}
	var data int64
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		data = result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		data = min + result.Int64()
	}

	return strconv.Itoa(int(data))
}

func zero(number int) int {
	return 100 / number
}
