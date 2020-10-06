// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics 
// Lomonosov Moscow State University

package acs

import (
	"math"
	"testing"
)

//Simple testing for correct Point Creation
func TestNewPoint(t *testing.T) {
	p := NewPoint(1, 1, 1)
	if (p.X != 1) || (p.X != 1) || (p.X != 1) {
		t.Log("Wrong assignment of the coordinates!")
		t.Fail()
	}
}

//Simple benchmarking
func BenchmarkNewPoint(b *testing.B) {
	for i := 0; i < 1e5; i++ {
		NewPoint(math.Pow(math.Sqrt(float64(i)), 102), 1, 1)
	}
}
