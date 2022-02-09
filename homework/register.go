/*
 * @Author: yingcai.huang
 * @Date: 2022-02-09 13:28:20
 * @LastEditTime: 2022-02-09 13:50:50
 * @LastEditors: yingcai.huang
 * @Description:
 * @FilePath: /Gonote/homework/register.go
 * 引用代码请说明出处
 */

package homework

import "fmt"

func Regisger() {
	for {

		var username, pass, vpass string
		fmt.Print("UserName:")
		fmt.Scanf("%s\n", &username)
		fmt.Print("Please input pass:")
		fmt.Scanf("%s\n", &pass)
		fmt.Print("Please conform pass again:")
		fmt.Scanf("%s\n", &vpass)
		if username == "" || pass == "" || vpass == "" {
			println("must not be nill")
			continue
		}
		if pass != vpass {
			println("pass not match vpass")
			continue
		}
		fmt.Println("Register Success")
		break
	}
}
