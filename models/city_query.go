package models

import "fmt"

type query struct {
	beach bool
	ski   bool
	month int
	name  string
}

type CityQuery interface {
	Beach() bool
	Ski() bool
	Month() int
	Name() string
}

func NewQuery(br bool, sr bool, m int, n string) (CityQuery, error) {
	q := &query{
		beach: br,
		ski:   sr,
		month: m,
		name:  n,
	}
	//fmt.Errorf("NewQuery not implemented")
	err := q.validate()
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (q query) validate() error {
	if q.month < 1 || q.month > 12 {
		return fmt.Errorf("invalid month should be from [1:12] got %v", q.month)
	}
	return nil
}

func (q query) Beach() bool {
	return q.beach
}

func (q query) Ski() bool {
	return q.ski
}

func (q query) Month() int {
	return q.month
}

func (q query) Name() string {
	return q.name
}
