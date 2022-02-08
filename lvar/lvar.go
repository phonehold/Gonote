/*
 * @Author: yingcai.huang
 * @Date: 2022-02-08 16:32:49
 * @LastEditTime: 2022-02-08 16:54:46
 * @LastEditors: yingcai.huang
 * @Description:
 * @FilePath: /Gonote/lvar/lvar.go
 * 引用代码请说明出处
 */
package lvar

import "fmt"

var V1 int = 10
var V2 int64 = 100

const Version = "v1.0.37"

var (
	V3 int  = 100
	V4 int8 = 6
	V5 int
	V6 float32
)

func Lvar1() {

	fmt.Printf("V1=%v", V1)
	fmt.Println("")
	fmt.Printf("V2=%v\n", V2)
	fmt.Printf("V3=%v\n V4=%v\n V5=%v\n V6=%v", V3, V4, V5, V6)
}
