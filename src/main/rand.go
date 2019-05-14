package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//随机打乱数组
func main() {
	//strs := []string{
	//	"1", "2", "3", "4", "5", "6", "7", "8",
	//}
	//a, _ := Random(strs, 3)
	//fmt.Println(a)
	fmt.Println(RemoveAllSame("1,4,9,5,8,6,7,10,11,3,2", "1,5,3,8,6,11,10,7,2"))
}

//随机
func Random(strings []string, length int) (string, error) {
	if len(strings) <= 0 {
		return "", errors.New("the length of the parameter strings should not be less than 0")
	}

	if length <= 0 || len(strings) < length {
		return "", errors.New("the size of the parameter length illegal")
	}

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < length; i++ {
		str += strings[i]
	}
	return str, nil
}

// 单个slice去重处理
func RemoveTheSame(s string) string {
	s1 := []string{}
	m := make(map[int]string)
	for _, v := range s {
		m[int(v)] = "ok"
	}
	for k, _ := range m {
		s1 = append(s1, string(k))
	}
	return strings.Join(s1, "")
}

//多个slice去重  [str1大范围,str2小范围]
func RemoveAllSame(str1 string, str2 string) string {
	str := strings.Split(str1, ",")
	data := ""
	for _, val := range str {
		if strings.Index(","+str2+",", ","+val+",") == -1 {
			//不存在
			data += val + ","
		}
	}
	return data[0 : len(data)-1]
}
