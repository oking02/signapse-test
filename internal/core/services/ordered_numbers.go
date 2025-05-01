package services

import (
	"context"
	"fmt"
	"slices"
	"sort"

	"github.com/oking02/signapse-test/internal/core/ports"
)

var _ ports.OrderedNumbersFetcher = (*OrderedNumbers)(nil)

type OrderedNumbers struct {
	fetcher ports.RawNumbersFetcher
}

func NewOrderedNumbers(fetcher ports.RawNumbersFetcher) *OrderedNumbers {
	return &OrderedNumbers{
		fetcher: fetcher,
	}
}

// FetchOrderedNumbers retrieves raw numbers, filter out non ints,
// processes them into an ordered format, with odd before even and returns the resulting sorted list.
// e.g. [3,5,7,4,6,8]
// note:
//
//	The striping of non-integer values could be done in the fetcher,
//	but as that is essentially business logic, it's best kept here
func (o OrderedNumbers) FetchOrderedNumbers(ctx context.Context) ([]int, error) {

	rawNumbers, err := o.fetcher.FetchRawNumbers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch raw numbers: %w", err)
	}

	return o.createOrderedNumbers(rawNumbers), nil
}

func (o OrderedNumbers) createOrderedNumbers(input []any) []int {

	if len(input) == 0 {
		return []int{}
	}

	intNumbers := toIntSlice(input)
	even, odd := splitByEvenOdd(intNumbers)

	sort.Ints(even)
	sort.Ints(odd)

	return slices.Concat(odd, even)
}

// toIntSlice converts a []any to []int
// discarding any values that are not int's
func toIntSlice(input []any) []int {
	var results []int

	for _, a := range input {

		// not exhaustive, in reality for this task it will only deal
		// with float64 as by default JSON numbers are parsed as float64's
		switch a.(type) {
		case int:
			results = append(results, a.(int))
		case float32:
			v := a.(float32)
			results = append(results, int(v))
		case float64:
			v := a.(float64)
			results = append(results, int(v))
		}

	}

	return results
}

// splitByEvenOdd splits a single []int into an even
// and odd []int
func splitByEvenOdd(input []int) (even []int, odd []int) {

	for _, num := range input {
		if x := num % 2; x == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}

	return even, odd
}
