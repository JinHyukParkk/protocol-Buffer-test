package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/JinHyukParkk/protocol_Buffer_test/tutorial"
	"github.com/gogo/protobuf/proto"
)

func writePerson(w io.Writer, p *pb.Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "  Name:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "  E-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case pb.Person_MOBILE:
			fmt.Fprint(w, "  Mobile phone #: ")
		case pb.Person_HOME:
			fmt.Fprint(w, "  Home phone #: ")
		case pb.Person_WORK:
			fmt.Fprint(w, "  Work phone #: ")
		}
		fmt.Fprintln(w, pn.Number)
	}
}

func listPeople(w io.Writer, book *pb.AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}

// Main reads the entire address book from a file and prints all the
// information inside.
func main() {
	// if len(os.Args) != 2 {
	// 	log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	// }
	// fname := os.Args[1]

	// [START unmarshal_proto]
	// Read the existing address book.

	// #### data 파일 읽기
	// in, err := ioutil.ReadFile(fname)
	// if err != nil {
	// 	log.Fatalln("Error reading file:", err)
	// }

	// #### data 파일 생성
	// book := pb.AddressBook{People: []*pb.Person{
	// 	{
	// 		Name:  "John Doe",
	// 		Id:    101,
	// 		Email: "john@example.com",
	// 	},
	// 	{
	// 		Name: "Jane Doe",
	// 		Id:   102,
	// 	},
	// 	{
	// 		Name:  "Jack Doe",
	// 		Id:    201,
	// 		Email: "jack@example.com",
	// 		Phones: []*pb.Person_PhoneNumber{
	// 			{Number: "555-555-5555", Type: pb.Person_WORK},
	// 		},
	// 	},
	// 	{
	// 		Name:  "Jack Buck",
	// 		Id:    301,
	// 		Email: "buck@example.com",
	// 		Phones: []*pb.Person_PhoneNumber{
	// 			{Number: "555-555-0000", Type: pb.Person_HOME},
	// 			{Number: "555-555-0001", Type: pb.Person_MOBILE},
	// 			{Number: "555-555-0002", Type: pb.Person_WORK},
	// 		},
	// 	},
	// 	{
	// 		Name:  "Janet Doe",
	// 		Id:    1001,
	// 		Email: "janet@example.com",
	// 		Phones: []*pb.Person_PhoneNumber{
	// 			{Number: "555-777-0000"},
	// 			{Number: "555-777-0001", Type: pb.Person_HOME},
	// 		},
	// 	},
	// 	{
	// 		Name:  "John Doe",
	// 		Id:    101,
	// 		Email: "john@example.com",
	// 	},
	// }}

	// [END unmarshal_proto]

	// ### binary file 로 직렬화
	// out, err := proto.Marshal(&book)
	// if err != nil {
	// 	log.Fatalln("Failed to encode address book:", err)
	// }
	// if err := ioutil.WriteFile("test.data", out, 0644); err != nil {
	// 	log.Fatalln("Failed to write address book:", err)
	// }

	fname := "test.data"
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}
	book1 := &pb.AddressBook{}
	// [START_EXCLUDE]
	if err := proto.Unmarshal(in, book1); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	listPeople(os.Stdout, book1)
}
