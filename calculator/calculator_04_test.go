package calculator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func TestCalculatorHttp(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/sum", SumHttp())
	ts := httptest.NewServer(r)
	defer ts.Close()

	assertCorrectMessage := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("0.3 plus 0.2 equals 0.5", func(t *testing.T) {
		res, err := mapRequestToValue(ts, "/v1/sum?a=0.3&b=0.2")
		if err != nil {
			t.Errorf(fmt.Sprintf("%v", err))
			return
		}
		got := roundToTwoDecimalPoints(res.Result)
		want := 0.5
		assertCorrectMessage(t, got, want)
	})

	t.Run("3.5 plus 2.5 equals 6.0", func(t *testing.T) {
		res, err := mapRequestToValue(ts, "/v1/sum?a=3.5&b=2.5")
		if err != nil {
			t.Errorf(fmt.Sprintf("%v", err))
			return
		}
		got := roundToTwoDecimalPoints(res.Result)
		want := 6.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("0.2 plus 0.04 equals 0.24", func(t *testing.T) {
		res, err := mapRequestToValue(ts, "/v1/sum?a=0.2&b=0.04")
		if err != nil {
			t.Errorf(fmt.Sprintf("%v", err))
			return
		}
		got := roundToTwoDecimalPoints(res.Result)
		want := 0.24
		assertCorrectMessage(t, got, want)
	})

	t.Run("0.36 plus 0.04 equals 0.4", func(t *testing.T) {
		res, err := mapRequestToValue(ts, "/v1/sum?a=0.36&b=0.04")
		if err != nil {
			t.Errorf(fmt.Sprintf("%v", err))
			return
		}
		got := roundToTwoDecimalPoints(res.Result)
		want := 0.4
		assertCorrectMessage(t, got, want)
	})

	t.Run("0.68 plus 0.04 equals 0.72", func(t *testing.T) {
		res, err := mapRequestToValue(ts, "/v1/sum?a=0.68&b=0.04")
		if err != nil {
			t.Errorf(fmt.Sprintf("%v", err))
			return
		}
		got := roundToTwoDecimalPoints(res.Result)
		want := 0.72
		assertCorrectMessage(t, got, want)
	})
}

func mapRequestToValue(ts *httptest.Server, url string) (*Operation, error) {
	result := &Operation{}
	res, err := http.Get(ts.URL + url)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, &result)
	return result, nil
}

func TestHttpSum(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/sum", sum())
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/v1/sum?a=5&b=5")
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}
	result := &Operation{}
	json.Unmarshal(body, &result)
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	if result.Result != 10.0 {
		t.Errorf("Expected %f, received %f", 10.0, result.Result)
	}
}

func sum() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a, err := strconv.ParseFloat(r.FormValue("a"), 64)
		b, err := strconv.ParseFloat(r.FormValue("b"), 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading bookmarks"))
		}
		data := &Operation{
			Operation: fmt.Sprintf("%v+%v", a, b),
			Result:    Sum(a, b),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading bookmarks"))
		}
	})
}
