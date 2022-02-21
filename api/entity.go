package api

import (
	"encoding/json"
)

type ErrorMessageResponse struct {
	Message string `json:"message"`
}

func (em *ErrorMessageResponse) Bytes() []byte {
	b, _ := json.Marshal(em)
	return b
}

type GetCountriesReponse struct {
	Data []*CountryData `json:"data"`
}

type CountryData struct {
	Name     string `json:"name"`
	DialCode string `json:"dialCode"`
	IsoCode  string `json:"isoCode"`
	Flag     string `json:"flag"`
}
