package chapter5_4

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 错误处理策略:
// 2. 重新尝试失败操作(错误的发生是偶然性的):限制重试时间和次数

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off(指数退避?:左移*2): 意味着每次沉睡的时间会越来越长
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
