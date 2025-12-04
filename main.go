package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// DATA STRUCTURES

type SensorData struct {
	ID        int     `json:"id"`
	Waktu     string  `json:"waktu"`
	Suhu      float64 `json:"suhu"`
	Tegangan  float64 `json:"tegangan"`
	RPM       int     `json:"rpm"`
	Kecepatan int     `json:"kecepatan"`
}

// Global State
var dataSensor SensorData
var historyData []SensorData
var idCounter = 0

// FUNCTIONAL CORE (PURE FUNCTIONS)

// Pure Function: Update data tanpa mutasi in-place
func updateSensorData(_ SensorData, newData SensorData) SensorData {
	return newData
}

// Tipe function untuk aturan status (First-Class Function)
type StatusCheck func(SensorData) (string, bool)

var checkOverheat = func(d SensorData) (string, bool) {
	if d.Suhu > 80 {
		return "Mesin Overheat", true
	}
	return "", false
}

var checkLowVoltage = func(d SensorData) (string, bool) {
	if d.Tegangan < 11 {
		return "Tegangan Kritis", true
	}
	return "", false
}

var checkOverspeed = func(d SensorData) (string, bool) {
	if d.Kecepatan > 100 {
		return "Kecepatan Berbahaya", true
	}
	return "", false
}

// Higher-Order Function: Komposisi status dari berbagai aturan
func resolveStatus(d SensorData, checks ...StatusCheck) string {
	var detectedIssues []string
	for _, check := range checks {
		if msg, ok := check(d); ok {
			detectedIssues = append(detectedIssues, msg)
		}
	}
	if len(detectedIssues) == 0 {
		return "Sistem Normal"
	}
	return strings.Join(detectedIssues, " + ")
}

func createEmptyHistory() []SensorData {
	return []SensorData{}
}

// HTTP HANDLERS

func handleSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var input SensorData
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// LOGIKA ID & WAKTU (REALTIME)
		idCounter++
		input.ID = idCounter
		input.Waktu = time.Now().Format("15:04:05")

		// Update State
		dataSensor = updateSensorData(dataSensor, input)
		historyData = append(historyData, input)

		fmt.Printf("[ID: %d] [%s] Data Masuk: %+v\n", input.ID, input.Waktu, input)
		json.NewEncoder(w).Encode(map[string]string{"msg": "Data updated", "time": input.Waktu})

	case http.MethodGet:
		json.NewEncoder(w).Encode(dataSensor)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	currentStatus := resolveStatus(dataSensor, checkOverheat, checkLowVoltage, checkOverspeed)

	json.NewEncoder(w).Encode(map[string]string{
		"status": currentStatus,
	})
}

func handleHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(historyData)

	case http.MethodDelete:
		historyData = createEmptyHistory()
		idCounter = 0
		fmt.Println("History & ID Counter Reset")
		json.NewEncoder(w).Encode(map[string]string{"msg": "History deleted"})

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/api/sensor", handleSensor)
	http.HandleFunc("/api/sensor/status", handleStatus)
	http.HandleFunc("/api/sensor/history", handleHistory)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
