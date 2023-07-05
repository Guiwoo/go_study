package observer

import "fmt"

type Subject interface {
	registerObserver(o Observer2)
	removeObserver(o Observer2)
	notifyObserver()
}

type Observer2 interface {
	update()
}
type DisplayElement interface {
	display()
}

type WeatherData struct {
	observers                []Observer2
	temp, humidity, pressure float32
}

func (w *WeatherData) registerObserver(o Observer2) {
	if w.observers == nil {
		w.observers = make([]Observer2, 0)
	}
	w.observers = append(w.observers, o)
}

func (w *WeatherData) removeObserver(o Observer2) {
	for i, v := range w.observers {
		if v == o {
			if i != len(w.observers)-1 {
				// 일반 슬라이스
				tmp := w.observers[:i]
				w.observers = append(tmp, w.observers[i+1:]...)
			} else {
				w.observers = w.observers[:len(w.observers)-1]
			}
		}
	}
}

func (w *WeatherData) notifyObserver() {
	for _, v := range w.observers {
		v.update()
	}
}

func (w *WeatherData) getTemperature() {

}
func (w *WeatherData) getHumidity() {

}
func (w *WeatherData) getPressure() {

}
func (w *WeatherData) measurementChanged() {
	w.notifyObserver()
}
func (w *WeatherData) setMeasurement(tmp, humidity, pressure float32) {
	w.temp = tmp
	w.humidity = humidity
	w.pressure = pressure
	w.measurementChanged()
}

var _ Subject = (*WeatherData)(nil)

type CurrentConditionDisplay struct {
	temperature, humidity float32
	weatherData           *WeatherData
}

func (c *CurrentConditionDisplay) update() {
	c.temperature = c.weatherData.temp
	c.humidity = c.weatherData.humidity
	c.display()
}
func (c *CurrentConditionDisplay) display() {
	fmt.Printf("현재 상태 온도 : %f, 습도 : %f\n", c.temperature, c.humidity)
}
func NewCurrentConditionDisplay(w *WeatherData) *CurrentConditionDisplay {
	ob := &CurrentConditionDisplay{weatherData: w}
	w.registerObserver(ob)
	return ob
}

type StatisticDisplay struct {
	average, highest, lowest float32
	weatherData              *WeatherData
}

func (s *StatisticDisplay) update() {
	temp := s.weatherData.temp
	s.average = (s.average + temp) / 2

	if s.highest < temp {
		s.highest = temp
	}
	if s.lowest > temp {
		s.lowest = temp
	}
	s.display()
}
func (s *StatisticDisplay) display() {
	fmt.Printf("평균/최고/최저 온도 : %f / %f / %f \n", s.average, s.highest, s.lowest)
}
func NewStatisticDisplay(w *WeatherData) *StatisticDisplay {
	ob := &StatisticDisplay{weatherData: w}
	w.registerObserver(ob)
	return ob
}

type ForecastDisplay struct {
	prevTemp, prevHumidity, prevPressure float32
	announcement                         string
	weatherData                          *WeatherData
}

func (f *ForecastDisplay) update() {
	temp := f.weatherData.temp
	humidity := f.weatherData.humidity
	pressure := f.weatherData.pressure
	if temp > f.prevTemp && humidity > f.prevHumidity && pressure > f.prevPressure {
		f.announcement = "날씨가 무척 더워질 예정입니다. 조심하세요"
	} else if temp < f.prevTemp && humidity < f.prevHumidity && pressure < f.prevPressure {
		f.announcement = "날씨가 무척 추워질 예정입니다. 조심하세요"
	} else {
		f.announcement = "어제와 유사한 날씨 입니다."
	}
	f.prevTemp = temp
	f.prevHumidity = humidity
	f.prevPressure = pressure
	f.display()
}
func (f *ForecastDisplay) display() {
	fmt.Println(f.announcement)
}
func NewForecastDisplay(w *WeatherData) *ForecastDisplay {
	ob := &ForecastDisplay{weatherData: w}
	w.registerObserver(ob)
	return ob
}
