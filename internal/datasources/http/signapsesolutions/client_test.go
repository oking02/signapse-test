package signapsesolutions

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/oking02/signapse-test/internal/core/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestRoundTripFunc func(req *http.Request) *http.Response

func (f TestRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestClient_FetchRawNumbers(t *testing.T) {

	var (
		baseURL = "https://technical-test.api.production.signapsesolutions.com"
		token   = "some-token"
	)

	tests := []struct {
		name       string
		tran       TestRoundTripFunc
		wantResult []any
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			tran: TestRoundTripFunc(func(req *http.Request) *http.Response {
				assert.Equal(t, fmt.Sprintf("Bearer %s", token), req.Header.Get("Authorization"))

				b, err := os.ReadFile("testdata/number1.json")
				require.NoError(t, err)

				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
				}
			}),
			wantResult: []any{
				float64(81),
				float64(10),
				"fe095273-2127-4493-b795-d297e52df06b",
				float64(18),
				float64(28),
				float64(76),
				float64(25),
				float64(45),
				float64(95),
				float64(41),
				"83b31517-a866-4f2b-afc9-36f8653088c9",
				"469693ba-5d47-450e-a82a-ea7e994c17a7",
				float64(12)},
			wantErr: assert.NoError,
		},
		{
			name: "500 response code",
			tran: TestRoundTripFunc(func(req *http.Request) *http.Response {
				assert.Equal(t, fmt.Sprintf("Bearer %s", token), req.Header.Get("Authorization"))

				return &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       nil,
				}
			}),
			wantResult: nil,
			wantErr:    assert.Error,
		},
		{
			name: "401 response code",
			tran: TestRoundTripFunc(func(req *http.Request) *http.Response {
				assert.Equal(t, fmt.Sprintf("Bearer %s", token), req.Header.Get("Authorization"))

				return &http.Response{
					StatusCode: http.StatusUnauthorized,
					Body:       nil,
				}
			}),
			wantResult: nil,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return errors.Is(err, domain.ErrAuthFailed)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cl := NewClient(
				&http.Client{
					Transport: tt.tran,
				},
				baseURL,
				token)

			res, err := cl.FetchRawNumbers(t.Context())
			tt.wantErr(t, err)
			assert.Equal(t, tt.wantResult, res)
		})
	}

}
