package services

import (
	"errors"
	"testing"

	"github.com/oking02/signapse-test/internal/core/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderedNumbers_FetchOrderedNumbers(t *testing.T) {

	testErr := errors.New("test error")

	tests := []struct {
		name    string
		fetcher *mocks.RawNumbersFetcher
		results []int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fetcher: func() *mocks.RawNumbersFetcher {
				m := mocks.NewRawNumbersFetcher(t)
				m.EXPECT().FetchRawNumbers(mock.Anything).
					Once().
					Return([]any{
						84, 92, 42,
						"84e35e86-7a26-492c-8c68-aedac37b3000",
						65,
						"3f4aaca0-50bc-49ff-92df-a92807154124",
						7, 74,
						"c3a7c0cc-7c65-4cd0-919d-bb79f22130f2",
						39, 66, 59, 9}, nil)
				return m
			}(),
			results: []int{
				7, 9, 39, 59, 65, 42, 66, 74, 84, 92,
			},
			wantErr: assert.NoError,
		},
		{
			name: "empty input",
			fetcher: func() *mocks.RawNumbersFetcher {
				m := mocks.NewRawNumbersFetcher(t)
				m.EXPECT().FetchRawNumbers(mock.Anything).
					Once().
					Return([]any{}, nil)
				return m
			}(),
			results: []int{},
			wantErr: assert.NoError,
		},
		{
			name: "error fetching raw numbers",
			fetcher: func() *mocks.RawNumbersFetcher {
				m := mocks.NewRawNumbersFetcher(t)
				m.EXPECT().FetchRawNumbers(mock.Anything).
					Once().
					Return(nil, testErr)
				return m
			}(),
			results: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return errors.Is(err, testErr)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewOrderedNumbers(tt.fetcher)
			result, err := svc.FetchOrderedNumbers(t.Context())
			tt.wantErr(t, err)
			assert.ElementsMatch(t, tt.results, result)
		})
	}
}
