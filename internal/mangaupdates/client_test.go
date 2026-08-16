package mangaupdates_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/saul178/mangareadertui/internal/mangaupdates"
	"github.com/stretchr/testify/assert"
)

/*
*
TODO: this is just to test that these functions return what i expect
1.) create a mock test for the api on actual errors where the api expects certain requirements
2.) normal 404 where no errors is reported EX: series not found is 404 but thats not an error so it returns just 404 and no response body
3.) mock test for a successful api response
*
*/
func TestTypeAssertionCase(t *testing.T) {
	sr := &mangaupdates.SearchResults{}
	r := &mangaupdates.Result{}
	s := &mangaupdates.Series{}
	f := &mangaupdates.SeriesSearchDataList{}
	srStr := fmt.Sprintf("SearchResult type matched: %T", sr)
	rStr := fmt.Sprintf("Result type matched: %T", r)
	sStr := fmt.Sprintf("Series type matched: %T", s)

	testCases := []struct {
		name      string
		modelType any
		expected  string
		err       error
	}{
		{
			name:      "Want ptr to SearchResults struct passed through",
			modelType: sr,
			expected:  srStr,
			err:       nil,
		},
		{
			name:      "Want ptr to Result struct passed through",
			modelType: r,
			expected:  rStr,
			err:       nil,
		},
		{
			name:      "Want ptr to Series struct passed through",
			modelType: s,
			expected:  sStr,
			err:       nil,
		},
		{
			name:      "Default case ptr to struct didn't match",
			expected:  "",
			modelType: f,
			err:       NoTypesMatched,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := testCheckTypeSwitchAssertion(t, tc.modelType)

			if err != nil {
				assert.ErrorIs(t, err, NoTypesMatched)
			}

			assert.Equal(t, got, tc.expected)
		})
	}
}

var NoTypesMatched error = errors.New("none of the types matched")

func testCheckTypeSwitchAssertion(t *testing.T, md any) (string, error) {
	t.Helper()
	switch model := md.(type) {
	case *mangaupdates.Result:
		return fmt.Sprintf("Result type matched: %T", model), nil
	case *mangaupdates.Series:
		return fmt.Sprintf("Series type matched: %T", model), nil
	case *mangaupdates.SearchResults:
		return fmt.Sprintf("SearchResult type matched: %T", model), nil
	default:
		return "", NoTypesMatched
	}
}

// TODO: implement test for search function on successful search, 400 bad request search, 200 ok but empty search
// func TestRequestAndDecode(t *testing.T) {
// 	t.Parallel()
// 	testCases := []struct {
// 		name          string
// 		statusCode    int
// 		responseBody  string
// 		expected      mangaupdates.Response
// 		expectedError bool
// 	}{
// 		{
// 			name:       "400 validation error",
// 			statusCode: http.StatusBadRequest,
// 			responseBody: `{
// 				"status": "exception",
// 				"reason": "Field Validation Error",
// 				"context": {
// 					"search": [
// 						{
// 							"index": 0,
// 							"errors": ["\"\" must have a length between 1 and 400"]
// 						}
// 					],
// 					"exclude": [
// 						{
// 							"index": 0,
// 							"errors": ["\"\" must have a length between 1 and 100"]
// 						}
// 					]
// 				}
// 			}`,
// 			expected: mangaupdates.Response{
// 				Status: "exception",
// 				Reason: "Field Validation Error",
// 				Context: map[string][]mangaupdates.ContextErrors{
// 					"search": {
// 						{Index: 0, Errors: []string{"\"\" must have a length between 1 and 400"}},
// 					},
// 					"exclude": {
// 						{Index: 0, Errors: []string{"\"\" must have a length between 1 and 100"}},
// 					},
// 				},
// 			},
// 		},
// 		// {
// 		// 	name:         "404 response",
// 		// 	statusCode:   http.StatusNotFound,
// 		// 	responseBody: `{}`,
// 		// 	expected:     mangaupdates.Response{},
// 		// },
// 		// {
// 		// 	name:       "200 success",
// 		// 	statusCode: http.StatusOK,
// 		// 	responseBody: `{
// 		// 		"status": "success",
// 		// 		"reason": "login successful"
// 		// 	}`,
// 		// 	expected: mangaupdates.Response{
// 		// 		Status: "success",
// 		// 		Reason: "login Succesful",
// 		// 	},
// 		// },
// 	}
//
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()
// 			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				w.Header().Set("Content-Type", "application/json")
// 				w.WriteHeader(tc.statusCode)
// 				_, err := w.Write([]byte(tc.responseBody))
// 				require.NoError(t, err)
// 			}))
// 			defer testServer.Close()
//
// 			// muClient := mangaupdates.NewMangaUpdatesClient()
//
//
// 			var got mangaupdates.Result
// 			// err := "TODO: implement search function then test"
// 			// require.NoError(t, err)
// 			assert.Equal(t, tc.expected, got)
// 			assert.Equal(t, tc.statusCode, got.StatusCode)
// 		})
// 	}
//
// }
