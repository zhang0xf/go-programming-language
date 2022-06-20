// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package chapter9_7

// 下面是我们要设计的cache的第一个“草稿”：
// Memo实例会记录需要缓存的函数f(类型为Func)，以及缓存内容(里面是一个string到result映射的map)。
// 每一个result都是简单的函数返回的值对儿--一个值和一个错误值。
// 继续下去我们会展示一些Memo的变种，不过所有的例子都会遵循上面的这些方面。

// A Memo caches the results of calling a Func.
type Memo1 struct {
	f     Func
	cache map[string]result
}

func New1(f Func) *Memo1 {
	return &Memo1{f: f, cache: make(map[string]result)}
}

// NOTE: not concurrency-safe!
func (memo *Memo1) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
