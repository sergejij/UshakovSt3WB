package creditHistory

type HistoryChecker interface {
	HistoryCheck(personName string) string
}
type creditHistory struct {

}
func (c *creditHistory) HistoryCheck(personName string) string {
	if personName == "Alex" || personName == "Fill" {
		return "bad"
	}
	return "good"
}

// NewHistoryChecker initializes the HistoryChecker.
func NewHistoryChecker() (HistoryChecker, error) {
	return &creditHistory{}, nil
}

