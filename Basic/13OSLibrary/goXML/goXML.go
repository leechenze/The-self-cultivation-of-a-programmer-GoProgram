package goXML

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func GoXML() {
	println("========================OS Library========================")
	println()

	// MarshalIndent And Marshal
	println()
	marshalIndentFn()

	// Unmarshal
	println()
	UnmarshalFn()

	// Decoder结合IO读取Json文件
	println()
	decoderFn()

	// Encoder结合IO写入Json文件
	println()
	encoderFn()

	println()
	println("========================OS Library========================")

}

func encoderFn() {
	person := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	file, _ := os.OpenFile("operate/write.xml", os.O_WRONLY, os.ModePerm)
	defer file.Close()
	encoder := xml.NewEncoder(file)
	encoder.Encode(person)
}

func decoderFn() {
	file, _ := ioutil.ReadFile("operate/read.xml")
	var person Person
	xml.Unmarshal(file, &person)
	fmt.Printf("decoder person: %v\n", person)
}

func UnmarshalFn() {
	strXML := `
		<person>
				<name>tom</name>
				<age>20</age>
				<email>tom@gmail.com</email>
		</person>
	`
	strXMLByte := []byte(strXML)
	var person Person
	xml.Unmarshal(strXMLByte, &person)
	fmt.Printf("xmlUnmarshal: %v\n", person)
}

func marshalIndentFn() {
	person := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	xmlMarshal, _ := xml.MarshalIndent(person, "", " 	")
	fmt.Printf("xmlMarshal: %v\n", string(xmlMarshal))
}

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}
