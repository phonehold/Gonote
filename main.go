/*
 * @Author: yingcai.huang
 * @Date: 2022-02-08 15:03:58
 * @LastEditTime: 2022-02-08 16:55:14
 * @LastEditors: yingcai.huang
 * @Description:
 * @FilePath: /Gonote/main.go
 * 引用代码请说明出处
 */
package main

import (
	"fmt"
	"gonote/lvar"
	"gonote/note"
)

func main() {
	note.SayHelloWorld()
	lvar.Lvar1()
	fmt.Println("Me")
	fmt.Println(lvar.Version)
}
