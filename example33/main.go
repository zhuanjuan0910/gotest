package main

import "fmt"

func main() {
	fmt.Println(doubleScore(0.0))  //0
	fmt.Println(doubleScore(20.0)) //40
	fmt.Println(doubleScore(50.0)) //50
}
func doubleScore(source float32) (score float32) {
	defer func() { //在defer中修改返回值的值
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	score = source * 2
	return
}
