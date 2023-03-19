package tools

import (
	"fmt"
	"sync"
	"time"
)

var requests = make(map[string]interface{})

func RefreshRequestsData(errch chan<- error, mutex *sync.Mutex) {
	for {
		for key, _ := range requests {
			req, err := ResponseMap(key)
			if err != nil {
				errch <- err
			} else {
				mutex.Lock()
				requests[key] = req
				mutex.Unlock()
			}
		}
		time.Sleep(time.Hour * 2)
	}
}

func ErrorHandler(errch <-chan error) {
	for err := range errch {
		fmt.Print(err)
	}
}

func GarbageCollector() {
	for {
		time.Sleep(time.Hour * 24)
		for key, _ := range requests {
			delete(requests, key)
		}
	}
}
