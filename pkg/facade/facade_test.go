package facade

import (
	"fasad/pkg/comparison"
	"fasad/pkg/history"
	"fasad/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type person struct {
	name 			string
	balance 		int
	desiredCredit 	int
}

var (
	numOfPerson = 3
	persons = []person{
		{
			"John",
			100,
			2200,
		},
		{
			"Nik",
			100,
			500,
		},
		{
			"Alex",
			200,
			500,
		},

	}
	expects = []bool{false, true, true}
)

func TestCredit_TakeCredit(t *testing.T) {
	var status bool
	for i := 0; i < numOfPerson; i++ {
		verifier := comparison.NewVerifier(models.MaxCredit)
		if persons[i].name != models.UnreliablePerson1 &&
			persons[i].name != models.UnreliablePerson2 {
			status = true
		}
		reputation := history.NewCreditReputation(persons[i].name, status)
		credit := NewCredit(persons[i].name, persons[i].balance, persons[i].desiredCredit)
		result := credit.Take(verifier, reputation)
		assert.Equal(t, expects[i], result, persons[i].name)
		status = false
	}
}
