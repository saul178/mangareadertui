package mangaupdates_test


import (
	"fmt"
	"strings"
	"testing"
)

type testErrs struct {
	statusCode int
	status string
	reason string
	contexts map[string][]contextErrors

}

type contextErrors struct {
	err []string
}

/** 
TODO: this is just to test that these functions return what i expect
1.) create a mock test for the api on actual errors where the api expects certain requirements
2.) normal 404 where no errors is reported EX: series not found is 404 but thats not an error so it returns just 404 and no response body
3.) mock test for a successful api response
**/
func TestGetResponseDetails(t *testing.T) {
	contextMap := make(map[string][]contextErrors)
	contextMap["search"] = []contextErrors{
		{
			err: []string{"search needs more than 1 char"},
		},
	}
	contextMap["exclude_genre"] = []contextErrors{
		{
			err: []string{"must have a length between 1 and 100"},
		},
	}

	testErrs := &testErrs{
		statusCode: 404,
		status: "exception",
		reason: "mocking the api response structs",
		contexts: contextMap,
	}

	fmt.Printf("context map LEN: %d\n" ,len(contextMap))
	fmt.Println(testErrs.testGetErrors())
}

func (t *testErrs) testGetErrors() string {
	var b strings.Builder
	if t.status != "" || t.reason != "" {
		fmt.Fprintf(&b, "STATUS CODE: %d\nSTATUS: %s\nREASON: %s", t.statusCode, t.status, t.reason)
	}

	if len(t.contexts) < 1 {
		return b.String()
	}

	fields := make([]string, 0, len(t.contexts))
	for c := range t.contexts {
		fields = append(fields, c)
	}

	for _, f := range fields {
		for _, ce := range t.contexts[f] {
			fmt.Fprintf(&b, "\n%s: %s", f, strings.Join(ce.err, "; "))
		}
	}

	return b.String()

}

/*
TODO:
create a test client and send requests to endpoints and see if we get a response back either by httpstatuscode or getting a struct of info back
*/
func TestDoRequest(t *testing.T) {}

