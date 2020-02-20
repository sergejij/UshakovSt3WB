package main

import (
	"fmt"

	"task1/facade/pkg/facade"
	"task1/facade/pkg/history"
	"task1/facade/pkg/models"
	"task1/facade/pkg/validator"
)

const (
	name        = "Nik"
	myMoney     = 400
	desiredLoan = 500
	MaxCredit   = 1500
)

var (
	unreliablePersons = []string{"John", "Nik", "Fill"}
)

func main() {
	nameValidator := validator.NewNameValidator(MaxCredit)
	reputation := history.NewReputation(unreliablePersons)
	credit := facade.NewCredit(nameValidator, reputation)
	result := credit.Take(name, myMoney, desiredLoan)
	if result {
		fmt.Printf("%s\n%s %d$.", models.CreditApproved,
			models.CurrentBalance, myMoney+desiredLoan)
	} else {
		fmt.Printf("%s\n%s %d$.", models.CreditReject,
			models.CurrentBalance, myMoney)
	}
}
