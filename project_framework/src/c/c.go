package main

import(
	"a/a1"		// a1.go 里面：package a1
	"a/a2"		// a2.go 里面：package a2
	"b"		// 注意：这是文件夹名
)

func main(){
	a1.PrintA1()
	a2.PrintA2()
	b3.PrintB1()	// 文件夹b里面的包b3
	b3.PrintB2()
}
