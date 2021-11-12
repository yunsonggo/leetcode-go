package string

import (
	"fmt"
	"strings"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

func noRepe(s string) {
	if len(s) == 0 {
		fmt.Printf("字符串:%s 中 最长的无重复字符子串为:0\n",s)
		return
	}
	key := 0
	var arr []string
	for {
		bs := s[key:]
		if len(bs) == 0 {
			break
		}
		temp := ""
		for _,v := range bs {
			vs := fmt.Sprintf("%s",string(v))
			count := strings.Count(temp,vs)
			if count >= 1 {
				arr = append(arr,temp)
				break
			}
			temp = fmt.Sprintf("%s%s",temp,vs)
			key ++
		}
	}
	var resValue = ""
	for _,value := range arr {
		if len(value) > len(resValue) {
			resValue = value
		}
	}
	fmt.Printf("字符串:%s 中 最长的无重复字符子串为:%s\n",s,resValue)
}

