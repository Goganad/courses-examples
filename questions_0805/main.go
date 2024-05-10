package main

import (
	"errors"
	"fmt"
	"sync"
)

///some/path/to/gopath/
//├── src/
//	├── myproject/
//		├── cmd/
//			└── main.go
//		├── tools/
//			└── arithmetic.go

type Bird interface {
	Sing() string
}

type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	if d != nil {
		return ""
	}

	return d.voice
}

func main() {
	d := NewDuck()
	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)

	americanSocket := &Adapter{
		EuropeanDevice: &EuropeanLamp{},
	}
	americanSocket.InsertIntoAmericanSocket()
}

// Answer to question number 7
func NewDuck() *Duck {
	return &Duck{
		voice: "krya",
	}
}

func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}

// Factory
type Database struct {
	ConnectionString string
}

func NewDatabase(connectionString string) *Database {
	return &Database{
		ConnectionString: connectionString,
	}
}

// singleton
type Singleton struct {
	Count int
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

// adapter
type EuropeanSocket interface {
	InsertIntoEuropeanSocket()
}

type AmericanSocket interface {
	InsertIntoAmericanSocket()
}

type Adapter struct {
	EuropeanDevice EuropeanSocket
}

func (a *Adapter) InsertIntoAmericanSocket() {
	a.EuropeanDevice.InsertIntoEuropeanSocket()
}

type EuropeanLamp struct{}

func (e *EuropeanLamp) InsertIntoEuropeanSocket() {
	fmt.Println("European Lamp connected to European socket.")
}
