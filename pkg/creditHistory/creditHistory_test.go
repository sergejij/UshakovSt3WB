package creditHistory

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testCreditHistory = "test credit history check"
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
		for i, name := range persons {
			story := NewCreditHistory(name)
			result := story.HistoryCheck(name)
			assert.Equal(t, want[i], result, fmt.Sprintf("%s - %s.",
				personName, persons[i]))
		}
	})
}
