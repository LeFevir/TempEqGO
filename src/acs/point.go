// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics 
// Lomonosov Moscow State University

package acs

import ()

type Point struct {
	X, Y, Z float64
}

func NewPoint(x, y, z float64) (point Point) {
	point.X = x
	point.Y = y
	point.Z = z
	return
}
