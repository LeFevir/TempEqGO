// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"math"
)

func GenerateTestField() (tempField *Field) {
	tempField = new(Field)
	tempField.SetNodesNumber(101, 101, 101)
	tempField.SetNodesSteps(0.1e-03, 0.1e-03, 0.1e-03)
	tempField.SetGridBottomBorder(-5.0e-03, -5.0e-03, -5.0e-03)
	tempField.PrepareGrid()
	return
}

func (tempField *Field) SetInitTemp(T0 float64, r0 float64) {
	for i := 0; i < tempField.Nx; i++ {
		for j := 0; j < tempField.Ny; j++ {
			for k := 0; k < tempField.Nz; k++ {
				t0 := T0 * math.Exp(-(tempField.X[i]*tempField.X[i]+tempField.Y[j]*tempField.Y[j]+tempField.Z[k]*tempField.Z[k])/(r0*r0))
				tempField.Put(t0, i, j, k)
			}
		}
	}
	return
}

func (refTempField *Field) CalcReferenceField(T0 float64, r0 float64, a float64, t float64) {
	for i := 0; i < refTempField.Nx; i++ {
		for j := 0; j < refTempField.Ny; j++ {
			for k := 0; k < refTempField.Nz; k++ {
				T := T0 * math.Pow((1.0/(1.0+4*a*t/(r0*r0))), 3.0/2.0) * math.Exp(-(refTempField.X[i]*refTempField.X[i]+refTempField.Y[j]*refTempField.Y[j]+refTempField.Z[k]*refTempField.Z[k])/(r0*r0+4.0*a*t))
				refTempField.Put(T, i, j, k)
			}
		}
	}
	return
}
