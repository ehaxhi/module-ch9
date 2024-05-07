package models

import (
	"data.go/data"
	"sort"
	"strings"
)

type cities struct {
	cityMap map[string]CityTemp
}

type Cities interface {
	Filter(cq CityQuery) []CityTemp
}

func (c cities) Filter(cq CityQuery) []CityTemp {
	if !cq.Beach() && !cq.Ski() && cq.Month() == 0 && len(cq.Name()) == 0 {
		return c.listAll()
	}
	return c.filterHelper(cq)
}

func (c cities) filterHelper(cq CityQuery) []CityTemp {
	var cs []CityTemp
	for _, rc := range c.cityMap {
		if matchFilter(rc, cq) {
			cs = append(cs, rc)
		}
	}
	sortAlphabetically(cs)
	return cs
}

func matchFilter(rc CityTemp, cq CityQuery) bool {
	if cq.Beach() && rc.BeachVacationReady(cq) {
		return true
	}
	if cq.Ski() && rc.SkiVacationReady(cq) {
		return true
	}
	if cq.Name() != "" && strings.Contains(strings.ToLower(rc.Name()), strings.ToLower(cq.Name())) {
		return true
	}
	return false
}

func NewCities(reader data.DataReader) (Cities, error) {
	d, err := reader.ReadData()
	if err != nil {
		return nil, err
	}
	cmap := make(map[string]CityTemp)
	for _, r := range d {
		cmap[r.Id] = NewCity(r.Name, r.TempC, r.HasBeach, r.HasMountain)
	}
	return &cities{
		cityMap: cmap,
	}, nil
}

func (c cities) listAll() []CityTemp {
	var cs []CityTemp
	for _, v := range c.cityMap {
		cs = append(cs, v)
	}
	sortAlphabetically(cs)
	return cs
}

func sortAlphabetically(cs []CityTemp) {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Name() < cs[j].Name()
	})
}
