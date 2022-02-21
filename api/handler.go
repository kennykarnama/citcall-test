package api

import (
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/kennykarnama/citcall-test/client/citcall"
)

type HttpHandler struct {
	citcallClient          citcall.Client
	externalRequestTimeout time.Duration
}

func NewHttpHandler(citcallClient citcall.Client, externalRequestTimeout time.Duration) *HttpHandler {
	return &HttpHandler{
		citcallClient:          citcallClient,
		externalRequestTimeout: externalRequestTimeout,
	}
}

func (hh *HttpHandler) GetCountries(w http.ResponseWriter, req *http.Request) {
	timeoutCtx, _ := context.WithTimeout(context.Background(), hh.externalRequestTimeout)
	countries, err := hh.citcallClient.GetCountries(timeoutCtx)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		errorResp := ErrorMessageResponse{
			Message: err.Error(),
		}
		if err == citcall.ErrBodyEmpty {
			w.WriteHeader(http.StatusNoContent)
		} else {

			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(errorResp.Bytes())
		return
	}
	var countryRespItems []*CountryData
	for _, c := range countries {
		countryRespItems = append(countryRespItems, &CountryData{
			Name:     c.Name,
			DialCode: c.DialCode,
			IsoCode:  c.IsoCode,
			Flag:     c.Flag,
		})
	}
	resp := &GetCountriesReponse{
		Data: countryRespItems,
	}
	// Template
	tmpl, err := template.ParseFiles("./views/countries.html")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		errorResp := ErrorMessageResponse{
			Message: err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorResp.Bytes())
		return
	}
	tmpl.Execute(w, resp)
}
