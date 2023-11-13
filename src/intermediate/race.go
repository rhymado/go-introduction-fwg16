package intermediate

import (
	"fmt"
	"sync"
)

var num = 0

func Racing() {
	var mtx = sync.RWMutex{}
	for num < 10 {
		wg.Add(2)

		mtx.RLock() // lock bersifat blocking jika mutex di lock
		go read(&mtx)

		mtx.Lock() // lock bersifat blocking jika mutex di lock
		go increment(&mtx)

		wg.Wait() // wait bersifat lock hingga counter wg 0
	}
}

func read(mtx *sync.RWMutex) {
	defer func() {
		mtx.RUnlock()
		wg.Done()
	}()
	fmt.Println(num)
}

func increment(mtx *sync.RWMutex) {
	defer func() {
		mtx.Unlock()
		wg.Done()
	}()
	num++
}
