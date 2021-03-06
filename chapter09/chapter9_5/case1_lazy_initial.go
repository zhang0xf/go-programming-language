package chapter9_5

import (
	"image"
	"sync"
)

// 如果初始化成本比较大的话，那么将初始化延迟到需要的时候再去做就是一个比较好的选择。
// 如果在程序启动的时候就去做这类初始化的话，会增加程序的启动时间，并且因为执行的时候可能也并不需要这些变量，所以实际上有一些浪费。

var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

// 懒初始化(lazy initialization)
// NOTE: not concurrency-safe!
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons() // one-time initialization
	}
	return icons[name]
}

func loadIcon(name string) image.Image {
	// ...
	return nil
}

// 如果一个变量只被一个单独的goroutine所访问的话，我们可以使用上面的这种模板，但这种模板在Icon被并发调用时并不安全。
// 就像前面银行的那个Deposit(存款)函数一样，Icon函数也是由多个步骤组成的：首先测试icons是否为空，然后load这些icons，之后将icons更新为一个非空的值。
// 直觉会告诉我们最差的情况是loadIcons函数被多次访问会带来数据竞争。
// 当第一个goroutine在忙着loading这些icons的时候，另一个goroutine进入了Icon函数，发现变量是nil，然后也会调用loadIcons函数。不过这种直觉是错误的。
// 我们希望你从现在开始能够构建自己对并发的直觉，也就是说对并发的直觉总是不能被信任的！
// 回忆一下9.4节。因为缺少显式的同步，编译器和CPU是可以随意地去更改访问内存的指令顺序，以任意方式，只要保证每一个goroutine自己的执行顺序一致。
// 其中一种可能loadIcons的语句重排是下面这样。它会在填写icons变量的值之前先用一个空map来初始化icons变量。
func loadIcons1() {
	icons = make(map[string]image.Image)
	icons["spades.png"] = loadIcon("spades.png")
	icons["hearts.png"] = loadIcon("hearts.png")
	icons["diamonds.png"] = loadIcon("diamonds.png")
	icons["clubs.png"] = loadIcon("clubs.png")
}

// 因此，一个goroutine在检查icons是非空时，也并不能就假设这个变量的初始化流程已经走完了。
// 最简单且正确的保证所有goroutine能够观察到loadIcons效果的方式，是用一个mutex来同步检查。
var mu sync.Mutex // guards icons

// Concurrency-safe.
func Icon2(name string) image.Image {
	mu.Lock()
	defer mu.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

// 然而使用互斥访问icons的代价就是没有办法对该变量进行并发访问，即使变量已经被初始化完毕且再也不会进行变动。这里我们可以引入一个允许多读的锁：
var rwMu sync.RWMutex // guards icons

// Concurrency-safe.
func Icon3(name string) image.Image {
	rwMu.RLock()
	if icons != nil {
		icon := icons[name]
		rwMu.RUnlock()
		return icon
	}
	rwMu.RUnlock()

	// acquire an exclusive lock
	rwMu.Lock()
	if icons == nil { // NOTE: must recheck for nil
		loadIcons()
	}
	icon := icons[name]
	rwMu.Unlock()
	return icon
}

// 上面的代码有两个临界区。goroutine首先会获取一个读锁，查询map，然后释放锁。如果条目被找到了(一般情况下)，那么会直接返回。
// 如果没有找到，那goroutine会获取一个写锁。不释放共享锁的话，也没有任何办法来将一个共享锁升级为一个互斥锁，所以我们必须重新检查icons变量是否为nil，以防止在执行这一段代码的时候，icons变量已经被其它gorouine初始化过了。

// 上面的模板使我们的程序能够更好的并发，但是有一点太复杂且容易出错。
// 幸运的是，sync包为我们提供了一个专门的方案来解决这种一次性初始化的问题：sync.Once。
// 概念上来讲，一次性的初始化需要一个互斥量mutex和一个boolean变量来记录初始化是不是已经完成了；互斥量用来保护boolean变量和客户端数据结构。
// Do这个唯一的方法需要接收初始化函数作为其参数。让我们用sync.Once来简化前面的Icon函数吧：
var loadIconsOnce sync.Once

// Concurrency-safe.
func Icon4(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

// 每一次对Do(loadIcons)的调用都会锁定mutex，并会检查boolean变量。
// 在第一次调用时，boolean变量的值是false，Do会调用loadIcons并会将boolean变量设置为true。
// 随后的调用什么都不会做，但是mutex同步会保证loadIcons对内存产生的效果能够对所有goroutine可见。
// 用这种方式来使用sync.Once的话，我们能够避免在变量被构建完成之前和其它goroutine共享该变量。
