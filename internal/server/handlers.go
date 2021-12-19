package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/leandersteiner/iot-backend/internal/model"
	"github.com/leandersteiner/iot-backend/internal/repository"
	"io/ioutil"
	"net/http"
	"strconv"
)

var heartRateRepo = repository.HeartRateLog{
	HeartRates: make([]model.HeartRate, 0),
}

func healthCheck(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(`{"status":"ok"}`))
}

func getHeartRate(rw http.ResponseWriter, req *http.Request) {
	heartRates, err := heartRateRepo.GetLast()
	if err != nil {
		rw.WriteHeader(400)
		return
	}
	jsonResonse, err := json.Marshal(heartRates)
	if err != nil {
		rw.WriteHeader(500)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonResonse)
}

func getHeartRateRange(rw http.ResponseWriter, req *http.Request) {
	count, err := strconv.Atoi(chi.URLParam(req, "count"))
	if err != nil || count <= 0 {
		rw.WriteHeader(400)
	}
	heartRates, err := heartRateRepo.GetRange(count)
	if err != nil {
		rw.WriteHeader(400)
	}
	jsonResponse, err := json.Marshal(heartRates)
	if err != nil {
		rw.WriteHeader(500)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonResponse)
}

func getAllHeartRates(rw http.ResponseWriter, req *http.Request) {

	heartRates, err := heartRateRepo.GetAll()
	if err != nil {
		rw.WriteHeader(400)
	}
	jsonResponse, err := json.Marshal(heartRates)
	if err != nil {
		rw.WriteHeader(500)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonResponse)
}

func postHeartRate(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	var heartRates []model.HeartRate
	err = json.Unmarshal(body, &heartRates)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	for _, rate := range heartRates {
		heartRateRepo.Add(rate)
	}
	rw.Write([]byte(`{"status":"ok"}`))
}
