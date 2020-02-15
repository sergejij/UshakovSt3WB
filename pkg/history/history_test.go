package history

import (
	"fasad/pkg/models"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testCreditHistory = "test credit history comparison"
	personName = "Person name"
)

var (
	persons = []string{
		"John",
		"Brad",
		"Alex",
		"Fill",
		"Nik",
		"Jack",
	}

	want = []bool {
		false,
		true,
		true,
		false,
		true,
		true,
	}
)

func TestCreditHistory_HistoryCheck(t *testing.T) {
	t.Run(testCreditHistory, func(t *testing.T) {
		var status bool
		for i, name := range persons {
			if name != models.UnreliablePerson1 && name != models.UnreliablePerson2 {
				status = true
			}
			story := NewCreditReputation(name, status)
			result := story.CheckHistory(name)
			assert.Equal(t, want[i], result, fmt.Sprintf("%s - %s.",
				personName, persons[i]))
			status= false
		}
	})
}
