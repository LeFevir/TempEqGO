// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package main

import (
	"./acs"
	"flag"
	// "log"
	"os"
)

const (
	DEF_MEDIUM_FILE = "medium.txt"
	// DEF_GRID_CALC_FILE = "grid_calc.txt"
	DEF_FIELD_FILE = "int_field.gob"

	INTEREST_POINT_I = 18
	INTEREST_POINT_J = 18
	INTEREST_POINT_K = 16
)

var (
	procNum, numt                                                         int
	tau                                                                   float64
	array_file, elements_file, elements_off_file, medium_file, focus_file string
	grid_calc_file, work_folder, field_file, mult_grid_file               string
	mode                                                                  string
	gendir, genbin, gengob, offel                                         bool
	oneElementCalc, normalize                                             bool
)

func main() {
	parseFlags()

	if gendir {
		acs.GenerateOutDirAndRedirectLog(work_folder)
	} else {
		acs.SetOutDirAndRedirectLog(work_folder)
	}

	// 3. Начать расчет температуры
	//calcField.NumericCalcPressureField(procNum, array, medium, numOfTris, true)

	switch mode {
	case "test":
		outDir := acs.OutDir
		// 2. Считать параметры среды
		medium := acs.NewMediumFromFile(medium_file)

		tempField := acs.GenerateTestField()
		tempField.SetInitTemp(100.0, 1.0e-03)

		acs.AddDirToOutDir(outDir, 0)
		tempField.PrintTempFieldXBinary(false, 50, 50)

		//heatSourcesField := acs.GenerateTestField()

		refTempField := acs.GenerateTestField()
		refTempField.CalcReferenceField(100.0, 1.0e-03, 1.29e-07, 1.0)

		acs.AddDirToOutDir(outDir, 1)
		refTempField.PrintTempFieldXBinary(false, 50, 50)

		tempField.CalcTempAndDose(acs.GenerateFieldAs(tempField), acs.GenerateFieldAs(tempField), medium, 0.0, 0.1, 10)

		acs.AddDirToOutDir(outDir, 2)
		tempField.PrintTempFieldXBinary(false, 50, 50)

		os.Exit(0)

	case "calc":
		medium := acs.NewMediumFromFile(medium_file)
		// 6.0e+06 = 6 MHz from Transducer
		medium.ApplyFrequencyForAbsorption(6.0e+06)

		intField := new(acs.Field)
		intField.RestoreFieldFromFile(field_file)

		tempField := acs.GenerateFieldAs(intField)
		tempField.FillFieldWithValue(37.0)

		doseField := acs.GenerateFieldAs(intField)
		doseField.FillFieldWithValue(0.0)

		//log.Println("TempField has been filled, in interest point it has", tempField.Get(50, 50, 50))

		tempField.CalcTempAndDose(doseField, intField, medium, 37.0, tau, numt)

		// log.Println("TempField has been calculated, in interest point it has", tempField.Get(50, 50, 50))

		tempField.PrintTempFieldZBinary(false, INTEREST_POINT_I, INTEREST_POINT_J)

		tempField.PrintTempFieldXBinary(false, INTEREST_POINT_J, INTEREST_POINT_K)
		tempField.PrintTempFieldYBinary(false, INTEREST_POINT_I, INTEREST_POINT_K)

		tempField.PrintTempFieldXZBinary(false, INTEREST_POINT_J)
		tempField.PrintTempFieldYZBinary(false, INTEREST_POINT_I)
		tempField.PrintTempFieldXYBinary(false, INTEREST_POINT_K)

		doseField.PrintDoseFieldXZBinary(false, INTEREST_POINT_J)
		doseField.PrintDoseFieldYZBinary(false, INTEREST_POINT_I)
		doseField.PrintDoseFieldXYBinary(false, INTEREST_POINT_K)

		// 2. Считать параметры среды
		// medium := acs.NewMediumFromFile(medium_file)

		// 1. Считать параметры поля из файла
		// presField := new(acs.Field)
		// presField.RestoreFieldFromFile(field_file)
		// prepareHeatSources

	}

	// normalize = false

	// if genbin {
	// 	presField.PrintAbsFieldBinary(normalize)
	// }

	// if gengob {
	// 	presField.DumpField()
	// }
}

func parseFlags() {
	flag.IntVar(&procNum, "proc", 0, "number of threads to use. 0 for autodetection")
	flag.StringVar(&mode, "mode", "test", "mode:\ncalc - calculate field with focus in point which has set in focus.txt\nscan-main - scan focus for main maxes amplitudes\nmultY - create series of fields with steering focus by coordinates from mult")
	flag.StringVar(&field_file, "field", DEF_FIELD_FILE, "path to pressure field gob-file")
	flag.StringVar(&medium_file, "medium", DEF_MEDIUM_FILE, "path to txt file with medium params")

	//flag.StringVar(&grid_calc_file, "gridcalc", DEF_GRID_CALC_FILE, "path to txt file with calculation grid params")

	flag.StringVar(&work_folder, "dir", "", "optional path to output folder")
	flag.BoolVar(&gendir, "gendir", false, "generate or not output folder with time stamp")
	flag.BoolVar(&genbin, "genbin", true, "generate or not binary field files")
	flag.BoolVar(&gengob, "gengob", false, "dump or not field gob file")

	flag.IntVar(&numt, "numt", 1, "numbers of time steps")
	flag.Float64Var(&tau, "tau", 0.1, "step on time in s")
	flag.Parse()
}
