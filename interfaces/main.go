package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"strconv"
)

type Formatter interface {
	Int()
	Float()
}

type myType float32

func (mt *myType) Int() {
	fmt.Printf("Целое число: %.f", mt)
	fmt.Println()
}

func (mt *myType) Float() {}

type Temperature interface {
	Show()
}

type Kelvin float32

func (k Kelvin) Show() {
	fmt.Printf("Температура: %.2f K", k)
	fmt.Println()
}

func main() {
	err3 := fmt.Errorf("ошибка3:%w", fmt.Errorf("ошибка2:%w", errors.New("ошибка1")))
	fmt.Println(errors.Is(err3, errors.New("ошибка1")))

	err31 := errors.New("ошибка1")
	err3 = fmt.Errorf("ошибка3:%w", fmt.Errorf("ошибка2:%w", err31))
	fmt.Println(errors.Is(err3, err31))

	var a interface{}
	fmt.Printf("Тип: %T, значение: %v, f == nil -> %v\n", a, a, a == nil)

	a = struct {
		a int
		b int
	}{
		a: 6,
		b: 7,
	}
	fmt.Printf("Тип: %T, значение: %v, f == nil -> %v\n", a, a, a == nil)

	a = []int{1, 2, 3, 4}
	fmt.Printf("Тип: %T, значение: %v, f == nil -> %v\n", a, a, a == nil)

	switch a.(type) {
	case map[string]int:
		fmt.Printf("it is a map of [string]int! %v\n", a)
	case map[string]string:
		fmt.Printf("it is a [string]string! %v\n", a)
	case map[int]int:
		fmt.Printf("it is a [string]string! %v\n", a)
	default:
		fmt.Println("unsupported type")
	}

	formatter, _ := getFormatter()
	if formatter == nil {
		log.Fatal("bad")
	}
}

func run(t Temperature) {
	fmt.Printf("Тип: %T, значение: %v", t, t)
	fmt.Println()
}

func run1(mt myType) {
	fmt.Printf("run1: %+v", mt)
	fmt.Println()
}

func run2(f Formatter) {
	fmt.Printf("run2: %+v", f)
	fmt.Println()
}

func run3(w io.Writer) {
	n, err := w.Write([]byte{})
	if err != nil {
		return
	}
	fmt.Printf(strconv.Itoa(n))
}

func getFormatter() (Formatter, error) {
	var f Formatter

	if false {
		var mt *myType
		f = mt
	}

	fmt.Printf("Тип: %T, значение: %v, f == nil -> %v\n", f, f, f == nil)

	var mt *myType
	f = mt

	fmt.Printf("Тип: %T, значение: %v, f == nil -> %v\n", f, f, f == nil)

	var d myType = 5
	mt = &d
	f = mt

	fmt.Printf("Тип: %T, значение: %v, f == nil -> %v\n", f, f, f == nil)

	return f, nil
}
