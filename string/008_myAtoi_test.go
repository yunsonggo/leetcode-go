package string

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	s := "-=+134 2a01 52dfg+-asdF#%^dfg9+__)-"
	n, err := myAtoi(s)
	fmt.Printf("字符串为:%s\n Atoi: %v\n,错误:%v\n", s, n, err)
}
