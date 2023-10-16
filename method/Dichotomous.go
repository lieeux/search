package method

import "fmt"

func Dichotomous(arr *[10]int, left int, right int, goal int) { //二分法
	if left > right {
		fmt.Println("找不到")
		return
	}
	mid := (left + right) / 2 //自动取整
	if goal < (*arr)[mid] {
		Dichotomous(arr, left, mid-1, goal)
	} else if goal > (*arr)[mid] {
		Dichotomous(arr, mid+1, right, goal)
	} else {
		fmt.Printf("%v下标为%v\n", goal, mid)
	}
}
