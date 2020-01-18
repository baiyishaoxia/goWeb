/**
 * 工具
 * @desc 接口返回的通用类型：方便控制器json格式化
 * ---------------------------------------------------------------------
 * @author		Super <super@papa.com.cn>
 * @date		2019-11-07
 * @copyright	cooper
 * ---------------------------------------------------------------------
 */

package app

import (
	"encoding/json"
	"errors"
	"strconv"
)

//NewMP ...
func NewMP() Mp {
	return make(map[string]interface{}, 0)
}

//ListJoin 分页列表
func ListJoin(slice, total, page, size interface{}) Mp {
	return NewMP().Set("list", slice).Set("count", total).Set("page", page).Set("size", size)
}

//Mp ...
type Mp map[string]interface{}

func (that Mp) Init(mode interface{}) (Mp, error) {
	_json, _err := json.Marshal(mode)
	if _err != nil {
		return nil, _err
	}
	if _err := json.Unmarshal(_json, &that); _err != nil {
		return nil, _err
	}
	return that, nil
}

//Set ...
func (that Mp) Set(key string, value interface{}) Mp {
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

func (that Mp) Copy() Mp {
	_mp := NewMP()
	for _k, _v := range that {
		_mp.Set(_k, _v)
	}
	return _mp
}

//Get ...
func (that Mp) Get(key string) interface{} {
	return that[key]
}

//Del ...
func (that Mp) Del(key string) Mp {
	if that.Has(key) {
		delete(that, key)
	}
	return that
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
