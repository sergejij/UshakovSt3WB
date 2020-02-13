package main

import (
	"fasad/pkg/facade"
	"fmt"
)

func main() {
	credit, err := facade.NewCreditTaker("Jonh", 2100, 112)
	if err != nil {
		fmt.Errorf("Error: the facade is not working.\n")
	}
	res := credit.CreditTake()
	fmt.Println(res)
}

