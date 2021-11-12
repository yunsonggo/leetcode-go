package numbers

import "fmt"

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 数字数组 两两下标相减 乘以 二者中的小者 得出结果最大的两数

func mostWater(arr []int) {
	var (
		max  int
		numA int
		numB int
	)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			// 两下标只差为 x轴长度
			long := i - j
			tempMax := 0
			// 两者小值
			temp := 0
			if arr[i]-arr[j] <= 0 {
				temp = arr[i]
			} else {
				temp = arr[j]
			}

			if long <= 0 {
				long = -long
				tempMax = long * temp
			} else {
				tempMax = long * temp
			}
			if tempMax > max {
				max = tempMax
				numA = arr[i]
				numB = arr[j]
			}
		}
	}
	fmt.Printf("最大容量:%d,numA:%d,numB:%d\n", max, numA, numB)
}
