package chapter4_4

import (
	"fmt"
	"time"
)

// 一个命名为S的结构体类型将不能再包含S类型的成员(该限制同样适用于数组)
// 但是S类型的结构体可以包含*S指针类型的成员，这可以让我们创建递归的数据结构，比如链表和树结构等。

// 如果结构体成员名字是以大写字母开头的，那么该成员就是导出的；这是Go语言导出规则决定的。一个结构体可能同时包含导出和未导出的成员。
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// 结构体成员的输入顺序也有重要的意义。
// 交换Name和Address出现的先后顺序，那样的话就是定义了不同的结构体类型。
type Employee_short struct {
	ID            int
	Name, Address string // 合并
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func Struct() {

	var dilbert Employee

	// .
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code

	// 取地址
	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia

	// 取地址 + .
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	(*employeeOfTheMonth).Position += " (proactive team player)"

	// 函数返回结构体
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

	// 更新值
	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no real reason

}

// EmployeeByID函数将根据给定的员工ID返回对应的员工信息结构体的指针
// 如果将EmployeeByID函数的返回值从*Employee指针类型改为Employee值类型，那么更新语句将不能编译通过，因为在赋值语句的左边并不确定是一个变量（译注：调用函数返回的是值，并不是一个可取地址的变量）
func EmployeeByID(id int) *Employee {
	/* ... */
	// 省略
	return &Employee{}
}
