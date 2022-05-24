package driver

import (
	"database/sql"
	"testing"
)

func TestDriverRegistration(t *testing.T) {
	for _, d := range sql.Drivers() {
		if d == "manticore" {
			return
		}
	}

	t.Error("manticore driver not registered")
}

func TestOpen(t *testing.T) {
	_, err := sql.Open("manticore", "")
	if err != nil {
		t.Error(err)
	}
}
