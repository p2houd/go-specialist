package main

import (
	"strings"
	"testing"
)

func TestCheckOrders(t *testing.T) {
	tests := []struct {
		test_name string
		name      string
		input     string
		quantity  string
		fio       string
		tel       string
		ind       string
		city      string
		street    string
		bld       string
		fl        string
		valid     bool
	}{
		{
			test_name: "just ok",
			name:      "cocacola",
			quantity:  "131",
			fio:       "ivan",
			tel:       "8800200060",
			ind:       "100023",
			city:      "Moscow",
			street:    "Vorontsovska",
			bld:       "35",
			fl:        "5",
			valid:     true,
		},
		{
			test_name: "invalid name",
			name:      "coc1acola",
			quantity:  "131",
			fio:       "ivan",
			tel:       "8800200060",
			ind:       "100023",
			city:      "Moscow",
			street:    "Vorontsovska",
			bld:       "35",
			fl:        "5",
			valid:     false,
		},
		{
			test_name: "wrong tel",
			name:      "cocacola",
			quantity:  "131",
			fio:       "ivan",
			tel:       "880020006023",
			ind:       "100023",
			city:      "Moscow",
			street:    "Vorontsovska",
			bld:       "35",
			fl:        "5",
			valid:     false,
		},
		{
			test_name: "wrong index",
			name:      "cocacola",
			quantity:  "131",
			fio:       "ivan",
			tel:       "8800200060",
			ind:       "10002233",
			city:      "Moscow",
			street:    "Vorontsovska",
			bld:       "35",
			fl:        "5",
			valid:     false,
		},
		{
			test_name: "wrong name",
			name:      strings.Repeat("cocacola", 100),
			quantity:  "131",
			fio:       "ivan",
			tel:       "8800200060",
			ind:       "10002233",
			city:      "Moscow",
			street:    "Vorontsovska",
			bld:       "35",
			fl:        "5",
			valid:     false,
		},
		{
			test_name: "empty name",
			name:      "",
			quantity:  "131",
			fio:       "ivan",
			tel:       "8800200060",
			ind:       "10002233",
			city:      "Moscow",
			street:    "Vorontsovska",
			bld:       "35",
			fl:        "5",
			valid:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			o := new(Order)
			gotOutput := o.New(tt.name, tt.quantity, tt.fio, tt.tel, tt.ind, tt.city, tt.street, tt.bld, tt.fl)
			if gotOutput == nil == tt.valid {
				t.Errorf("Order is not created but expected to be")
			}
		})
	}
}
