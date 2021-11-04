package chapter4_2

import "fmt"

type T int

func Slice() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "Macth",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	fmt.Println(summer[:20]) // panic: out of range

	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"
	// 注:获得summer切片的第0个位置到第5个位置(超出了长度3,没有超出capcity,扩展了summer)

	var s []int    // len(s) == 0, s == nil
	s = nil        // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{}    // len(s) == 0, s != nil
	// 注:判断一个slice是否为空,使用len(s) == 0来判断,而不该使用s == nil,见上文:s = []int{}
	fmt.Printf("%T", s)

	// make创建匿名数组,并返回一个slice
	len := 3
	cap := 5
	t1 := make([]T, len)
	t2 := make([]T, len, cap) // same as make([]T, cap)[:len]
	fmt.Printf("%T", t1)
	fmt.Printf("%T", t2)
}

// 判断两个slice是否相等(深相等,需手动实现)
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
