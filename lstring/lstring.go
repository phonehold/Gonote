/*
 * @Author: yingcai.huang
 * @Date: 2022-02-08 21:13:02
 * @LastEditTime: 2022-02-08 21:26:47
 * @LastEditors: yingcai.huang
 * @Description:
 * @FilePath: /Gonote/lstring/lstring.go
 * 引用代码请说明出处
 */
package lstring

import "fmt"

func SayLstring() {
	var (
		s = "Hello "
	)
	fmt.Println(s + "Google")
	fmt.Printf("The length is %v", len(s+"Google"))
}
