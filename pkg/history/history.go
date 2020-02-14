package history

type creditReputation struct {
	person map[string]bool
}

type CreditReputation interface {
	CheckHistory(personName string) bool
}

func (c *creditReputation) CheckHistory(personName string) bool {
	return c.person[personName]
}

// NewHistoryChecker initializes the HistoryChecker.
func NewCreditReputation(person string, status bool) CreditReputation {
	return &creditReputation{
		person: map[string]bool{
			person:status,
		},
	}
}
