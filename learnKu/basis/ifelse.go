// if-else 结构
package main

import "fmt"

func main() {
	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is between 0 and 5\n")
	} else {
		fmt.Printf("first is 5 or greater\n")
	}
	// if condition的用法
	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n")
	}

	supplierLevel, platinumFlag, regionalDiamondLevel, diamondFlag := 1, 1, 1, 0
	getSupplierLevel(supplierLevel, platinumFlag, regionalDiamondLevel, diamondFlag)
}

//
// 获取供应商等级
//
func getSupplierLevel(supplierLevel, platinumFlag, regionalDiamondLevel, diamondFlag int) {
	var memberLevel int32 = 1
	if supplierLevel == 1 {
		if diamondFlag == 1 {
			memberLevel = 5
		} else if regionalDiamondLevel == 1 {
			memberLevel = 4
		} else if platinumFlag == 1 {
			memberLevel = 3
		} else {
			memberLevel = 2
		}
	}
	fmt.Printf("memberLevel is %d\n", memberLevel)
}
