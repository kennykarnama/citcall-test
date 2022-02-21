package api_test

import (
	"context"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kennykarnama/citcall-test/api"

	"github.com/kennykarnama/citcall-test/client/citcall"

	"net/http"

	"os"

	mockCitcall "github.com/kennykarnama/citcall-test/mocks/client/citcall"
	"github.com/stretchr/testify/mock"
)

func init() {
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
}

func TestHandlerGetCountriesSuccess(t *testing.T) {
	fmt.Println(os.Getwd())
	// mock citcall client
	mockCitcallClient := &mockCitcall.Client{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	countriesData := []*citcall.Country{
		&citcall.Country{
			Name:     "test-A",
			DialCode: "+92",
			IsoCode:  "222",
			Flag:     "http://test.png",
		},
	}
	mockCitcallClient.On("GetCountries", mock.Anything).Return(citcall.Countries(countriesData), nil)
	countries, err := mockCitcallClient.GetCountries(ctx)
	if err != nil {
		t.Errorf("found error: %v", err)
		return
	}
	if len(countries) == 0 {
		t.Errorf("Expected countries len is: %v but found: %v", 1, 0)
		return

	}
	httpHandler := api.NewHttpHandler(mockCitcallClient, 10*time.Second)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	req, err := http.NewRequest("GET", "/countries", nil)
	if err != nil {
		t.Errorf("found error: %v", err)
		return
	}
	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(httpHandler.GetCountries)
	handlerFunc.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		fmt.Println(rr.Body.String())
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandlerGetCountriesStatusNoContent(t *testing.T) {
	fmt.Println(os.Getwd())
	// mock citcall client
	mockCitcallClient := &mockCitcall.Client{}
	mockCitcallClient.On("GetCountries", mock.Anything).Return(nil, errors.New("error"))
	httpHandler := api.NewHttpHandler(mockCitcallClient, 10*time.Second)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	req, err := http.NewRequest("GET", "/countries", nil)
	if err != nil {
		t.Errorf("found error: %v", err)
		return
	}
	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(httpHandler.GetCountries)
	handlerFunc.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
		return
	}
}

func TestHandlerGetCountriesError(t *testing.T) {
	fmt.Println(os.Getwd())
	// mock citcall client
	mockCitcallClient := &mockCitcall.Client{}
	mockCitcallClient.On("GetCountries", mock.Anything).Return(nil, citcall.ErrBodyEmpty)
	httpHandler := api.NewHttpHandler(mockCitcallClient, 10*time.Second)
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	req, err := http.NewRequest("GET", "/countries", nil)
	if err != nil {
		t.Errorf("found error: %v", err)
		return
	}
	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(httpHandler.GetCountries)
	handlerFunc.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
		return
	}
}
