package main

import (
	"fasad/pkg/comparison"
	"fasad/pkg/facade"
	"fasad/pkg/history"
	"fasad/pkg/models"
	"fmt"
)

const (
	name = 			"Nik"
	myMoney = 		1200
	desiredLoan = 	500
)

var (
	unreliablePersons = []string{"John", "Nik", "Fill"}
)

func main() {
	verifier := comparison.NewVerifier(models.MaxCredit)
	reputation := history.NewCreditReputation(unreliablePersons)
	credit := facade.NewCredit(name, myMoney, desiredLoan)
	result := credit.Take(verifier, reputation)
	if result {
		fmt.Printf("%s\n%s %d$.", models.CreditApproved,
			models.CurrentBalance, myMoney + desiredLoan)
	} else {
		fmt.Printf("%s\n%s %d$.", models.CreditReject,
			models.CurrentBalance, myMoney)
	}
}
