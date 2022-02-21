package citcall

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	GetCountriesEndpoint = "/test/countries.json"
)

var (
	ErrBodyEmpty = errors.New("http response body is empty")
)

type httpClient struct {
	Cli     *http.Client
	BaseURL string
}

func NewHttpClient(cli *http.Client, baseURL string) *httpClient {
	return &httpClient{
		Cli:     cli,
		BaseURL: baseURL,
	}
}

func (hc *httpClient) GetCountries(ctx context.Context) (Countries, error) {
	targetURL := fmt.Sprintf("%v%v", hc.BaseURL, GetCountriesEndpoint)
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("action=GetCountries err=%v", err)
	}
	req = req.WithContext(ctx)
	resp, err := hc.Cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("action=GetCountries.doReq err=%v", err)
	}
	if resp.Body == nil {
		return nil, ErrBodyEmpty
	}
	defer resp.Body.Close()
	var countries Countries
	err = json.NewDecoder(resp.Body).Decode(&countries)
	if err != nil {
		return nil, fmt.Errorf("action=GetCountries.unmarshalResponse err=%v", err)
	}
	return countries, nil
}
