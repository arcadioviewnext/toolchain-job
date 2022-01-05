package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hi, TEST 2, from a batch job! My index is: %s\n", os.Getenv("JOB_INDEX"))
}
