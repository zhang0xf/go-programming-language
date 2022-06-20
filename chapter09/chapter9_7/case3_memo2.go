package chapter9_7

import "sync"

type Memo2 struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

func New2(f Func) *Memo2 {
	return &Memo2{f: f, cache: make(map[string]result)}
}

// Get is concurrency-safe.
func (memo *Memo2) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
