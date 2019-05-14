package grpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func (t *GrpcInfo) PostTest(info map[string]interface{}, reply *string) (err error) {
	id, _ := strconv.ParseInt(info["id"].(string), 10, 64)
	fmt.Println(id)
	data, _ := json.Marshal(GrpcError{Code: 200, Message: "签名成功"})
	err = errors.New(string(data))
	return
}
