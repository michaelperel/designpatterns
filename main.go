package main

import (
	"fmt"
	"github.com/michaelperel/designpatterns/strategy"
)

func main() {
	// Strategy Design Pattern
	sqd := strategy.NewSlowQuietDuck()
	fmt.Println(sqd.Display(), sqd.Fly(), sqd.Quack(), sqd.Swim(), sqd.Sleep())

	fld := strategy.NewFastLoudDuck()
	fmt.Println(fld.Display(), fld.Fly(), fld.Quack(), fld.Swim(), fld.Run())
}
