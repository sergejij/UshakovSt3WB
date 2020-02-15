package history

type Reputation interface {
	CheckHistory(personName string) bool
}

type reputation struct {
	person map[string]bool
}

func (c *reputation) CheckHistory(personName string) bool {
	return c.person[personName]
}

// NewCreditReputation initializes the Reputation.
func NewCreditReputation(person string, status bool) Reputation {
	return &reputation{
		person: map[string]bool{
			person:status,
		},
	}
}
