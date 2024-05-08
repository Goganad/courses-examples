package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type contract struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date,omitempty"`
	Landlord string `xml:"-"`
	tenat    string `xml:"tenat"`
}

type contracts struct {
	List []contract `xml:"contract"`
}

func main() {
	d := `
<?xml version="1.0" encoding="UTF-8"?>
<contracts>
	<contract>
		<number>1</number>
		<sign_date>2023-09-01</sign_date>
		<landlord>Остап Бендер</landlord>
		<tenat>Шура Балаганов</tenat>
	</contract>
	<contract>
		<number>1</number>
		<sign_date>2023-09-01</sign_date>
		<landlord>Остап Бендер</landlord>
		<tenat>Шура Балаганов</tenat>
	</contract>
</contracts>
`

	c := contracts{}
	err := xml.Unmarshal([]byte(d), &c)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", c)
	fmt.Println()

	fileName := "contracts.xml"
	//f, err := os.Create(fileName)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//n, err := f.WriteString(xml.Header)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//fmt.Println(n)

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	res := contracts{}
	err = xml.NewDecoder(f).Decode(&res)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
