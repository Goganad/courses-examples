package main

import (
	"fmt"
	"time"
)

func main() {
	//go func() {
	//	fmt.Println("Hello, first goroutine!")
	//}()
	//
	//fmt.Println("Hello, main goroutine!")

	//runtime.GOMAXPROCS(1)
	//
	//var i int
	//for i < 10000 {
	//	go fmt.Print(0)
	//	fmt.Print(1)
	//	i++
	//}
	//
	//time.Sleep(2 * time.Second)

	//ch := make(chan int)
	//
	//go func() {
	//	var n int
	//	n++
	//	fmt.Println("goroutine: ", n)
	//	ch <- n
	//}()
	//
	//a := <-ch
	//
	//fmt.Println("main goroutine:", a)

	//ch := make(chan int)
	//fmt.Println("сейчас будет запись в канал")
	//ch <- 1
	//fmt.Println("этот код не выполнится")

	//ch := make(chan int)
	//fmt.Println("Сейчас будет чтение из канала")
	//<-ch
	//fmt.Println("Никогда не выполнится")

	//ch := make(chan int)
	//go write(ch)
	//res := read(ch)
	//fmt.Println(res)

	//ch := make(chan int, 2)
	//ch <- 1
	//ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//ch := make(chan int)
	//
	//go func() {
	//	for i := range []int{1, 2, 3, 4, 5} {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	//
	//for i := range ch {
	//	fmt.Println(i)
	//}

	//result := make(chan int)
	//stop := make(chan struct{})
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	for i := 0; i < 3; i++ {
	//		result <- i
	//	}
	//	stop <- struct{}{}
	//}()
	//
	//for {
	//	select {
	//	case <-stop:
	//		fmt.Println("завершение работы главной горутины")
	//		return
	//	case v := <-result:
	//		fmt.Println("результат:", v)
	//	default:
	//		fmt.Println("ожидание результата...")
	//		time.Sleep(time.Second)
	//	}
	//}

	ch := make(chan int)
	stop := make(chan struct{}, 2)

	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)
			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()

	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()

	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	fmt.Println("завершение работы главной горутины")
}

func read(ch <-chan int) int {
	fmt.Println("чтение из канала")
	return <-ch
}

func write(ch chan<- int) {
	fmt.Println("запись в канал 7")
	ch <- 7
}
