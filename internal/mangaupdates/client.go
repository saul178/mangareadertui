package mangaupdates

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

// TODO: create a options function so that the user can have the option to be authenticated
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

// [RequestAndDecode()] returns back a defined struct filled after decoding json response
func (muc *MangaUpdatesClient) RequestAndDecode(ctx context.Context, method string, url string, body io.Reader, respModelType any) error {
	resp, err := muc.doRequest(ctx, method, url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TODO: maybe we can make this generic so that we get compiler warnings instead of runtime errors?
	switch mt := respModelType.(type) {
	case *SearchedSeriesResponse:
		err = json.NewDecoder(resp.Body).Decode(mt)
	case *Result:
		err = json.NewDecoder(resp.Body).Decode(mt)
	case *Series:
		err = json.NewDecoder(resp.Body).Decode(mt)
	default:
		return fmt.Errorf("[ERROR] no defined struct type matched: %T\nDid you pass the correct struct type?", mt)
	}

	return err
}

// [EncodeAndRequest()] sends a json object payload to the requested url and returns the response decoded to the appropriate struct
func (muc *MangaUpdatesClient) EncodeAndRequest(ctx context.Context, method string, url string, payload any, respModelType any) error {
	// create body data buffer and encode/write the payload request to the &body
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(payload)
	if err != nil {
		return err

	}
	err = muc.RequestAndDecode(ctx, method, url, &body, respModelType)
	return err
}
