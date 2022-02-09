/*
 * @Author: yingcai.huang
 * @Date: 2022-02-08 15:03:58
 * @LastEditTime: 2022-02-09 14:44:49
 * @LastEditors: yingcai.huang
 * @Description:
 * @FilePath: /Gonote/main.go
 * 引用代码请说明出处
 */
package main

import (
	"fmt"
	"gonote/lfunc"
	"gonote/logger"
	"gonote/lpoint"
	"gonote/lstring"
	"gonote/lvar"
	"gonote/note"
)

func main() {

	note.SayHelloWorld()
	lvar.Lvar1()
	fmt.Println("Me")
	fmt.Println(lvar.Version)
	lstring.SayLstring()
	lpoint.Increase(lpoint.Src)
	logger.Newmain()
	// homework.Regisger()
	//lable.Rlable()
	res1, res2 := lfunc.FuncSum(100, 7)
	fmt.Printf("res1 %v, res2 %v", res1, res2)
}