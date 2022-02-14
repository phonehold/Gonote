/*
 * @Author: yingcai.huang
 * @Date: 2022-02-11 16:58:24
 * @LastEditTime: 2022-02-14 19:04:24
 * @LastEditors: yingcai.huang
 * @Description:
 * @FilePath: /Gonote/lslice/lslice.go
 * 引用代码请说明出处
 */
package lslice

import "fmt"

func NewSlice() {
	//TODO fix me
	s3 := make([]int, 4, 5)
	fmt.Println("s3=", s3)
	// //define slice var s1
	// var s1 []int
	// //define slice s2
	// s2 := []int{}
	// //define array
	// var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// //init slice s3
	// var s3 []int = make([]int, 0)

	// 	全局：
	// var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// var slice0 []int = arr[start:end]
	// var slice1 []int = arr[:end]
	// var slice2 []int = arr[start:]
	// var slice3 []int = arr[:]
	// var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
	// 局部：
	// arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	// slice5 := arr[start:end]
	// slice6 := arr[:end]
	// slice7 := arr[start:]
	// slice8 := arr[:]
	// slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素

}
