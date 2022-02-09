/*
 * @Author: yingcai.huang
 * @Date: 2022-02-08 21:29:37
 * @LastEditTime: 2022-02-09 11:21:23
 * @LastEditors: yingcai.huang
 * @Description:
 * @FilePath: /Gonote/lpoint/lpoint.go
 * 引用代码请说明出处
 */
package lpoint

import "fmt"

var Src int = 2022

func Increase(n int) {
	NewFunction(Src)
	erwr := 1
	erwr++
	n++
	fmt.Printf("\r n的值%v\r\n n的地址%v\r\n", n, &n)
}

func NewFunction(n int) {
	csgoogle := 34
	csgoogle++
	fmt.Printf("\r \n newfun value %v", csgoogle)
}
