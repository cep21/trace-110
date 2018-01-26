package trace_110

import (
	"testing"
	"sync"
	"fmt"
	"time"
	"net/http"
)

func startServer() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	start := time.Now()
	go func() {
		defer wg.Done()
		j := 3
		for time.Since(start) < time.Second {
			for i :=1; i < 1000000; i++ {
				j *= i
			}
		}
		fmt.Println(j)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
	}()
	go func() {
		defer wg.Done()
		errCount := 0
		for time.Since(start) < time.Second * 4{
			_, err := http.Get("https://www.google.com")
			if err != nil {
				errCount++
			}
		}
		fmt.Println(errCount)
	}()
	wg.Wait()
}


func TestServer(t *testing.T) {
	startServer()
}