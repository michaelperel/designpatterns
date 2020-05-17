package observer

import (
	"math/rand"
	"sync"
)

type WeatherSubject struct {
	observers []Observer
}

func (*WeatherSubject) Temperature() int {
	return rand.Intn(100)
}

func (*WeatherSubject) PH() float64 {
	return rand.Float64() * (14)
}

func (w *WeatherSubject) Register(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherSubject) Notify() {
	wg := sync.WaitGroup{}
	for _, o := range w.observers {
		wg.Add(1)
		go func(o Observer) {
			defer wg.Done()
			o.Update(w)
		}(o)
	}
	wg.Wait()
}
