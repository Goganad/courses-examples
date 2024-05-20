package main

import (
	"fmt"
	"sync"
	"time"
)

var loadOnce sync.Once

func main() {
	wg := sync.WaitGroup{}

	n := 5

	wg.Add(n)
	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()
			fmt.Println("горутина No", i)
		}()
	}

	timeBackup := time.Now()

	wg.Wait()

	fmt.Printf("we were waiting 5 goroutines for %f", time.Now().Sub(timeBackup).Seconds())

	wg = sync.WaitGroup{}

	statusCode := 0

	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			loadOnce.Do(func() {
				status := load()
				statusCode = status
			})
			fmt.Println("горутина: ", i)
		}()
	}

	wg.Wait()

	fmt.Println(statusCode)
}

func load() int {
	time.Sleep(3 * time.Second)
	fmt.Println("load")

	return 404
}
