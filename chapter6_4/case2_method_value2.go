package chapter6_4

import "time"

// 在一个包的API需要一个函数值、且调用方希望操作的是某一个绑定了对象的方法的话，方法"值"会非常实用
// 下面例子中的time.AfterFunc这个函数的功能是在指定的延迟时间之后来执行一个函数。且这个函数操作的是一个Rocket对象r

type Rocket struct {
	/* ... */
}

func (r *Rocket) Launch() {
	/* ... */
}

func MethodValue2() {
	r := new(Rocket)

	// 不使用方法值的一般用法
	time.AfterFunc(10*time.Second, func() { r.Launch() })

	// 直接用方法"值"传入AfterFunc的话可以更为简短：
	time.AfterFunc(10*time.Second, r.Launch)
}
