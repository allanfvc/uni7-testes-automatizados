package calculator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Operation struct {
	Operation string  "json:'operation'"
	Result    float64 "json:'result'"
}

//Sum -
func Sum(a, b float64) float64 {
	return a + b
}

//Subtract -
func Subtract(a, b float64) float64 {
	return a - b
}

//Multiply -
func Multiply(a, b float64) float64 {
	return a * b
}

//Divider -
func Divider(a, b float64) float64 {
	return a / b
}

func SumHttp() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a, err := strconv.ParseFloat(r.FormValue("a"), 64)
		b, err := strconv.ParseFloat(r.FormValue("b"), 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading parameters"))
		}
		data := &Operation{
			Operation: fmt.Sprintf("%v+%v", a, b),
			Result:    Sum(a, b),
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading parameters"))
		}
	})
}
