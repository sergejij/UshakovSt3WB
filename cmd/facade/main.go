package main

import (
	"fasad/pkg/blacklist"
	"fasad/pkg/comparison"
	"fasad/pkg/facade"
	"fasad/pkg/history"
	"fasad/pkg/models"
	"fmt"
)

const (
	name = 			"Nik"
	myMoney = 		1000
	desiredLoan = 	1200
)

var (
	unreliablePersons = []string{"John", "Nik", "Fill"}
)

func main() {
	verifier := comparison.NewVerifier(models.MaxCredit)

	blockedPersons := blacklist.NewBlocked(unreliablePersons)
	status := blockedPersons.Check(name) // можно ли метод вызвать в майне?

	reputation := history.NewCreditReputation(name, status)
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
