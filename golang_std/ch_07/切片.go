package main

import (
	"fmt"
	"sort"
)

type Animal struct {
	person string
	height int
	weight int
}

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	//长度是切片中元素的个数，容量是切片第一个元素到切片所指向的底层数组中的末尾元素的个数
	s1 := arr[0:3]
	printSlice(s1)
	s1 = append(s1, 6)
	printSlice(s1)
	fmt.Println(arr)
	//追加元素后，没有发生扩容，容量与指针地址没有变化，底层数组发生了变化
	//[1 2 3 4 5]
	//len=3 cap=5 0xc0000ca030 [1 2 3]
	//len=4 cap=5 0xc0000ca030 [1 2 3 6]
	//[1 2 3 6 5]

	arr2 := [5]int{1, 2, 3, 4}
	fmt.Println(arr2)
	//长度是切片中元素的个数，容量是切片第一个元素到切片所指向的底层数组中的末尾元素的个数
	s2 := arr2[2:]
	printSlice(s2)
	s2 = append(s2, 6)
	printSlice(s2)
	fmt.Println(arr2)
	//当切片的底层数组不足以容纳所有给定值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组
	//这个切片在追加元素后，发生了扩容，其容量和指针地址发生了变化，指向了心得底层数组，原有底层数组未变
	//[1 2 3 4 0]
	//len=3 cap=3 0xc00000c400 [3 4 0]
	//len=4 cap=6 0xc00000c450 [3 4 0 6]
	//[1 2 3 4 0]

	//切片扩容的规律，当长度与容量一样时，追加元素，长度加1，容量变成原来的两倍

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %p %v\n", len(s), cap(s), s, s)
}

func sliceSort() {

	mySlice := []Animal{}

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})

}
