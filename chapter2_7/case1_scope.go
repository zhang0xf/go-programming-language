package chapter2_7

import (
	"fmt"
	"log"
	"os"
)

var g = "g"

func f() {}

func ScopeTest1() {
	f := "f"
	fmt.Println(f) // "f"; local var f shadows package-level func f
	fmt.Println(g) // "g"; package-level var
	// fmt.Println(h) // compile error: undefined: h
}

func ScopeTest2() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}
}

func ScopeTest3() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
}

func f1() int {
	return 0
}

func g1(x int) int {
	return 0
}

func ScopeTest4() {
	if x := f1(); x == 0 {
		fmt.Println(x)
	} else if y := g1(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, y) // compile error: x and y are not visible here
}

func ScopeTest5() error {
	var fname string
	// f只在if语句可见
	// if f, err := os.Open(fname); err != nil { // compile error: unused: f
	// 	return err
	// }
	// f.ReadByte() // compile error: undefined f
	// f.Close()    // compile error: undefined f

	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	// f.ReadByte()
	f.Close()

	return nil
}

var cwd string

func init() {
	// cwd内部申明屏蔽了外部申明
	cwd, err := os.Getwd() // compile error: unused: cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	// 全局变量cwd并不会被正确地初始化,看似正常的日志输出会使得这个bug更加隐晦!
	log.Printf("Working directory = %s", cwd)
}

func init() {
	// 单独申明err,以防止:=的简短申明方式
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
