package main

import (
	"bufio"
	"fmt"
	"os"
)

func runRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println("echoing", text)
}
