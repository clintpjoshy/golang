package search

import (
  "log"
	"sync" // sync needed for synchronizing goroutines
)

// creates a map with matchers for searching
// package level variable and therefore is accessble throughout the package.
var matchers = make(map[string]Matcher)

//search logic
//exportable	
// this funciton has the business logic. 
// This is a good example, that gives an overview of how GO programs must be structured

func Run(searchTerm string) {

	//retrieve all the feeds to match with specifed search term
	// A function can have multiple return values.
	// Short variable declaratin operator (:=) can be used in  a function. declare and initialize at a time.
	// Compiler detemines the type of each variable depending on the return form the funciton call.
  feeds, err := RetrieveFeeds()
	if err != nil {
		// if error then the program is terminated using Fatal. Logs to termainal window.
		log.Fatal(err)
	}

	//unbuffered channel to receive results
	// make can be used to create a channel as well.
	// channel is needed to communicate (share data) with goroutes.
	results := make(chan *Result)

	// Wait group to process all the feeds
	// Waitgroups are needed to not to terminate the program before all the search is completed.
	// This is a counting semaphore that keeps track of when a routine is finished running.
	var waitGroup sync.WaitGroup

	// Number of Go routines needed to process each feed
	// This will be the same as the number of GR's we are going to launch (see below in the for loop)
	waitGroup.Add(len(feeds))

	// GR for each feed to find results
	//range can be used with slices, strings, arrays, channels, maps
	// blank (_) is identifier can be used here as well. First item indicates the index and second indicates the value.
	for _, feed := range feeds{
	  //find matchers for each search
		matcher, exists := matchers[feed.Type]

		if !exists {
			// if no matcher found assign "default"
			// this ensures the program to run without any interruptions or hiccups.
		  matcher = matchers["default"]
		}

    // Launch go routine to search the search term
		// Each feed will be concurrently processed without depending on each other.
		// this anonymous function takes 2 parameters namely matcher with type Mathcer and feed which is a pointer value
		// This anonymous func is invoked immediately with parameters.
		// Each GR calls a function called Match (found in match.go).
		// After function call we decrease the counting semaphore by 1.
		// Closures can be used in GOLang as well. In this case, searchTerm, results are not passed into the anonymous function. Variables from outside of function is accessbile in a inner function. 
		// Outer funciton is allready executed from main.go. But the variables are available for inner functions.
		// for matcher and feed if we use closure, it'll be the same value from the outer scope (mostly the last value after the for loop finishes running). In order to avoid this, we can pass in the value as parameters and parameters are by value and therefore will be independent of outer variables..
		// results are channels that are used to share data b/w GRs.
  	go func(matcher Matcher, feed *Feed) {
   		Match(matcher, feed, searchTerm, results)
  		// Decrements the counter by 1.
  		waitGroup.Done()
  	}(matcher, feed)
  }

  // GR for monitoring after all the work is done
	// This is needed to make sure that the main func is not shut down. When main func is shut down, the program terminates.
	// WaitGroup and results are available in this GR via closures.
	// Wait method makes sure that the main func is 
  go func() {
  	// Wait to complete all the processing
		// Wait until waitGroup counting semaphore becomes 0
    waitGroup.Wait()

  	//Indicate when to exit the program
		// Once waitGroup becomes 0, close is called on the channel which will terminate the program.
  	close(results)
  }()

	// Found in match.go
  Display(results)
}

// Matcher value will be added to the map of registered matchers.
// These registrations must happen before main and there init is a great place to do this registration.
func Register(feedType string, matcher Matcher) {
  if _, exists := matchers[feedType]; exists {
	  log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
