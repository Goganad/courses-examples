package main

import (
	"encoding/json"
	"fmt"
)

/*
Необходимо представить в виде json структуру contract
contract{
Number: 2,
Landlord: "Остап",
tenat: "Шура",
}
*/
type contract struct {
	Number   int    `json:"number"`
	SignDate string `json:"-"`
	Landlord string `json:"landlord"`
	Tenat    string `json:"tenat"`
}

func main() {
	c := contract{
		Number:   2,
		SignDate: "2023-09-01",
		Landlord: "Остап",
		Tenat:    "Шура",
	}
	//fmt.Printf("структура %+v", c)
	//fmt.Println()
	res, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%v", string(res))
	//fmt.Println()
}
