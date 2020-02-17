package facade

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

/*type person struct {
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
*/

/*func TestCredit_TakeCredit(t *testing.T) {
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
}*/
type verifierMock struct {
	mock.Mock
}

type reputationMock struct {
	mock.Mock
}

/*type VerifierMock interface {
	Verify(desiredCredit int) (bool, error)
}*/

func (v *verifierMock) Verify(desiredCredit int) (bool, error) {
	args := v.Called(desiredCredit)
	return args.Bool(0), args.Error(1)
}
func (c *reputationMock) CheckHistory(name string) (bool, error) {
	args := c.Called(name)
	return args.Bool(0), args.Error(1)
}
/*func NewVerifierMock(maxCredit int) VerifierMock {
	return &verifierMock {}
}*/


func TestCredit_Take(t *testing.T) {
	money := 1000
	desiredCredit := 500
	name := "Alex"
	testVerifier := new(verifierMock)
	testVerifier.On("Verify", desiredCredit).Return(true, nil)
	//testVerifier.Verify(desiredCredit)
	//testVerifier.AssertExpectations(t)


	testReputation := new(reputationMock)
	testReputation.On("CheckHistory", name).Return(true, nil)

	credit := NewCredit(name, money, 1500)
	credit.Take(testVerifier, testReputation)
}
