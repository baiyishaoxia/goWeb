package main

import (
	"errors"
	"fmt"
	"strconv"
)

//NewMP ...
func NewMP() Mp {
	return make(map[string]interface{}, 0)
}

//Mp ...
type Mp map[string]interface{}

//Set ...
func (that Mp) Set(key string, value interface{}) interface{} {
	that[key] = value
	return that
}

//Has ...
func (that Mp) Has(key string) bool {
	if _, _ok := that[key]; _ok {
		return true
	}
	return false
}

//Get ...
func (that Mp) Get(key string) interface{} {
	return that[key]
}

func (that Mp) String(key string) (string, error) {
	if that.Has(key) {
		switch n := that.Get(key).(type) {
		case int64:
			return strconv.FormatInt(n, 10), nil
		case string:
			return n, nil
		case int:
			return strconv.Itoa(n), nil
		case float32:
			return strconv.Itoa(int(n)), nil
		case float64:
			return strconv.Itoa(int(n)), nil
		case []byte:
			return string(n), nil
		case bool:
			if n {
				return "1", nil
			}
			return "0", nil
		}
		return "", errors.New("the type of data in map , can not to string")
	}
	return "0", errors.New("map not have this key")
}

//DefaultString ...
func (that Mp) DefaultString(key string, defaultValue string) string {
	if that.Has(key) {
		switch n := that.Get(key).(type) {
		case int64:
			return strconv.FormatInt(n, 10)
		case string:
			return n
		case int:
			return strconv.Itoa(n)
		case float32:
			return strconv.Itoa(int(n))
		case float64:
			return strconv.Itoa(int(n))
		case []byte:
			return string(n)
		case bool:
			if n {
				return "1"
			}
			return "0"
		}
		return "0"
	}
	return defaultValue
}

//map示例
func main() {
	way:="authcode"
	_mp := NewMP()
	switch way {
	case "authcode":
		_mp.Set("order_no", 1)
		_mp.Set("transaction_no",2)
	case "web":
		_mp.Set("code_url", 3)
	default:
		_mp.Set("payment_way", 6)
		_mp.Set("price", 4)
	}
	fmt.Println(_mp)
}
