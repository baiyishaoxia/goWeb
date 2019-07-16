package main

import (
	"errors"
	"fmt"
	"sync"
)

//线程同步(锁)
func main() {
	m := &MyMap{mp: make(map[string]int), mutex: new(sync.Mutex)}
	go SetValue(m)
	go m.Display()
	var str string
	fmt.Scan(&str)
}

type MyMap struct {
	mp    map[string]int
	mutex *sync.Mutex
}

func (this *MyMap) Get(key string) (int, error) {
	this.mutex.Lock()
	i, ok := this.mp[key]
	this.mutex.Unlock()
	if !ok {
		return i, errors.New("不存在")
	}
	return i, nil
}

func (this *MyMap) Set(key string, val int) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.mp[key] = val
}

func (this *MyMap) Display() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	for key, val := range this.mp {
		fmt.Println(key, "=", val)
	}
}

func SetValue(m *MyMap) {
	var a rune
	a = 'a'
	for i := 0; i < 10; i++ {
		m.Set(string(a+rune(i)), i)
	}
}
