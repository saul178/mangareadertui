package mangaupdates_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/saul178/mangareadertui/internal/mangaupdates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
TODO: this is just to test that these functions return what i expect
1.) create a mock test for the api on actual errors where the api expects certain requirements
2.) normal 404 where no errors is reported EX: series not found is 404 but thats not an error so it returns just 404 and no response body
3.) mock test for a successful api response
**/

// TODO: implement test for search function on successful search, 400 bad request search, 200 ok but empty search
func TestRequestAndDecode(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name          string
		statusCode    int
		responseBody  string
		expected      mangaupdates.Response
		expectedError bool
	}{
		{
			name:       "400 validation error",
			statusCode: http.StatusBadRequest,
			responseBody: `{
				"status": "exception",
				"reason": "Field Validation Error",
				"context": {
					"search": [
						{
							"index": 0,
							"errors": ["\"\" must have a length between 1 and 400"]
						}
					],
					"exclude": [
						{
							"index": 0,
							"errors": ["\"\" must have a length between 1 and 100"]
						}
					]
				}
			}`,
			expected: mangaupdates.Response{
				Status: "exception",
				Reason: "Field Validation Error",
				Context: map[string][]mangaupdates.ContextErrors{
					"search": {
						{Index: 0, Errors: []string{"\"\" must have a length between 1 and 400"}},
					},
					"exclude": {
						{Index: 0, Errors: []string{"\"\" must have a length between 1 and 100"}},
					},
				},
			},
		},
		// {
		// 	name:         "404 response",
		// 	statusCode:   http.StatusNotFound,
		// 	responseBody: `{}`,
		// 	expected:     mangaupdates.Response{},
		// },
		// {
		// 	name:       "200 success",
		// 	statusCode: http.StatusOK,
		// 	responseBody: `{
		// 		"status": "success",
		// 		"reason": "login successful"
		// 	}`,
		// 	expected: mangaupdates.Response{
		// 		Status: "success",
		// 		Reason: "login Succesful",
		// 	},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.statusCode)
				_, err := w.Write([]byte(tc.responseBody))
				require.NoError(t, err)
			}))
			defer testServer.Close()

			// muClient := mangaupdates.NewMangaUpdatesClient()


			var got mangaupdates.Result
			// err := "TODO: implement search function then test"
			// require.NoError(t, err)
			assert.Equal(t, tc.expected, got)
			assert.Equal(t, tc.statusCode, got.StatusCode)
		})
	}

}

/*
TODO:
create a test client and send requests to endpoints and see if we get a response back either by httpstatuscode or getting a struct of info back
*/
func TestDoRequest(t *testing.T) {}
