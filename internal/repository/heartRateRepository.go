package repository

import (
	"errors"
	"github.com/leandersteiner/iot-backend/internal/model"
	"sync"
)

type HeartRateLog struct {
	Mut        sync.RWMutex
	HeartRates []model.HeartRate
}

func (hrl *HeartRateLog) Add(rate model.HeartRate) {
	hrl.Mut.Lock()
	defer hrl.Mut.Unlock()
	hrl.HeartRates = append(hrl.HeartRates, rate)
}

func (hrl *HeartRateLog) Get(index int) (model.HeartRate, error) {
	if index >= len(hrl.HeartRates) {
		return model.HeartRate{}, errors.New("index out of bounds")
	}
	hrl.Mut.RLock()
	defer hrl.Mut.RUnlock()
	return hrl.HeartRates[index], nil
}

func (hrl *HeartRateLog) GetLast() (model.HeartRate, error) {
	if len(hrl.HeartRates) == 0 {
		return model.HeartRate{}, errors.New("range out of bounds")
	}
	hrl.Mut.RLock()
	defer hrl.Mut.RUnlock()
	return hrl.HeartRates[len(hrl.HeartRates)-1], nil
}

func (hrl *HeartRateLog) GetRange(count int) ([]model.HeartRate, error) {
	if count >= len(hrl.HeartRates) {
		return nil, errors.New("range out of bounds")
	}
	hrl.Mut.RLock()
	defer hrl.Mut.RUnlock()
	result := hrl.HeartRates[len(hrl.HeartRates)-count:]
	return result, nil
}

func (hrl *HeartRateLog) GetAverageOfRange(count int) (float32, error) {
	hrl.Mut.RLock()
	defer hrl.Mut.RUnlock()
	var result float32 = 0
	rates, err := hrl.GetRange(count)
	size := len(rates)
	if size <= 0 {
		return result, errors.New("no data available")
	}

	if err != nil {
		return result, err
	}
	for _, rate := range rates {
		result += rate.Rate
	}
	result /= float32(size)
	return result, nil
}
