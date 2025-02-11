package service

import (
	"reflect"
	"testing"
)

func TestDefaultPackCalculator_Calculate(t *testing.T) {
	tests := []struct {
		name        string
		orderAmount int
		packSizes   []int
		want        map[int]int
	}{
		{
			name:        "Single item",
			orderAmount: 1,
			packSizes:   []int{250, 500, 1000, 2000, 5000},
			want:        map[int]int{250: 1}, // Smallest pack size for minimum items
		},
		{
			name:        "Exact match with single pack",
			orderAmount: 250,
			packSizes:   []int{250, 500, 1000, 2000, 5000},
			want:        map[int]int{250: 1},
		},
		{
			name:        "Need to round up to next pack size",
			orderAmount: 251,
			packSizes:   []int{250, 500, 1000, 2000, 5000},
			want:        map[int]int{500: 1}, // Rule 2: minimum items (500) over minimum packs (2x250)
		},
		{
			name:        "Multiple packs needed",
			orderAmount: 501,
			packSizes:   []int{250, 500, 1000, 2000, 5000},
			want:        map[int]int{500: 1, 250: 1}, // Rule 2: 750 items is better than 1000
		},
		{
			name:        "Large order",
			orderAmount: 12001,
			packSizes:   []int{250, 500, 1000, 2000, 5000},
			want:        map[int]int{5000: 2, 2000: 1, 250: 1}, // 12250 items total
		},
	}

	calculator := &DefaultPackCalculator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Calculate(tt.orderAmount, tt.packSizes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackServiceImpl_CalculatePacks(t *testing.T) {
	packSizes := []int{250, 500, 1000, 2000, 5000}
	calculator := &DefaultPackCalculator{}
	service := NewPackService(calculator, packSizes)

	tests := []struct {
		name        string
		orderAmount int
		wantTotal   int
		wantPacks   map[int]int
	}{
		{
			name:        "Simple order",
			orderAmount: 250,
			wantTotal:   250,
			wantPacks:   map[int]int{250: 1},
		},
		{
			name:        "Complex order",
			orderAmount: 751,
			wantTotal:   1000,                 // Rule 2: Minimum items possible
			wantPacks:   map[int]int{1000: 1}, // Rule 3: Fewer packs (1x1000 better than 2x500)
		},
		{
			name:        "Large order",
			orderAmount: 12001,
			wantTotal:   12250,                                 // Rule 2: Minimum items possible
			wantPacks:   map[int]int{5000: 2, 2000: 1, 250: 1}, // Matches example in requirements
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := service.CalculatePacks(tt.orderAmount)
			if response.TotalItems != tt.wantTotal {
				t.Errorf("CalculatePacks() total = %v, want %v", response.TotalItems, tt.wantTotal)
			}
			if response.OrderAmount != tt.orderAmount {
				t.Errorf("CalculatePacks() orderAmount = %v, want %v", response.OrderAmount, tt.orderAmount)
			}
			if !reflect.DeepEqual(response.Packs, tt.wantPacks) {
				t.Errorf("CalculatePacks() packs = %v, want %v", response.Packs, tt.wantPacks)
			}
		})
	}
}
