// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"fmt"
	"log"
	"math"
)

func (tempField *Field) CalcTempAndDose(doseField *Field, intField *Field, medium *Medium, T0 float64, tau float64, numt int) {

	const (
		INTEREST_POINT_I = 18
		INTEREST_POINT_J = 18
		INTEREST_POINT_K = 16

		R0  = 0.5 // Constant for Dose Calc
		R02 = R0 / 2.0
	)

	// Constants for calculations (FOR PERFOMANCE)
	// --------------------------------------
	g1 := (0.5 * medium.ThermalDiffusivity * tau) / (tempField.Dx * tempField.Dx)
	g2 := (0.5 * medium.ThermalDiffusivity * tau) / (tempField.Dy * tempField.Dy)
	g3 := (0.5 * medium.ThermalDiffusivity * tau) / (tempField.Dz * tempField.Dz)
	A1 := g1
	A2 := g2
	A3 := g3
	C1 := 1.0 + 2.0*g1
	C2 := 1.0 + 2.0*g2
	C3 := 1.0 + 2.0*g3

	heatSourcesField := GenerateHeatSourcesFieldFromIntField(intField, medium)

	for i := 1; i < heatSourcesField.Nx; i++ {
		for j := 1; j < heatSourcesField.Ny; j++ {
			for k := 1; k < heatSourcesField.Nz; k++ {
				heatSourcesField.Put(heatSourcesField.Get(i, j, k)*tau, i, j, k) // For Perfomance
			}
		}
	}

	T1 := GenerateFillTinyFieldFromField(tempField)
	T2 := GenerateFillTinyFieldFromField(tempField)
	T3 := GenerateFillTinyFieldFromField(tempField)

	// Main cycle on time --------------------
	for jt := 1; jt < numt; jt++ {

		tempField.krankNicolsonSchemeCalc(g1, g2, g3, A1, A2, A3, C1, C2, C3, T0, T1, T2, T3, heatSourcesField)

		// DOSE CALCULATION
		doseField.doseCalc(tempField, tempField.Get(INTEREST_POINT_I, INTEREST_POINT_J, INTEREST_POINT_K), R0, R02, tau)

		fmt.Printf("Time is %.3f s. Temp in point (%.3f, %.3f, %.3f) mm is %.3f oC\n", float64(jt)*tau, tempField.X[INTEREST_POINT_I]*1000.0, tempField.Y[INTEREST_POINT_J]*1000.0, tempField.Z[INTEREST_POINT_K]*1000.0, tempField.Get(INTEREST_POINT_I, INTEREST_POINT_J, INTEREST_POINT_K))
		log.Printf("Time is %.3f s. Temp in point (%.3f, %.3f, %.3f) mm is %.3f oC\n", float64(jt)*tau, tempField.X[INTEREST_POINT_I]*1000.0, tempField.Y[INTEREST_POINT_J]*1000.0, tempField.Z[INTEREST_POINT_K]*1000.0, tempField.Get(INTEREST_POINT_I, INTEREST_POINT_J, INTEREST_POINT_K))

		// if mod(jt,jPrint)==0 {
		// 	// Printing on Display Main and Side Temperature
		// 	fmt.Println('(f6.2, A, f7.4, A)',T(x_c,y_c,z_c),' in Center Max. Time: ', jt*tau, 's')
		// 	// write (*, '(A, f6.2)') 'Without Diffusion: ', T0+jt*Q(x_c,y_c,z_c)
		// 	write (*, '(f6.2, A)') T(x_c,y_s,z_c),' in Side Max.'
		// 	// write (*, '(A, f7.3)') '-------------Time left (min): ', (numt-jt)*(c_time)/60

		// 	// Printing of TempRise in File
		// 	write (10, '(2x,f7.4,5x,f6.2)') jt*tau, T(x_c,y_c,z_c) !Main Peak
		// 	write (11, '(2x,f7.4,5x,f6.2)') jt*tau, T(x_c,y_s,z_c) !Side Peak
		// }

		// if jt==t_5 {
		// 	fmt.Println('-------------5s---------------')
		// 	write (100) T
		// 	write (110) t56
		// 	// Dose OUT
		// 	write (111, *) numy, numz
		// 	i=x_c
		// 	do j=1,numy
		// 				do k=1,numz
		// 	if (t56(i,j,k)>100.0) then
		// 	write (111, '(2x,f12.5,2x,f12.5,2x,i5)') Y(j), Z(k), 100
		// 	else
		// 	write (111, '(2(2x,f12.5),(2x,f14.2))') Y(j), Z(k), t56(i,j,k)
		// 	endif
		// 	end do
		// 	// write (*, '(A, f6.2, A)') 'Filling Text: ', 100.0*real(i)/real(numx), '% completed'
		// 	end do

		// 	// Temp OUT
		// 	write (101, *) numy, numz
		// 	i=x_c
		// 	do j=1,numy
		// 	do k=1,numz
		// 	write (101, '(2(2x,f12.5),(2x,f6.2))') Y(j), Z(k), T(i,j,k)
		// 	end do
		// 	end do
		// }

	}
	// End of Cycle on Time ------------------

	// fmt.Printf("Numeric field calculation has finished. It took %s\n", time.Since(startTime))
	// log.Printf("Расчет температуры и дозы закончен. Он занял %s\n", time.Since(startTime))
	return
}

func (T *Field) krankNicolsonSchemeCalc(g1, g2, g3, A1, A2, A3, C1, C2, C3 float64, T0 float64, T1, T2, T3, Q *FieldTiny) {

	// Krank-Nicolson Scheme------------------

	// --T1------------------

	xsi := make([]float64, T.Nx)
	eta := make([]float64, T.Nx)

	// Cycle on Y,Z
	for k := 1; k < T.Nz-1; k++ {
		for j := 1; j < T.Ny-1; j++ {
			// boundary conditions for i1
			xsi[T.Nx-1] = 0.0
			eta[T.Nx-1] = T0
			// Calculating for X
			for i := T.Nx - 2; i > 0; i-- {
				xsi[i] = A1 / (C1 - A1*xsi[i+1])
				eta[i] = (A1*eta[i+1] + (T.Get(i, j, k) + g1*(T.Get(i-1, j, k)+T.Get(i+1, j, k)-2.0*T.Get(i, j, k)))) / (C1 - A1*xsi[i+1])
			}

			// Calculating T1
			for i := 0; i < T.Nx-1; i++ {
				T1.Put((xsi[i+1]*T1.Get(i, j, k) + eta[i+1]), i+1, j, k)
			}
		}
	}
	//End of Cycle on Y,Z
	// ---T1---------------------

	// ---T2---------------------
	xsi = make([]float64, T.Ny)
	eta = make([]float64, T.Ny)
	// Cycle on X,Z
	for k := 1; k < T.Nz-1; k++ {
		for i := 1; i < T.Nx-1; i++ {
			// boundary conditions for i2
			xsi[T.Ny-1] = 0.0
			eta[T.Ny-1] = T0
			// Calculating for Y
			for j := T.Ny - 2; j > 0; j-- {
				xsi[j] = A2 / (C2 - A2*xsi[j+1])
				eta[j] = (A2*eta[j+1] + (T1.Get(i, j, k) + g2*(T1.Get(i, j-1, k)+T1.Get(i, j+1, k)-2.0*T1.Get(i, j, k)))) / (C2 - A2*xsi[j+1])
			}

			// Calculating T2
			for j := 0; j < T.Ny-1; j++ {
				T2.Put((xsi[j+1]*T2.Get(i, j, k) + eta[j+1]), i, j+1, k)
			}
		}
	}
	// End of Cycle on X,Z
	// ---T2---------------------

	// ---T3---------------------
	xsi = make([]float64, T.Nz)
	eta = make([]float64, T.Nz)
	// Cycle on X,Y
	for j := 1; j < T.Ny-1; j++ {
		for i := 1; i < T.Nx-1; i++ {

			// boundary conditions for i3
			xsi[T.Nz-1] = 0.0
			eta[T.Nz-1] = T0
			// Calculating for Z
			for k := T.Nz - 2; k > 0; k-- {
				xsi[k] = A3 / (C3 - A3*xsi[k+1])
				eta[k] = (A3*eta[k+1] + (T2.Get(i, j, k) + g3*(T2.Get(i, j, k-1)+T2.Get(i, j, k+1)-2.0*T2.Get(i, j, k)) + Q.Get(i, j, k))) / (C3 - A3*xsi[k+1])
			}

			// Calculating T(j+1)
			for k := 0; k < T.Nz-1; k++ {
				T.Put((xsi[k+1]*T.Get(i, j, k) + eta[k+1]), i, j, k+1)
			}
		}
	}
	// End of Cycle on X,Y
	// ---T3---------------------

	// !WE HAVE T(j+1)!!!!!!
}

func GenerateHeatSourcesFieldFromIntField(intField *Field, medium *Medium) (heatSourcesField *FieldTiny) {
	log.Println("Generating Heat Source Field...")
	heatSourcesField = new(FieldTiny)
	heatSourcesField.SetNodesNumber(intField.Nx, intField.Ny, intField.Nz)
	heatSourcesField.PrepareGrid()

	for i := 1; i < intField.Nx; i++ {
		for j := 1; j < intField.Ny; j++ {
			for k := 1; k < intField.Nz; k++ {
				// Q = 2*alpha*I/ (ro* Cp)
				val := 2 * medium.AbsorptionM * intField.Get(i, j, k) / (medium.Density * medium.HeatCapacity)
				val = val * 2.7 / (0.004 * 0.005)

				val = val * math.Exp(-2*(intField.Z[k])*medium.AbsorptionM)
				heatSourcesField.Put(val, i, j, k)
			}
		}
	}
	log.Println("Generating Heat Source Field has been completed")
	return
}

func (doseField *Field) doseCalc(tempField *Field, T, R0, R02, tau float64) {

	var val float64

	for i := range doseField.X {
		for j := range doseField.Y {
			for k := range doseField.Z {

				if T < 43.0 {
					val = doseField.Get(i, j, k) + tau*math.Pow(R02, (56.0-tempField.Get(i, j, k)))
				} else {
					val = doseField.Get(i, j, k) + tau*math.Pow(R0, (56.0-tempField.Get(i, j, k)))
				}

				if val > 100.0 {
					val = 100.0
				}

				doseField.Put(val, i, j, k)
			}
		}
	}

	return
}
