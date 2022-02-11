/*
 * @Author: yingcai.huang
 * @Date: 2022-02-11 16:13:43
 * @LastEditTime: 2022-02-11 16:24:30
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
}
