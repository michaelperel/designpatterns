package observer

import (
	"fmt"
)

type TempDisplay struct{}

func (*TempDisplay) Update(o interface{}) {
	w, ok := o.(*WeatherSubject)
	if !ok {
		return
	}
	fmt.Println("Temp: ", w.Temperature())
}

type PHDisplay struct{}

func (*PHDisplay) Update(o interface{}) {
	w, ok := o.(*WeatherSubject)
	if !ok {
		return
	}
	fmt.Println("PH: ", w.PH())
}
