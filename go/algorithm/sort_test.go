package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type Test[T any] struct {
		name string
		in   []T
		want []T
		less func(i, j int, x []T) bool
	}

	testsInt := []Test[int]{
		{
			name: "random int",
			in:   []int{2, 1, -4, 5, 100, 84, 3, 9, -23, 0},
			want: []int{-23, -4, 0, 1, 2, 3, 5, 9, 84, 100},
			less: func(i, j int, x []int) bool {
				return x[i] < x[j]
			},
		},
	}

	testsFloat64 := []Test[float64]{
		{
			name: "random float64",
			in:   []float64{2, 1.2, -4, 5, 100, 84.3, 3, 9, -23, 0},
			want: []float64{-23, -4, 0, 1.2, 2, 3, 5, 9, 84.3, 100},
			less: func(i, j int, x []float64) bool {
				return x[i] < x[j]
			},
		},
	}

	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.in, tt.less)
			if !reflect.DeepEqual(tt.want, tt.in) {
				t.Errorf("expected %v, got %v", tt.want, tt.in)
			}
		})
	}

	for _, tt := range testsFloat64 {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.in, tt.less)
			if !reflect.DeepEqual(tt.want, tt.in) {
				t.Errorf("expected %v, got %v", tt.want, tt.in)
			}
		})
	}
}
