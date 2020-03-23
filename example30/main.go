package main

import "fmt"

func main(){
	test()
}
func test(){
	defer func(){
		fmt.Println("打印前")
	}()
	defer func(){
		fmt.Println("打印中")
	}()
	defer func(){
		fmt.Println("打印后")
	}()
	panic("打印异常")

}
// 打印后
// 打印中
// 打印前
// 打印异常  defer中先进后出
