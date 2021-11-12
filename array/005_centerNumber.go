package array

import (
	"fmt"
	"sort"
)

//  005
//	给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和 nums2。
//	请你找出并返回这两个正序数组的 中位数 。
//	算法的时间复杂度应该为 O(log (m+n)) 。



func centerNumberInArray(nums1,nums2 []int) {
	fmt.Printf("nums1 :%v\n nums2 :%v\n",nums1,nums2)
	len1 := len(nums1)
	len2 := len(nums2)
	res1 := merge1(nums1,nums2)
	fmt.Printf("方法一: res:%v\n",res1)
	res2 := merge2(nums1,len1,nums2,len2)
	fmt.Printf("方法二: res:%v\n",res2)
	if len(res2) % 2 == 0 {
		fmt.Printf("合并数组长度:%d,除以2:%d,中位数为:%d,%d\n",len(res2),len(res2)/2,res2[len(res2)/2-1],res2[len(res2)/2])
	} else {
		fmt.Printf("合并数组长度:%d,除以2:%d,中位数为:%d\n",len(res2),len(res2)/2,res2[len(res2)/2])
	}
	return
}


// 合并两个有序数组方案一  合并重新排序
func merge1(nums1 []int,nums2 []int) []int {
	nums1 = append(nums1,nums2...)
	sort.Ints(nums1)
	return nums1
}
// 合并两个有序数组方案二 双指针 正排
func merge2(nums1 []int,len1 int,nums2 []int,len2 int) []int {
	temp := make([]int,len1+len2)
	i,j,k := 0,0,0
	for i<len1 && j< len2 {
		if nums1[i] < nums2[j] {
			temp[k] = nums1[i]
			i ++
		} else {
			temp[k] = nums2[j]
			j++
		}
		k ++
		//fmt.Printf("i:%d,len1:%d,j:%d,len2:%d,k:%d\n",i,len1,j,len2,k)
	}
	if i  < len1 {
		copy(temp[k:],nums1[i:])
	}
	if j < len2 {
		copy(temp[k:],nums2[j:])
	}
	copy(nums1,temp)
	return temp
}