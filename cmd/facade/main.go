package main

import (
	"fasad/pkg/facade"
	"fasad/pkg/models"
	"fmt"
)

const (
	name = 			"Alex"
	myMoney = 		200
	desiredLoan = 	100
)

func main() {
	credit := facade.NewCredit(name, myMoney, desiredLoan)
	res := credit.TakeCredit()
	if res {
		fmt.Printf("%s\n%s %d$.", models.CreditApproved,
			models.CurrentBalance, myMoney + desiredLoan)
	} else {
		fmt.Printf("%s\n%s %d$.", models.CreditReject,
			models.CurrentBalance, myMoney)
	}
}
