package main

import (
	"fmt"
	"proto"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
)

func test_write() {
	var p1 = tutorial.Person {
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*tutorial.Person_PhoneNumber{
			{Number: "555-4321", Type: tutorial.Person_HOME},
		},
	}
	var p2 = tutorial.Person {
		Id:    5678,
		Name:  "Rob pike",
		Email: "rob@example.com",
		Phones: []*tutorial.Person_PhoneNumber{
			{Number: "666-abcd", Type: tutorial.Person_MOBILE},
		},
	}

	// 地址簿加人
	book := &tutorial.AddressBook{}
	book.People = append(book.People, &p1)
	book.People = append(book.People, &p2)

	// 写文件
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("./write.bin", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func test_read() {
	// 读文件
	in, err := ioutil.ReadFile("./write.bin")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("./write.bin: File not found.  Creating new file.\n")
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	// 解析
	book := &tutorial.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	// 输出内容
	peoples := book.GetPeople()
	for i:=0; peoples != nil && i < len(peoples); i++ {
		log.Println("name: ", peoples[i].Name)
	}
}

func main() {
	test_write()
	test_read()

	fmt.Println("hello, world")
}
