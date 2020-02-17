package history

// Reputation checks whether a person has a good credit reputation.
type Reputation interface {
	CheckHistory(name string) bool
}

// reputation ...
type reputation struct {
	blackList 	[]string
}

// CheckHistory ...
func (c *reputation) CheckHistory(name string) bool {
	for i := 0; i < len(c.blackList); i++ {
		if name == c.blackList[i] {
			return false
		}
	}
	return true
}

// NewCreditReputation initializes the Reputation.
func NewCreditReputation(blackList []string) Reputation {
	return &reputation{
		blackList:blackList,
	}
}
