package belajar_golang_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Muhammad",
		MiddleName: "Abyan",
		LastName: "Kamal",
	}

	encoder.Encode(customer)
}