package chapter9_7

import "sync"

type Memo3 struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

func New3(f Func) *Memo3 {
	return &Memo3{f: f, cache: make(map[string]result)}
}

func (memo *Memo3) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the two critical sections, several goroutines
		// may race to compute f(key) and update the map.
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
