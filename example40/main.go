package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	userAges := &UserAges{
		ages: make(map[string]int),
	}
	for i := 0; i < 10000; i++ {
		go userAges.Add("oldboy", i)
		go func() {
			age := userAges.Get("oldboy")
			fmt.Println(age)
		}()
	}
}
