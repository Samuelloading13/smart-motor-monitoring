package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SensorData struct {
	Suhu      float64 `json:"suhu"`
	Tegangan  float64 `json:"tegangan"`
	RPM       int     `json:"rpm"`
	Kecepatan int     `json:"kecepatan"`
}

var dataSensor SensorData

func updateSensorData(_ SensorData, newData SensorData) SensorData {
	return newData
}

// Handler POST
func postSensorData(w http.ResponseWriter, r *http.Request) {
	var input SensorData
	json.NewDecoder(r.Body).Decode(&input)

	// update data
	dataSensor = updateSensorData(dataSensor, input)

	// simpan ke riwayat
	historyData = append(historyData, input)

	fmt.Println("POST dari:", r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Data sensor diterima"})
}

// Handler GET
func getSensorData(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataSensor)
}

var historyData []SensorData

// Status sederhana
func getStatus(data SensorData) string {
	if data.Suhu > 80 && data.Tegangan < 11 {
		return "Mesin Panas dan Aki Lemah"
	} else if data.Tegangan < 11 {
		return "Aki Lemah"
	} else if data.Suhu > 80 {
		return "Mesin Panas"
	}
	return "Normal"
}

// Handler status
func getSensorStatus(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": getStatus(dataSensor),
	})
}

// Handler riwayat
func getSensorHistory(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(historyData)
}

// Main function
func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/api/sensor", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			postSensorData(w, r)
		case http.MethodGet:
			getSensorData(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/sensor/status", getSensorStatus)
	http.HandleFunc("/api/sensor/history", getSensorHistory)

	http.ListenAndServe(":8080", nil)
}
