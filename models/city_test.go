package models_test

import (
	"models.go/models"
	"testing"
)

func TestCity(t *testing.T) {
	expectedName := "San Francisco"
	temperatures := []float64{-5, 10, 20}
	city := models.NewCity(expectedName, temperatures, true, true)

	t.Run("name", func(t *testing.T) {
		got := city.Name()
		if got != expectedName {
			t.Errorf("got %q, want %q", got, expectedName)
		}
	})

	t.Run("temperatures", func(t *testing.T) {
		cq := createQuery(t, true, true, 2, "")
		want := temperatures[1]
		got := city.TempC(cq)
		if got != want {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})

	t.Run("temperatures", func(t *testing.T) {
		cq := createQuery(t, true, true, 2, "")
		want := (temperatures[1] * 9 / 5) + 32
		got := city.TempF(cq)
		if got != want {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})

	t.Run("temperatures", func(t *testing.T) {
		cq := createQuery(t, false, true, 1, "")
		want := true
		got := city.SkiVacationReady(cq)
		if got != want {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})

	t.Run("temperatures", func(t *testing.T) {
		cq := createQuery(t, true, false, 2, "")
		want := false
		got := city.BeachVacationReady(cq)
		if got != want {
			t.Errorf("Expected '%v', but got '%v'", want, got)
		}
	})
}

func createQuery(t *testing.T, beach bool, ski bool, month int, name string) models.CityQuery {
	t.Helper()
	cq, err := models.NewQuery(beach, ski, month, name)
	if err != nil {
		t.Fatalf("Error creating city query:%v", err)
	}
	return cq
}
