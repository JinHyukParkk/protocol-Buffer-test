package main

import (
	"github.com/JinHyukParkk/protocol_Buffer_test/action"
)

// Main reads the entire address book from a file and prints all the
// information inside.
func main() {
	// if len(os.Args) != 2 {
	// 	log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	// }
	// fname := os.Args[1]

	// #### data 파일 읽기
	// in, err := ioutil.ReadFile(fname)
	// if err != nil {
	// 	log.Fatalln("Error reading file:", err)
	// }
	// action.Write()
	action.Read()

}
