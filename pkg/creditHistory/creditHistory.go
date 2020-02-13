package creditHistory

type HistoryChecker interface {
	HistoryCheck(personName string) bool
}

type creditHistory struct {
	person map[string]bool
}

func (c *creditHistory) HistoryCheck(personName string) bool {
	if c.person[personName] {
		return false
	}
	return true
}

// NewHistoryChecker initializes the HistoryChecker.
func NewHistoryChecker(name string) HistoryChecker {
	return &creditHistory{
		person: map[string]bool{name:true},
	}
}

