package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//  curl -X POST -H "Content-Type: application/json" http://localhost:10010/measure  -d '{"temp": 270, "hum": 300}'
//  curl -X GET -H "Content-Type: application/json" http://localhost:80/hi

type PostMeasureResponse struct {
	Temp int `json:"temp"`
	Hum  int `json:"hum"`
}

func decodeBody(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /measure", r.GetBody)
	var m PostMeasureResponse

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Temperature: %02d.%d°C, ", m.Temp/10, m.Temp%10)
	fmt.Printf("Humidity: %02d.%d%%\n", m.Hum/10, m.Hum%10)
}

func main() {
	fmt.Println("Listening on port 10010")
	mux := http.NewServeMux()

	// POST /measure
	mux.HandleFunc("/measure", decodeBody)

	// GET /hi
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello, you")
	})

	err := http.ListenAndServe(":10010", mux)
	if err != nil {
		panic(err)
	}
}
