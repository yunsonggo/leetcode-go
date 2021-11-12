package string

import (
	"errors"
	"regexp"
	"strings"
)

/*
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

函数 myAtoi(string s) 的算法如下：

读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−2 31,  2 31 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
返回整数作为最终结果。
注意：

本题中的空白字符只包括空格字符 ' ' 。
除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-to-integer-atoi
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

const INT32_MAX = int32(^uint32(0) >> 1)

//const INT32_MIN = int32(^INT32_MAX)

func myAtoi(s string) (int32, error) {

	// step1: 去掉首尾空格
	s0 := strings.TrimSpace(s)
	// 保存去除首尾字符
	var s1 string
	// 保存首位符号
	var s2 string
	// 保存踢出非数字后的字符串
	var s3 string
	// 用户保存结果
	n := 0
	var err error
	// step2: 判断首尾是否正负号
	if s0[0] == '-' || s0[0] == '+' {
		s2 = string(s0[0])
		s1 = s0[1:]
		if len(s1) < 1 {
			err = errors.New("只有正负号吗?转嘛呢")
			return 0, err
		}
	} else {
		s1 = s0
	}
	// 我这里使用正则 挨个过滤 凡不是数字的字符 返回 空字符
	for i := 0; i < len(s1); i++ {
		s3 += verifyNum(string(s1[i]))
	}
	// 循环纯数字的字符串s3
	for _, ch := range []byte(s3) {
		ch -= '0' // rune int32的别名类型
		if ch > 9 {
			err = errors.New("单个数字字符格式不是0-9的rune")
			return 0, err
		}
		n = n*10 + int(ch)
		// 如果大于32位最大值-1 返回
		if n >= int(INT32_MAX-1) {
			err = errors.New("已超出32位最大值")
			if s2 == "-" {
				n = -n
			}
			return int32(n), err
		}
	}
	if s2 == "-" {
		n = -n
	}
	return int32(n), err
}

func verifyNum(s string) string {
	pattern := `^[0-9]`
	reg := regexp.MustCompile(pattern)
	res := reg.MatchString(s)
	if !res {
		return ""
	}
	return s
}
