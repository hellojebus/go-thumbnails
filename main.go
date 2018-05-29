//usage:
// $ go run main.go
//   myfile.jpg
//   ^D
package main

import (
	"bufio"
	"os"
	"github.com/hellojebus/go-thumbnails/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.F
	}
}