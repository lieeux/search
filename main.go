package main

import (
	"search/method"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	method.Dichotomous(&arr, 0, len(arr)-1, 5)
}
