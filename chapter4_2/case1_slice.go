package chapter4_2

import "fmt"

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
}
