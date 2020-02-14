package facade

import (
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
	for i := 0; i < numOfPerson; i++ {
		facade := NewCredit(persons[i].name, persons[i].balance, persons[i].desiredCredit)
		result := facade.TakeCredit()
		assert.Equal(t, expects[i], result, persons[i].name)
	}
}
