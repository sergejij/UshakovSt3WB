package main

import (
	"fasad/pkg/facade"
	"fmt"
)

func main() {
	name := "Fill"
	money, desiredLoan, maxCredit := 2100, 125, 200
	credit := facade.NewCreditTaker(name, money, desiredLoan, maxCredit)
	res := credit.CreditTake()
	fmt.Println(res)
}
