package main

import (
	"app/vendors/anytype/models"
	"fmt"
)

//任意类型转任意类型
func main() {
	var dst1 float64
	var dst2 int64
	src := "111111"
	if err := models.ConvertAssign(&dst1, src); err != nil {
		fmt.Println("convert failed, %v", err)
	} else {
		fmt.Println("convert ok: %f", dst1)
	}
	if err := models.ConvertAssign(&dst2, src); err != nil {
		fmt.Println("convert failed, %v", err)
	} else {
		fmt.Println("convert ok: %d", dst2)
	}
}
