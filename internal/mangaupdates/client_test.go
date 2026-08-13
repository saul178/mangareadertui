package mangaupdates_test

import (
	"testing"

	"github.com/saul178/mangareadertui/internal/mangaupdates"
	// "github.com/stretchr/testify/assert"
)

/**
TODO: this is just to test that these functions return what i expect
1.) create a mock test for the api on actual errors where the api expects certain requirements
2.) normal 404 where no errors is reported EX: series not found is 404 but thats not an error so it returns just 404 and no response body
3.) mock test for a successful api response
**/

// test response struct  that api validation errors are reported,
func TestRequestAndDecode(t *testing.T) {
	mockResp := []mangaupdates.Response{
		{
			StatusCode: 400,
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
		{
			StatusCode: 404,
		},
		{
			StatusCode: 200,
			Status: "success",
			Reason: "login successful",
		},
	}

	mockCases := []struct {
		name string
		expected *mangaupdates.Response
	}{
		{
			name: "400 status code validation or service error returns a response",
			expected: &mockResp[0],
		},
		{
			name: "404 status code returned only, no response body",
			expected: &mockResp[1],
		},
		{
			name: "200 status code returned with status and reason",
			expected: &mockResp[2],
		},
	}
}

/*
TODO:
create a test client and send requests to endpoints and see if we get a response back either by httpstatuscode or getting a struct of info back
*/
func TestDoRequest(t *testing.T) {}

