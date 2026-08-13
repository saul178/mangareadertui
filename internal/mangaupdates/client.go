package mangaupdates

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// TODO:figure out if authtoken is needed, if needed create a default client and an authclient
type MangaUpdatesClient struct {
	client    *http.Client
	header    http.Header
	baseURL   string
	authToken string
}

func NewMangaUpdatesClient() *MangaUpdatesClient {
	client := http.Client{}
	header := http.Header{}
	header.Set("Content-Type", "application/json")

	return &MangaUpdatesClient{
		client:  &client,
		header:  header,
		baseURL: BaseAPIURL,
	}
}

func (muc *MangaUpdatesClient) doRequest(ctx context.Context, method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header = muc.header

	resp, err := muc.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (muc *MangaUpdatesClient) RequestAndDecode(ctx context.Context, method string, url string, body io.Reader, responseType any) error {
	resp, err := muc.doRequest(ctx, method, url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TODO: assert that a pointer to a struct is being passed to decode()
	err = json.NewDecoder(resp.Body).Decode(responseType)
	return err
}
