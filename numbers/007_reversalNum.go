package numbers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围 [−2 31次方,  2 31次方 − 1] ，就返回 0。
// 假设环境不允许存储 64 位整数（有符号或无符号）。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-integer
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
const INT32_MAX = int32(^uint32(0) >> 1)
const INT32_MIN = int32(^INT32_MAX)

func reversalNum() {
	fmt.Printf("int32 max:%v,min:%v\n", INT32_MAX, INT32_MIN)
	var a int32
	rand.Seed(time.Now().Unix())
	tag := rand.Intn(2)
	n := rand.Uint32()
	if tag == 0 {
		a = -int32(n)
	} else {
		a = int32(n)
	}
	strA := strconv.FormatInt(int64(a), 10)
	fmt.Printf("tag:%d,随机数为:%v\n", tag, a)
	fmt.Printf("FormatInt:%v,len:%v\n", strA, len(strA))
	var resStr string
	if a < 0 {
		resStr += string(strA[0])
		for i := len(strA) - 1; i > 0; i-- {
			resStr += string(strA[i])
		}
	} else {
		for i := len(strA) - 1; i >= 0; i-- {
			resStr += string(strA[i])
		}
	}
	fmt.Printf("res:%v\n", string(resStr))
	res64, _ := strconv.ParseInt(resStr, 10, 64)
	if res64 > int64(INT32_MAX) || res64 < int64(INT32_MIN) {
		fmt.Printf("结果: %v不在int32范围内\n", res64)
	} else {
		fmt.Printf("结果: %v\n", int32(res64))
	}
}
