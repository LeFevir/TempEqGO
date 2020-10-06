// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"log"
)

const (
	// Множитель для сохранения точности вычисления вещественных чисел в определённых пределах.
	// При выборе множителя = 1000000 можно работать с точностью координат сетки до 1/1000000 метра, то есть до микрометра.
	FLOAT_ACCURACY_MLTPL = 1000000
)

type Field struct {
	Nx int
	Ny int
	Nz int

	MinX float64
	MinY float64
	MinZ float64

	Dx float64
	Dy float64
	Dz float64

	X []float64
	Y []float64
	Z []float64

	Value []float64

	// Тип поля (объем, плоскости xy, yz, xz, оси x, y, z)
	Type string
}

func (field *Field) SetNodesNumber(nX, nY, nZ int) {
	field.Nx = nX
	field.Ny = nY
	field.Nz = nZ
}

func (field *Field) SetNodesSteps(dX, dY, dZ float64) {
	field.Dx = dX
	field.Dy = dY
	field.Dz = dZ
}

func (field *Field) SetGridBottomBorder(minX, minY, minZ float64) {
	field.MinX = minX
	field.MinY = minY
	field.MinZ = minZ
}

//This function prepares the Grid
func (field *Field) PrepareGrid() {
	field.X = make([]float64, field.Nx)
	field.Y = make([]float64, field.Ny)
	field.Z = make([]float64, field.Nz)

	// Так как вещественные числа не сохраняются точно в бинарном виде, то погрешность расчётов растёт при математических операциях.
	// В данном месте программы (задание сетки расчётов) наиболее важна точность координат сетки.
	// Для избежания ошибок вещественные числа умножаются на множитель FLOAT_ACCURACY_MLTPL = ,например, 1000000, преобразуются к целым,
	// производятся математические операции, а затем результат преобразуется обратно в вещественное число и делится на множитель FLOAT_ACCURACY_MLTPL.
	// Таким образом, при выборе множителя = 1000000 можно работать с точностью координат сетки до 1/1000000 метра, то есть до микрометра.

	for i := 0; i < field.Nx; i++ {
		field.X[i] = float64(int(field.MinX*FLOAT_ACCURACY_MLTPL)+i*int(field.Dx*FLOAT_ACCURACY_MLTPL)) / FLOAT_ACCURACY_MLTPL
	}

	for i := 0; i < field.Ny; i++ {
		field.Y[i] = float64(int(field.MinY*FLOAT_ACCURACY_MLTPL)+i*int(field.Dy*FLOAT_ACCURACY_MLTPL)) / FLOAT_ACCURACY_MLTPL
	}

	for i := 0; i < field.Nz; i++ {
		field.Z[i] = float64(int(field.MinZ*FLOAT_ACCURACY_MLTPL)+i*int(field.Dz*FLOAT_ACCURACY_MLTPL)) / FLOAT_ACCURACY_MLTPL
	}

	field.Value = make([]float64, field.Nx*field.Ny*field.Nz)
}

//This function gets the indexes of 3D field and outputs the corresponding value at field's point.
//Go hasn't got a multidimensional slices to provide runtime generation of arrays.
//So the values of 3D field have to be held in 1D slice with corresponding indexing.
func (field *Field) Get(i, j, k int) (v float64) {
	ind := i + j*field.Nx + k*field.Nx*field.Ny
	v = field.Value[ind]
	return
}

//This function gets the indexes of 3D field and puts the incoming value into field's point.
//Go hasn't got a multidimensional slices to provide runtime generation of arrays.
//So the values of 3D field have to be held in 1D slice with corresponding indexing.
func (field *Field) Put(v float64, i, j, k int) {
	ind := i + j*field.Nx + k*field.Nx*field.Ny
	field.Value[ind] = v
}

// func (field *Field) MakeCopy() (newField *Field) {
// 	newField = new(Field)
// 	newField.SetNodesNumber(field.Nx, field.Ny, field.Nz)
// 	newField.PrepareGrid()
// 	copy(newField.Value, field.Value)
// 	return
// }

//This function finds maximum value of field and its indexes. Indexes fills FieldComplex struct
func (field *Field) MaxValue() (max float64) {
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

func GenerateFieldAs(field *Field) (newField *Field) {
	newField = new(Field)
	newField.SetNodesNumber(field.Nx, field.Ny, field.Nz)
	newField.SetNodesSteps(field.Dx, field.Dy, field.Dz)
	newField.SetGridBottomBorder(field.MinX, field.MinY, field.MinZ)
	newField.Type = field.Type
	newField.PrepareGrid()
	log.Println("Field As Field has been generated")
	return
}

func (field *Field) FillFieldWithValue(val float64) {
	for i := 0; i < field.Nx; i++ {
		for j := 0; j < field.Ny; j++ {
			for k := 0; k < field.Nz; k++ {
				field.Put(val, i, j, k)
			}
		}
	}
	log.Println("Field has been filled with value")
	return
}
