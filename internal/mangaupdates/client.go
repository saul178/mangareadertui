package mangaupdates

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// TODO:figure out if authtoken is needed, if needed create a default client and an authclient
type MangaUpdatesClient struct {
	client    *http.Client
	header    http.Header
	baseURL   string
	AuthToken string
}

func NewMangaUpdatesClient() *MangaUpdatesClient {
	client := http.Client{}
	header := http.Header{}
	header.Set("Content-Type", "applicaton/json")

	return &MangaUpdatesClient{
		client:  &client,
		header:  header,
		baseURL: BaseAPIURL,
	}
}

// need test soon for these funcs to see if they work as intended
func (r *Response) GetResponsDetails() string {
	var b strings.Builder
	if r.Status != "" || r.Reason != "" {
		fmt.Fprintf(&b, "[STATUS CODE]: (%d)\n[STATUS]: %s\n[REASON]: %s\n", r.StatusCode, r.Status, r.Reason)
	}
	// if context is empty we can assume that we didnt encountered errors since the api only returns Contexts on errors
	if len(r.Context) == 0 {
		return b.String()
	}

	fields := make([]string, 0, len(r.Context))
	for c := range r.Context {
		fields = append(fields, c)
	}

	for _, f := range fields {
		for _, ce := range r.Context[f] {
			fmt.Fprintf(&b, "%s: %s\n", f, strings.Join(ce.Errors, "; "))
		}
	}

	return b.String()

}

func (muc *MangaUpdatesClient) DoRequest(ctx context.Context, method string, url string, body io.Reader) (any, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = muc.header
	resp, err := muc.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()

		var er Response
		if err = json.NewDecoder(resp.Body).Decode(&er); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("non 200 http response: %v", er.GetResponsDetails())
	}
	return resp, nil
}
