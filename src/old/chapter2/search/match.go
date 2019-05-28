package search

import (
	"fmt"
	"log"
)

// Result type
type Result struct {
	Field   string
	Content string
}

// Interface type.
// Interface declare behaviors.
// Behaviors are defined by the methods declared in the interface type.
// This is required to satisfy the behavior to be implemented by struct type.
// Search is a method.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a GR for each feed.
// Search concurrently.
// Search against specified matcher
// Actual search is done here with either value or pointers

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// assign results in channel
	for _, result := range searchResults {
		results <- result
	}
}

// Writes to the terminal as results are found.
//
func Display(results chan *Result) {
	// for loop terminates once the channel closes.
	// because of the use of channels and range keywords are used all the results are processed before the channel being closed.
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
