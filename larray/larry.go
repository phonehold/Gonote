/*
 * @Author: yingcai.huang
 * @Date: 2022-02-11 16:13:43
 * @LastEditTime: 2022-02-11 17:36:41
 * @LastEditors: yingcai.huang
 * @Description:
 * @FilePath: /Gonote/larray/larry.go
 * 引用代码请说明出处
 */
package larray

import "fmt"

func NewArray() {
	var a = [...]int{
		886, 9, 4, 180, 293, 512, 1024,
	}

	a[1] = 3
	a[5] = 6
	fmt.Println("\r\n Array value example")
	for i := 0; i < len(a); i++ {
		fmt.Printf("index %v value is %v\r", i, a[i])
	}
	fmt.Println("\r\n Array value range example")
	for i, v := range a {
		fmt.Printf("index %v value is %v\r", i, v)
	}
	var TwoArray [3][4]int = [3][4]int{
		{1, 2, 3, 4},
		{7, 8, 9, 0},
		{66, 77, 88, 99},
	}
	for i, v := range TwoArray {
		for i1, v2 := range v {
			fmt.Printf("a[%v][%v] value is %v\t\r", i, i1, v2)
		}
	}

}
