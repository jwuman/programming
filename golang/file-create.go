package main

import (
	"fmt"
	"io/utils"
	"log"
	"os"
)

func main() {
	file, err := os.Create("/tmp/myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.writeString("This is some text")
	file.Close()
	stream, err := ioutil.ReadFile("/tmp/myfile.txt")

}
