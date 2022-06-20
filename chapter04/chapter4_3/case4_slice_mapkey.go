package chapter4_3

import "fmt"

// slice作为map的key,但是slice并不满足可比较的特性

var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
