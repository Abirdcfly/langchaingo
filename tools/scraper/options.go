package scraper

type Options func(*Scraper)

// WithMaxDepth sets the maximum depth for the Scraper.
//
// Default value: 1
//
// maxDepth: the maximum depth to set.
// Returns: an Options function.
func WithMaxDepth(maxDepth int) Options {
	return func(o *Scraper) {
		o.MaxDepth = maxDepth
	}
}

// WithParallelsNum sets the number of maximum allowed concurrent
// requests of the matching domains
//
// Default value: 2
//
// parallels: the number of parallels to set.
// Returns: the updated Scraper options.
func WithParallelsNum(parallels int) Options {
	return func(o *Scraper) {
		o.Parallels = parallels
	}
}

// WithDelay creates an Options function that sets the delay of a Scraper.
//
// The delay parameter specifies the amount of time in milliseconds that
// the Scraper should wait between requests.
//
// Default value: 3
//
// delay: the delay to set.
// Returns: an Options function.
func WithDelay(delay int64) Options {
	return func(o *Scraper) {
		o.Delay = delay
	}
}

// WithAsync sets the async option for the Scraper.
//
// Default value: true

// async: The boolean value indicating if the scraper should run asynchronously.
// Returns a function that sets the async option for the Scraper.
func WithAsync(async bool) Options {
	return func(o *Scraper) {
		o.Async = async
	}
}

// WithNewBlacklist creates an Options function that replaces
// the list of url endpoints to be excluded from the scraping,
// with a new list.
//
// Default value:
//
//	[]string{
//		"login",
//		"signup",
//		"signin",
//		"register",
//		"logout",
//		"download",
//		"redirect",
//	},
//
// blacklist: slice of strings with url endpoints to be excluded from the scraping.
// Returns: an Options function.
func WithNewBlacklist(blacklist []string) Options {
	return func(o *Scraper) {
		o.Blacklist = blacklist
	}
}

// WithBlacklist creates an Options function that appends
// the url endpoints to be excluded from the scraping,
// to the current list
//
// Default value:
//
//	[]string{
//		"login",
//		"signup",
//		"signin",
//		"register",
//		"logout",
//		"download",
//		"redirect",
//	},
//
// blacklist: slice of strings with url endpoints to be excluded from the scraping.
// Returns: an Options function.
func WithBlacklist(blacklist []string) Options {
	return func(o *Scraper) {
		o.Blacklist = append(o.Blacklist, blacklist...)
	}
}

// WithHandleLinks sets the handleLinks option for the Scraper.
//
// Default value: false

// handleLinks: The boolean value indicating if the scraper will handle links of href tag
// Returns a function that sets the handleLinks option for the Scraper.
func WithHandleLinks(handleLinks bool) Options {
	return func(o *Scraper) {
		o.HandleLinks = handleLinks
	}
}

// WithMaxScrapedDataLength sets the max scraped data length option for the Scraper.
//
// Default value: 15000

// WithMaxScrapedDataLength: The int value indicating the max scraped data length
// Returns a function that sets the max scraped data length for the Scraper.
func WithMaxScrapedDataLength(maxScrapedDataLength int) Options {
	return func(o *Scraper) {
		o.MaxScrapedDataLength = maxScrapedDataLength
	}
}
