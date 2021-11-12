package algorithm

import (
	"fmt"
	"math/rand"
	"time"
)

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
type answer struct {
	a num
	b num
}

type num struct {
	key   int
	value int
}

func findTowNum() {
	var answers []answer
	var ans answer
	var nums []int

	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(1000)
	fmt.Printf("-------------\n this round the target's : %d\n --------------\n",target)

	for i := 1; i <= 100; i++ {
		x := rand.Intn(1000)
		nums = append(nums, x)
	}
	for i := 0; i < len(nums) - 1; i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans.a.key = i
				ans.a.value = nums[i]
				ans.b.key = j
				ans.b.value = nums[j]
				answers = append(answers, ans)
				break
			}
		}
	}

	if len(answers) == 0 {
		fmt.Println("no answer")
	} else {
		for _, v := range answers {
			fmt.Printf("num1's key = %d,value=%d,num2's key = %d,value = %d,sum=%d\n", v.a.key, v.a.value, v.b.key, v.b.value, v.a.value+v.b.value)
		}
	}
	return
}
