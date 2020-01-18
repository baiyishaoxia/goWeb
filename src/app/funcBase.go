/**
 * 工具
 * @desc 常用的工具函数
 * ---------------------------------------------------------------------
 */

package app

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"
)

//ToInt ...
func ToInt(d interface{}) int {
	switch n := d.(type) {
	case int:
		return n
	case string:
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(n + err.Error())
		}
		return i
	case int64:
		return int(n)
	case float32:
		return int(n)
	case float64:
		return int(n)
	case bool:
		if n {
			return 1
		}
		return 0
	case []byte:
		i, err := strconv.Atoi(string(n))
		if err != nil {
			panic(err.Error())
		}
		return i
	}
	return 0
}

//ToInt64 ...
func ToInt64(d interface{}) int64 {
	switch n := d.(type) {
	case int64:
		return n
	case string:
		if n == "" {
			return 0
		}
		i, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			panic(n + err.Error())
		}
		return i
	case int:
		return int64(n)
	case float32:
		return int64(n)
	case float64:
		return int64(n)
	case bool:
		if n {
			return 1
		}
		return 0
	case []byte:
		i, err := strconv.Atoi(string(n))
		if err != nil {
			panic(err.Error())
		}
		return int64(i)
	}
	return 0
}

//ToString ...
func ToString(d interface{}) string {
	switch n := d.(type) {
	case int64:
		return strconv.FormatInt(n, 10)
	case string:
		return n
	case int:
		return strconv.Itoa(n)
	case int8:
		return strconv.Itoa(int(n))
	case float32:
		return strconv.FormatFloat(float64(n), 'f', 2, 64)
	case float64:
		return strconv.FormatFloat(n, 'f', 2, 64)
	case []byte:
		return string(n)
	case error:
		return n.Error()
	case bool:
		if n {
			return "1"
		}
		return "0"
	}
	return "0"
}

func ToFloat64(d interface{}) float64 {
	switch n := d.(type) {
	case string:
		float, err := strconv.ParseFloat(n, 64)
		if err != nil {
			panic(n + err.Error())
		}
		return float
	}
	return 0
}

//SliceIntToInt64 ...
func SliceIntToInt64(slice []int) []int64 {
	_result := []int64{}
	for _, _v := range slice {
		_result = append(_result, int64(_v))
	}
	return _result
}

//MapToSlice map 转 slice
func MapToSlice(m map[int64]int64) []int64 {
	s := make([]int64, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

//SliceToMap slice 转 map
func SliceToMap(m interface{}) interface{} {
	switch _slice := m.(type) {
	case []int:
		_map := make(map[int]int, 0)
		for _, _v := range _slice {
			_map[_v] = _v
		}
		return _map
	case []int64:
		_map := make(map[int64]int64, 0)
		for _, _v := range _slice {
			_map[_v] = _v
		}
		return _map
	case []string:
		_map := make(map[string]string, 0)
		for _, _v := range _slice {
			_map[_v] = _v
		}
		return _map
	default:
		return nil
	}
}

//ForMatToStr 格式化时间(当前格式，目标格式)
func ForMatToStr(times interface{}, baseFormat ...string) string {
	_forMat := "2006-01-02 15:04:05"
	_toMat := "2006-01-02 15:04:05"
	if len(baseFormat) > 0 {
		_forMat = baseFormat[0]
	}
	if len(baseFormat) > 1 {
		_toMat = baseFormat[1]
	}
	parseStrTime, _ := time.Parse(_forMat, ToString(times))
	return parseStrTime.Format(_toMat)
}

//TimeStr 格式化时间
func TimeStr(format string) string {
	format = strings.Replace(format, "Y", "2006", 1)
	format = strings.Replace(format, "m", "01", 1)
	format = strings.Replace(format, "d", "02", 1)
	format = strings.Replace(format, "h", "15", 1)
	format = strings.Replace(format, "i", "04", 1)
	format = strings.Replace(format, "s", "05", 1)
	return time.Now().Format(format)
}

//UnixToStr 时间戳转字符串日期
func UnixToStr(timestamp int64, params ...string) string {
	timeNow := time.Unix(timestamp, 0) //2017-08-30 16:19:19 +0800 CST
	format := "2006-01-02 15:04:05"
	if len(params) > 0 {
		format = strings.Replace(params[0], "Y", "2006", 1)
		format = strings.Replace(params[0], "m", "01", 1)
		format = strings.Replace(params[0], "d", "02", 1)
		format = strings.Replace(params[0], "h", "15", 1)
		format = strings.Replace(params[0], "i", "04", 1)
		format = strings.Replace(params[0], "s", "05", 1)
	}
	return timeNow.Format(format) //2015-06-15 08:52:32
}

//UID ...
func UID(preID, randID interface{}) string {
	_idSlice := []string{ToString(preID)}
	_idSlice = append(_idSlice, strings.Split(TimeStr("Ymdhis"), "")[2:]...)
	_idSlice = append(_idSlice, ToString(randID))
	return strings.Join(_idSlice, "")
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
	//endregion
}

//Md5Encode ...
func Md5Encode(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//MD5加密
func MD5(str string) string {
	data := []byte(str)
	sum := fmt.Sprintf("%x\n", md5.Sum(data))
	return sum
}

//判断
func InSlice(slice []int, x int) bool {
	if len(slice) < 1 {
		return false
	}
	sort.Sort(sort.IntSlice(slice))
	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= x
	})
	if len(slice) > index {
		if slice[index] == x {
			return true
		}
	}
	return false
}
