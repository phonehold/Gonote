/*
 * @Author: yingcai.huang
 * @Date: 2022-02-09 13:55:08
 * @LastEditTime: 2022-02-09 14:27:05
 * @LastEditors: yingcai.huang
 * @Description:
 * @FilePath: /Gonote/lable/rlable.go
 * 引用代码请说明出处
 */
package lable

import "fmt"

func Rlable() {
	//
outside:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("#  ")
			if i == 9 && j == 4 {
				break outside

			}
		}
		fmt.Println()

	}
}
