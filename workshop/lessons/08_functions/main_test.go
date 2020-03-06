package main

import (
	"math"
	"testing"
)

func Test_average(t *testing.T) {
	tests := []struct {
		name    string
		args    []float64
		want    float64
		wantErr bool
	}{
		{
			name: "one",
			args: []float64{1},
			want: 1,
		},
		{
			name:    "empty list",
			args:    []float64{},
			want:    0.0,
			wantErr: true,
		},
		{
			name: "two elements",
			args: []float64{10, 1},
			want: 5.5,
		},
		{
			name: "math.MaxFloat32",
			args: []float64{math.MaxFloat32, math.MaxFloat32, math.MaxFloat32},
			want: math.MaxFloat32,
		},
		{
			name:    "math.MaxFloat64",
			args:    []float64{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
			want:    0.0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := average(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("average() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
