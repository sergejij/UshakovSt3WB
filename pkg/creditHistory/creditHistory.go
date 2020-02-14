package creditHistory

import (
	"fasad/pkg/models"
)

type CreditHistory interface {
	HistoryCheck(personName string) bool
}

type creditHistory struct {
	person map[string]bool
}

func (c *creditHistory) HistoryCheck(personName string) bool {
	if c.person[personName] {
		return true
	}
	return false
}

// NewHistoryChecker initializes the HistoryChecker.
func NewCreditHistory(name string) CreditHistory {
	if name == models.UnreliablePerson1 || name == models.UnreliablePerson2 {
		return &creditHistory{
			person: map[string]bool{
				name:false,
			},
		}
	}
	return &creditHistory{
		person: map[string]bool{
			name:true,
		},
	}
}

