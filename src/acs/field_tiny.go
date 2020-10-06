// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

type FieldTiny struct {
	Nx int
	Ny int
	Nz int

	Value []float64
}

func (field *FieldTiny) SetNodesNumber(nX, nY, nZ int) {
	field.Nx = nX
	field.Ny = nY
	field.Nz = nZ
}

//This function prepares the Grid
func (field *FieldTiny) PrepareGrid() {
	field.Value = make([]float64, field.Nx*field.Ny*field.Nz)
}

//This function gets the indexes of 3D field and outputs the corresponding value at field's point.
//Go hasn't got a multidimensional slices to provide runtime generation of arrays.
//So the values of 3D field have to be held in 1D slice with corresponding indexing.
func (field *FieldTiny) Get(i, j, k int) (v float64) {
	ind := i + j*field.Nx + k*field.Nx*field.Ny
	v = field.Value[ind]
	return
}

//This function gets the indexes of 3D field and puts the incoming value into field's point.
//Go hasn't got a multidimensional slices to provide runtime generation of arrays.
//So the values of 3D field have to be held in 1D slice with corresponding indexing.
func (field *FieldTiny) Put(v float64, i, j, k int) {
	ind := i + j*field.Nx + k*field.Nx*field.Ny
	field.Value[ind] = v
}

func GenerateFillTinyFieldFromField(field *Field) (newField *FieldTiny) {
	newField = new(FieldTiny)
	newField.SetNodesNumber(field.Nx, field.Ny, field.Nz)
	newField.PrepareGrid()
	copy(newField.Value, field.Value)
	return
}

//This function finds maximum value of field and its indexes. Indexes fills FieldComplex struct
func (field *FieldTiny) MaxValue() (max float64) {
	for i := 0; i < field.Nx; i++ {
		for j := 0; j < field.Ny; j++ {
			for k := 0; k < field.Nz; k++ {
				if field.Get(i, j, k) >= max {
					max = field.Get(i, j, k)
				}
			}
		}
	}
	return
}
