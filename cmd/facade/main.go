package main

import (
	"fasad/pkg/comparison"
	"fasad/pkg/facade"
	"fasad/pkg/history"
	"fasad/pkg/models"
	"fmt"
)

const (
	name = 			"John"
	myMoney = 		1000
	desiredLoan = 	500
)

func main() {
	var status bool
	verifier := comparison.NewVerifier(models.MaxCredit)
	if name != models.UnreliablePerson1 && name != models.UnreliablePerson2 {
		status = true
	}
	creditReputation := history.NewCreditReputation(name, status)
	credit := facade.NewCredit(name, myMoney, desiredLoan)
	result := credit.Take(verifier, creditReputation)
	if result {
		fmt.Printf("%s\n%s %d$.", models.CreditApproved,
			models.CurrentBalance, myMoney + desiredLoan)
	} else {
		fmt.Printf("%s\n%s %d$.", models.CreditReject,
			models.CurrentBalance, myMoney)
	}
}
