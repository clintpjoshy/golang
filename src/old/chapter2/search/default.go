package search

// defaultMatcher implements the default matcher. If a matcher is not found we use default matcher in search.go
type defaultMatcher struct {}

// default matcher is registered
func init() {
  var matcher defaultMatcher
	Register("default", matcher);
}

// Search Behavior for default matcher
// Default behavior is to return nill for both error and feed.
// dafaultMatcher is the receiver type here.

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
  return nil, nil
}
