package models_test

import (
	"models.go/models"
	"testing"
)

func TestCityQuery(t *testing.T) {
	tests := []struct {
		testName string
		beach    bool
		ski      bool
		month    int
		name     string
		wantErr  string
	}{
		{testName: "valid beach", beach: true, ski: false, month: 1, name: "", wantErr: ""},
		{testName: "valid ski", beach: false, ski: true, month: 1, name: "", wantErr: ""},
	}

	verifyField := func(t *testing.T, want interface{}, got interface{}) {
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			cq, err := models.NewQuery(tt.beach, tt.ski, tt.month, tt.name)
			if err != nil {
				verifyField(t, tt.wantErr, err.Error())
				return
			}
			verifyField(t, tt.beach, cq.Beach())
		})
	}
}
