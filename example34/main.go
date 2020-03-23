package main

import "fmt"

type Student struct{
	Name string
	Age int
}
func pase_student()map[string]*Student{
	m:=make(map[string]*Student)
	stus:=[]Student{
		{Name:"zhou", Age: 24 },
		{Name: "li", Age:23},
		{Name:"wang", Age:22},
	}
	for i,_:=range stus{
		stu:=stus[i]//每次循环都创建一个新的变量stu
		m[stu.Name]=&stu
	}
	return m
}
func main(){
	students:=pase_student()
	for k,v:=range students{
		fmt.Printf("key=%s,value=%v\n",k,v)
	}
}