package main

import "fmt"

func main() {
	s := make([]int, 5) //长度为5的切片为空时，值为该类型的0值
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

//输出[0 0 0 0 0 1 2 3]
