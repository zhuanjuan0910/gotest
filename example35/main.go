package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i:", i)
			wg.Done()
		}() //输出十个10，for循坏出来之后i就变成10了

	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i:", i)
			wg.Done()
		}(i) //输出0~9，参数之前就已经先计算好

	}
	wg.Wait()
}
