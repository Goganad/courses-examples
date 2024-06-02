package main

import (
	"context"
	"fmt"
	"time"
)

//func main() {
//	ctx := context.Background()
//	d := time.Now().Add(-2 * time.Second)
//	ctx, _ = context.WithDeadline(ctx, d)
//
//	wg := sync.WaitGroup{}
//	for i := 0; i < 10; i++ {
//		i := i
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for {
//				select {
//				case <-ctx.Done():
//					fmt.Println("стоп горутина: ", i)
//					return
//				default:
//					fmt.Println("сложные вычисления горутины: ", i)
//				}
//			}
//		}()
//	}
//
//	wg.Wait()
//}

//type ctxKey string
//
//const (
//	ctxKeyProgramName = "program_name"
//	ctxKeyTime        = "time"
//)
//
//func main() {
//	ctx := context.Background()
//	ctx = context.WithValue(ctx, ctxKeyProgramName, "some program name")
//	do(ctx)
//}
//
//func do(ctx context.Context) {
//	fmt.Println("1:", ctx.Value(ctxKeyProgramName))
//}

func main() {
	start := time.Now()
	ctx := context.Background()
	d := time.Now().Add(10 * time.Second)
	ctx, cancel1 := context.WithDeadline(ctx, d)
	defer cancel1()
	d = time.Now().Add(time.Second)
	ctx, cancel2 := context.WithDeadline(ctx, d)
	defer cancel2()
OUT:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("завершение работы")
			break OUT
		default:
			fmt.Println("работа")
			time.Sleep(time.Second)
		}
	}

	duration := time.Since(start)
	fmt.Println("Время выполнения:", duration)
}
