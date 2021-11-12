package linkedList

import (
	"fmt"
	"strconv"
)

//给你两个非空的链表，表示两个非负的整数。
//它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//链接：https://leetcode-cn.com/problems/add-two-numbers

type numberList struct {
	value int
	next *numberList
}

//numberListSum.value = sum / dis % 10

// num1,num2 := 342,465
func towLinkSum(num1,num2 int) {
	sum := num1 + num2
	num1Str := strconv.Itoa(num1)
	sumStr := strconv.Itoa(sum)

	numberListOne := &numberList{}
	numberListTwo := &numberList{}
	numberListSum := &numberList{}

	for i:=0;i<len(num1Str);i++ {
		dis := divisor(i)
		fmt.Printf("i: %d,dis:%d\n",i,dis)
		n1 := &numberList{
			value: num1 / dis % 10,
		}
		n2 := &numberList{
			value: num2 / dis % 10,
		}
		insertNode(numberListOne,n1)
		insertNode(numberListTwo,n2)
	}

	for i:=0;i<len(sumStr);i++ {
		dis := divisor(i)
		fmt.Printf("i: %d,dis:%d\n",i,dis)
		s := &numberList{
			value: sum / dis % 10,
		}
		insertNode(numberListSum,s)
	}

	fmt.Printf("----------第一个数字是:%d 逆向单链表是:------------\n",num1)
	showList(numberListOne)
	fmt.Printf("----------第二个数字是:%d 逆向单链表是:------------\n",num2)
	showList(numberListTwo)
	fmt.Printf("----------和是:%d 逆向单链表是:------------\n",sum)
	showList(numberListSum)
}


// 添加一个节点
func insertNode(headNode ,newNode *numberList) {
	temp := headNode
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
}

// 显示链表
func showList(headNode *numberList) {
	temp := headNode
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	result := ""
	for {
		if temp.next == nil {
			break
		}
		result += fmt.Sprintf("%d",temp.next.value)
		if temp.next.next != nil {
			result += "->"
		}
		temp = temp.next
	}
	fmt.Println(result)
}

// 删除一个节点
func removeNode(headNode *numberList,value int) {
	temp := headNode
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.value == value {
				flag = true
				break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("没有找到要删除的值")
	}
}

func divisor(num int) int {
	a := "1"
	if num == 0 {
		return 1
	}
	for i:=1;i<=num;i++ {
		a += "0"
	}
	res,_ := strconv.Atoi(a)
	return res
}