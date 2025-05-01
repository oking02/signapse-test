package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oking02/signapse-test/internal/core/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetOrderedNumbers_ServeHTTPt(t *testing.T) {

	tests := []struct {
		name       string
		fetcher    *mocks.OrderedNumbersFetcher
		req        *http.Request
		assertResp func(t *testing.T, resp *httptest.ResponseRecorder)
	}{
		{
			name: "success",
			fetcher: func() *mocks.OrderedNumbersFetcher {
				m := mocks.NewOrderedNumbersFetcher(t)
				m.EXPECT().FetchOrderedNumbers(mock.Anything).
					Once().
					Return([]int{7, 9, 39, 59, 65, 42, 66, 74, 84, 92}, nil)
				return m
			}(),
			req: func() *http.Request {

				// if required, add headers, path variables, etc.

				return httptest.NewRequestWithContext(
					t.Context(), http.MethodGet, "/", nil)
			}(),
			assertResp: func(t *testing.T, resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, resp.Code)

				var result struct {
					Data []int `json:"data"`
				}
				require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &result))
				assert.ElementsMatch(t,
					[]int{7, 9, 39, 59, 65, 42, 66, 74, 84, 92},
					result.Data)
			},
		},
		{
			name: "success empty response",
			fetcher: func() *mocks.OrderedNumbersFetcher {
				m := mocks.NewOrderedNumbersFetcher(t)
				m.EXPECT().FetchOrderedNumbers(mock.Anything).
					Once().
					Return([]int{}, nil)
				return m
			}(),
			req: func() *http.Request {

				// if required, add headers, path variables, etc.

				return httptest.NewRequestWithContext(
					t.Context(), http.MethodGet, "/", nil)
			}(),
			assertResp: func(t *testing.T, resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, resp.Code)

				var result struct {
					Data []int `json:"data"`
				}
				require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &result))
				assert.ElementsMatch(t,
					[]int{},
					result.Data)
			},
		},
		{
			name: "fetch error",
			fetcher: func() *mocks.OrderedNumbersFetcher {
				m := mocks.NewOrderedNumbersFetcher(t)
				m.EXPECT().FetchOrderedNumbers(mock.Anything).
					Once().
					Return(nil, fmt.Errorf("test error"))
				return m
			}(),
			req: func() *http.Request {

				// if required, add headers, path variables, etc.

				return httptest.NewRequestWithContext(
					t.Context(), http.MethodGet, "/", nil)
			}(),
			assertResp: func(t *testing.T, resp *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, resp.Code)

				var result struct {
					Message string `json:"message"`
					Errors  any    `json:"errors"`
				}
				require.NoError(t, json.Unmarshal(resp.Body.Bytes(), &result))
				assert.Equal(t, "failed to get ordered numbers", result.Message)
				assert.Equal(t, "test error", result.Errors)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := NewGetOrderedNumbers(tt.fetcher)
			resp := httptest.NewRecorder()

			ctrl.ServeHTTP(resp, tt.req)
			tt.assertResp(t, resp)
		})
	}
}
