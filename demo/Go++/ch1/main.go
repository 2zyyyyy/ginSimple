/*
@Description：test.go
@Author : gilbert
@Date : 2022/5/9 3:42 PM
*/

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
	练习1.1:修改echo程序输出os.Args[0]，即命令的名字。
	练习1.2:修改echo程序，输出参数的索引和值，每行-一个。
	练习1.3:尝试测量可能低效的程序和使用strings.Join的程序在执行时间上的差异。
*/

func exercise1() {
	fmt.Println("os.args[0]:", os.Args[0])
}

func exercise2() {
	for k, v := range os.Args[1:] {
		fmt.Println(k, v)
	}
}

// join
func exercise3() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	nanoseconds := time.Since(start)
	fmt.Println(nanoseconds)
}

// 字符串拼接 for range
func exercise4() {
	start := time.Now()
	s, sep := "", ""
	for _, v := range os.Args[1:] {
		sep = " "
		s += v + sep
	}
	fmt.Println(s)
	nanoseconds := time.Since(start)
	fmt.Println(nanoseconds)
}

// 字符串拼接 for i
func exercise5() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	nanoseconds := time.Since(start)
	fmt.Println(nanoseconds)
}

func main() {
	exercise1()
	exercise2()
	exercise3() // 1.842µs
	exercise4() // 1.852µs
	exercise5() // 1.477µs
}
