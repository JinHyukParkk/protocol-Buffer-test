package action

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"

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

func Write() {
	// #### data 파일 생성
	book := pb.AddressBook{People: []*pb.Person{
		{
			Name:  "John Doe",
			Id:    101,
			Email: "john@example.com",
		},
		{
			Name: "Jane Doe",
			Id:   102,
		},
		{
			Name:  "Jack Doe",
			Id:    201,
			Email: "jack@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-555-5555", Type: pb.Person_WORK},
			},
		},
		{
			Name:  "Jack Buck",
			Id:    301,
			Email: "buck@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-555-0000", Type: pb.Person_HOME},
				{Number: "555-555-0001", Type: pb.Person_MOBILE},
				{Number: "555-555-0002", Type: pb.Person_WORK},
			},
		},
		{
			Name:  "Janet Doe",
			Id:    1001,
			Email: "janet@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-777-0000"},
				{Number: "555-777-0001", Type: pb.Person_HOME},
			},
		},
		{
			Name:  "John Doe",
			Id:    101,
			Email: "john@example.com",
		},
	}}
	// ### binary file 로 직렬화
	out, err := proto.Marshal(&book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("../data/test.data", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}
