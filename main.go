//usage:
// $ go run main.go
//   myfile.jpg
//   ^D
package main

import (
	"bufio"
	"os"
	"github.com/hellojebus/go-thumbnails/thumbnail"
	"log"
	"fmt"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}