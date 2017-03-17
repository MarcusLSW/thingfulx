package thingfulx

import (
	"time"

	"golang.org/x/net/context"
)

// FetcherBuilder is type definition for the constructor functions that new
// Fetchers must implement.
type FetcherBuilder func() (Fetcher, error)

// Fetcher is the main interface for things that know how to fetch some
// resource from somewhere handling any protocol weirdness or authentication
// requirements and return back a slice of parsed Thing instances. The
// interface is agnostic in terms of protocol so an implementing class is free
// to perform whatever steps required to get the data.
type Fetcher interface {
	Seeder

	Crawler

	Parser
}

// Seeder in the interface used by Pomelo to generate a minimum set of URLS that
// should be indexed for a host.
type Seeder interface {
	// URLS returns the smallest set of URLs required to completely index the upstream
	// data provider. The returned slice might be a single URL for hosts that publish
	// relatively few things, or it might be a whole bunch for hosts that have pages
	// of results.
	// This function takes as parameters an instance of the Context Interface for
	// request-scoped values, a Client implementation which the Fetcher must
	// use to make any outgoing requests, plus a delay Duration. Requests should
	// not be made faster than the time interval specified by delay.
	URLS(ctx context.Context, client Client, delay time.Duration) ([]string, error)
}

// Crawler is the interface that must be implemented for performing
// the task of going and getting data from some remote data provider, and
// returning a slice of Thing objects extracted from that data source.
// Certain providers might require multiple requests from different URLs to
// actually retrieve the data so within this interface they are free to do
// this.
type Crawler interface {
	// Crawl takes as parameters an instance of the Context Interface for
	// request-scoped values, the url we want to get data from, a Client object
	// that the crawler will use to actually make the request.
	// It returns the data provider response as a slice of bytes.
	Crawl(ctx context.Context, url string, client Client) ([]byte, error)
}

// Parser is the interface used to provide a full representation of a Thing.
// This representation must include both data and metadata required to describes
// a resource from both a factual and a semantic point of view.
// For a comprehensive definition of how to semantically define a Thing and what
// are the ontologies used by Thingful refer to https://github.com/thingful/schema.
type Parser interface {
	// Parse returns a slice of Thing objects extracted from that data source.
	// This function takes as parameters a slice of bytes representing the data
	// collected from the upstream data provider and a TimeProvider instance used
	// internally by the fetcher to record the indexing time of the parser.
	// This is to allow for easier testing.
	Parse(rawData []byte, timeProvider TimeProvider) ([]Thing, error)
}
