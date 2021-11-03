package chapter3_6

type Weekday int

// itoa常量生成器
const (
	Sunday Weekday = iota // 第一个声明的常量所在的行,itoa置为0,然后在每一个有常量声明的行+1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

//  较为复杂的itoa常量生成器:用于给一个无符号整数的最低5bit的每个bit指定一个名字
type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

// 每个常量都是1024的幂
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)
