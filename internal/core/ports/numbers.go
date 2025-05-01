package ports

import "context"

// RawNumbersFetcher is the interface for fetching the raw unordered number list
type RawNumbersFetcher interface {
	FetchRawNumbers(ctx context.Context) ([]any, error)
}

// OrderedNumbersFetcher is the interface for fetching a random odd even ordered number list
type OrderedNumbersFetcher interface {
	FetchOrderedNumbers(ctx context.Context) ([]int, error)
}
