package main

import (
	"fmt"
	"repo.abs.com/hiring/allocate/src/initiate"
)

func main() {
	err:=initiate.Initialise()
	if err != nil {
		fmt.Println(err)
	}
}
