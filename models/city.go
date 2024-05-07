package models

var (
	beachVacationThreshold float64 = 22
	skiVacationThreshold   float64 = -2
)

type city struct {
	name        string
	tempC       []float64
	hasBeach    bool
	hasMounting bool
}

type CityTemp interface {
	Name() string
	TempC(q CityQuery) float64
	TempF(q CityQuery) float64
	BeachVacationReady(q CityQuery) bool
	SkiVacationReady(q CityQuery) bool
}

func NewCity(n string, t []float64, hb bool, hm bool) CityTemp {
	return &city{
		name:        n,
		tempC:       t,
		hasBeach:    hb,
		hasMounting: hm,
	}
}

func (c city) Name() string {
	return c.name
}

func (c city) TempC(q CityQuery) float64 {
	return c.tempC[q.Month()-1]
}

func (c city) TempF(q CityQuery) float64 {
	return (c.tempC[q.Month()-1] * 9 / 5) + 32
}
func (c city) BeachVacationReady(q CityQuery) bool {
	return c.hasBeach && c.tempC[q.Month()-1] > beachVacationThreshold
}
func (c city) SkiVacationReady(q CityQuery) bool {
	return c.hasMounting && c.tempC[q.Month()-1] < skiVacationThreshold
}
