package main

import "fmt"

//region Remark:切片 Author:tang
func main() {
	var s1 []int
	fmt.Println("空的切片:", s1)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	//从0开始切,到下标3就截止
	s2 := a[0:3]
	fmt.Println("切片下标0,1,2:", s2)
	s2 = a[:5]
	fmt.Println("切片前5个:", s2)
	s2 = a[5:len(a)]
	fmt.Println("切片5-最后:", s2)

	//包含3个元素容量为10的slice 
	s3 := make([]int, 3, 10)
	s3 = a[0:5]
	fmt.Println("切片a数组", s3, "长度:", len(s3), "容量:", cap(s3))

	//slice与底层数组的关系 (但是不能越界,不能超过最开始的数组长度)
	s4 := s3[3:10]
	fmt.Println("切片后再切数组", s4, "容量", cap(s4))

	//追加
	s5 := append(s3[0:3], 11, 12)
	fmt.Printf("追加后的切片A:%v 地址:%p \n", s5, s5)
	s5 = append(s3, 14)
	fmt.Printf("追加后的切片B:%v 地址:%p \n", s5, s5)
	//s3的容量不足,将重新分配空间
	s5 = append(s3, 15, 16, 17, 18, 19, 20, 21)
	fmt.Printf("追加后的切片C:%v 地址:%p \n", s5, s5)

	//copy
	s6 := []int{1, 2, 3, 4, 5, 6}
	s7 := []int{7, 8, 9}
	copy(s6, s7)
	fmt.Println("s7 copy s6:", s6)
	s8 := []int{1, 2, 3, 4, 5, 6}
	s9 := []int{7, 8, 9, 10, 1, 1, 1, 1, 1}
	copy(s8, s9)
	fmt.Println("s9 copy s8:", s8)
	copy(s7[0:1], s9[3:4])
	fmt.Println("s9 copy s7", s7)
	s10 := s9[0:len(s9)] //简写  s10:=s9[:]
	fmt.Println("将s9完整copy到s10中:", s10)

}

//endregion
