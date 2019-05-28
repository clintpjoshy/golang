//<rss xmlns:npr="http://www.npr.org/rss/" xmlns:nprml="http://api
//    <channel>
//        <title>News</title>
//        <link>...</link>
//        <description>...</description>
//log.Fatalln(feedType, "Matcher already registered")
// Licensed to Mark Watson <nordickan@gmail.com>
//RSS matcher 33
//        <language>en</language>
//        <copyright>Copyright 2014 NPR - For Personal Use
//        <image>...</image>
//        <item>
//            <title>
//                Putin Says He'll Respect Ukraine Vote But U.S.
//            </title>
//            <description>
//                The White House and State Department have called on the
//            </description>

// The above format is how rss document xml format.
// this matcher is the same as the default matcher in search.
// Only change is the implmentation of interface method Search
// Decoding xml is very similar to decoding json found in feed.go
package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"chapter2/search"
)

// struct types to decode the document.
// multiple types can be defined as follows by using just one type keyword.
type (
	// items in the item tag in xml
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubData     string   `xml: "pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml: "link"`
		GUID        string   `xml: "guid"`
		GeoRssPoint string   `xml: "georss:point"`
	}

	//items in the image tag
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// items in the channel tag
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          string   `xml:"image"`
		Item           string   `xml:"item"`
	}

	// Fields in the rss document based on the xml above
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// This is for matcher interface
// This is just like default matcher. It can be an empty struct because we do not need to maintatin any states.
type rssMatcher struct{}

// Init func registers a value for rssMatcher just like for default matcher.
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Make http call and decode
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed URI provided")
	}

	// This returns either err or resp which is a pointer to a value of type response
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// response needs to be closed after the function returns
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// Decode xml.
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

// search for a specified search term
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s]\n", feed.Type, feed.Name, feed.URI)

	//data for search
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		//title of the search term
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)

		if err != nil {
			return nil, err
		}

		// save if a match is found
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// check description for search item
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// If a match is found save it
		// & is used to get the address of the new value
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}
