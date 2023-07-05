package observer

import (
	"fmt"
	"testing"
)

func Test_01(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a[0:0])
}

func Test_02(t *testing.T) {
	w := &WeatherData{}

	cur := NewCurrentConditionDisplay(w)
	stat := NewStatisticDisplay(w)
	fore := NewForecastDisplay(w)

	fmt.Println(cur, stat, fore)

	w.setMeasurement(80, 22.2, 32.7)
	w.setMeasurement(60, 20.2, 30.7)
	w.setMeasurement(90, 22.2, 32.9)
}
