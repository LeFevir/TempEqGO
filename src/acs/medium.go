// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"log"
)

//Structure medium implements the medium (water, tissue, etc) structure with parameters of the medium.
type Medium struct {
	Name                string
	Density             float64
	SpeedOfSound        float64
	HeatCapacity        float64
	ThermalConductivity float64
	AbsorptionDB        float64
	AbsorptionM         float64
	ThermalDiffusivity  float64
}

//Calculates ThermalDiffusivity (hi) from ThermalConductivity (k) by the hi = k/(ro*cp)
func (medium *Medium) CalcThermalDiffusivity() {
	medium.ThermalDiffusivity = medium.ThermalConductivity / (medium.Density * medium.HeatCapacity)
	log.Println("ThermalDiffusivity has been calculated. It is", medium.ThermalDiffusivity)
}

//Gets absorption in dB/cm/MHz and calculate and fill absorptionM class field in 1/m/MHz units
func (medium *Medium) SetAbsorption() {
	medium.AbsorptionM = medium.AbsorptionDB * 100 / 8.685889638
}

//Applyies frequency to Absorption and converts into 1/m unit
func (medium *Medium) ApplyFrequencyForAbsorption(freq float64) {
	medium.AbsorptionM = medium.AbsorptionM * freq / 1.0e06
}
