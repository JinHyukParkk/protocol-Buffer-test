package action

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/JinHyukParkk/protocol_Buffer_test/tutorial"
	"github.com/gogo/protobuf/proto"
)

func listPeople(w io.Writer, book *pb.AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}

func Read() {
	fname := "data/test.data"
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
