package main

import "fmt"
import "bufio"
import "os"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			fmt.Printf(string(line))
		}
		if err != nil {
			break
		}
	}
}
