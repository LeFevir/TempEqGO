// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var OutDir string

func AddDirToOutDir(oldDir string, newDirInt int) {
	newDir := strconv.Itoa(newDirInt)
	OutDir = filepath.Join(oldDir, newDir)
	error := os.Mkdir(OutDir, os.ModeDir)
	if error != nil {
		fmt.Printf("An error occurred on creating output folder %s\n", error)
		log.Printf("An error occurred on creating output folder %s\n", error)
		return
	}
	log.Println("Успешно создана папка для вывода", OutDir)
	return
}

//Generates output dir for current calculations titled with current date and time
func GenerateOutDirAndRedirectLog(workdir string) {
	workdir, _ = filepath.Abs(workdir)
	OutDir = filepath.Join(workdir, time.Now().Format("2006-01-02_15-04-05"))
	error := os.Mkdir(OutDir, os.ModeDir)
	if error != nil {
		fmt.Printf("An error occurred on creating output folder %s\n", error)
		log.Printf("An error occurred on creating output folder %s\n", error)
		return
	}
	redirectLog()
	log.Println("Успешно создана папка для вывода", OutDir)
}

func SetOutDirAndRedirectLog(workdir string) {
	OutDir, _ = filepath.Abs(workdir)
	redirectLog()
	log.Println("Задана папка для вывода", OutDir)
}

func redirectLog() {
	file := createTxtFileInOutDir("log")
	log.SetOutput(file)
	fmt.Println("Logging into", file.Name())
}

func PrintStatus(percents int, startTime time.Time) {
	fmt.Printf("Calculated: %d%%. ", percents)
	log.Printf("Рассчитано: %d%%\n", percents)
	if percents != 0 {
		d := time.Since(startTime)
		fmt.Println("Remaining:", time.Duration(100*int64(d)/int64(percents)-int64(d)))
	} else {
		fmt.Println("Can't define remaining time")
	}
}

func createFile(path string) (file *os.File) {
	file, erOpen := os.Create(path)
	if erOpen != nil {
		fmt.Printf("An error occurred on creating the files for output\n")
		log.Printf("An error occurred on creating the files for output\n")
	}
	return
}

func createGobFileInOutDir(name string) *os.File {
	path := filepath.Join(OutDir, name+".gob")
	return createFile(path)
}

func createBinFileInOutDir(name string) *os.File {
	path := filepath.Join(OutDir, name+".bin")
	return createFile(path)
}

func createTxtFileInOutDir(name string) *os.File {
	path := filepath.Join(OutDir, name+".txt")
	return createFile(path)
}

//This function saves (dumps) Field into file field.gob
func (field *Field) DumpField() {
	file := createGobFileInOutDir("field")
	defer file.Close()

	enc := gob.NewEncoder(file)
	erEnc := enc.Encode(field)
	if erEnc != nil {
		fmt.Printf("Error in dumping of Field %s\n", erEnc)
		log.Printf("Error in dumping of Field %s\n", erEnc)
		return
	}
	fmt.Println("Field dumping has acomplished")
	log.Println("Успешно закончено сохранение поля в", file.Name())
}

//This function restores Field from defined GOB-file
func (field *Field) RestoreFieldFromFile(filepath string) {
	file, erOpen := os.Open(filepath)
	if erOpen != nil {
		fmt.Printf("An error occurred on opening the file %s\n", erOpen)
		log.Printf("An error occurred on opening the file %s\n", erOpen)
		return
	}
	defer file.Close()

	erDec := gob.NewDecoder(file).Decode(&field)
	if erDec != nil {
		fmt.Printf("Error in restoring Field from file %s\n", erDec)
		return
	}

	fmt.Println("Restoring of field has finished")
	log.Println("Восстановление поля из файла закончено")
}

//This function prints planes XY, YZ, XZ of the FieldComplex into binary files (LittleEndian)
func (field *Field) PrintTempFieldBinary(normalize bool) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	switch field.Type {
	case "yz":
		field.PrintTempFieldYZBinary(normalize, 0)
	case "xz":
		field.PrintTempFieldXZBinary(normalize, 0)
	case "xy":
		field.PrintTempFieldXYBinary(normalize, 0)
	case "z":
		field.PrintTempFieldZBinary(normalize, 0, 0)
	case "y":
		field.PrintTempFieldYBinary(normalize, 0, 0)
	case "x":
		field.PrintTempFieldXBinary(normalize, 0, 0)
	case "volume":
		field.PrintTempField2DBinary(normalize, field.Nx/2, field.Ny/2, field.Nz/2)
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

//This function prints planes XY, YZ, XZ of the Field into binary files (LittleEndian)
func (field *Field) PrintTempField2DBinary(normalize bool, pointOnX int, pointOnY int, pointOnZ int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	field.PrintTempFieldXZBinary(normalize, pointOnY)
	field.PrintTempFieldYZBinary(normalize, pointOnX)
	field.PrintTempFieldXYBinary(normalize, pointOnZ)

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintTempFieldYZBinary(normalize bool, pointOnX int) {
	fmt.Println("Binary FieldComplex output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("TempField_YZ")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Ny, field.Nz)

	i := pointOnX
	for j, yj := range field.Y {
		for k, zk := range field.Z {
			writeValuesBin(file, yj, zk, field.Get(i, j, k)*norm)
		}
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintTempFieldXZBinary(normalize bool, pointOnY int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("TempField_XZ")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Nx, field.Nz)

	j := pointOnY
	for i, xi := range field.X {
		for k, zk := range field.Z {
			writeValuesBin(file, xi, zk, field.Get(i, j, k)*norm)
		}
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintTempFieldXYBinary(normalize bool, pointOnZ int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("TempField_XY")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Nx, field.Ny)

	k := pointOnZ
	for i, xi := range field.X {
		for j, yj := range field.Y {
			writeValuesBin(file, xi, yj, field.Get(i, j, k)*norm)
		}
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintTempFieldZBinary(normalize bool, pointOnX, pointOnY int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("TempField_Z")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Nz)

	i := pointOnX
	j := pointOnY
	for k, zk := range field.Z {
		writeValuesBin(file, zk, field.Get(i, j, k)*norm)
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintTempFieldYBinary(normalize bool, pointOnX, pointOnZ int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("TempField_Y")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Ny)

	i := pointOnX
	k := pointOnZ
	for j, yj := range field.Y {
		writeValuesBin(file, yj, field.Get(i, j, k)*norm)
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintTempFieldXBinary(normalize bool, pointOnY, pointOnZ int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("TempField_X")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Nx)

	j := pointOnY
	k := pointOnZ
	for i, xi := range field.X {
		writeValuesBin(file, xi, field.Get(i, j, k)*norm)
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func writeSizesBin(file *os.File, values ...int) {
	for _, v := range values {
		binary.Write(file, binary.LittleEndian, int64(v))
	}
}

func writeValuesBin(file *os.File, values ...float64) {
	for _, v := range values {
		binary.Write(file, binary.LittleEndian, v)
	}
}

func (field *Field) PrintDoseFieldYZBinary(normalize bool, pointOnX int) {
	fmt.Println("Binary FieldComplex output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("DoseField_YZ")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Ny, field.Nz)

	i := pointOnX
	for j, yj := range field.Y {
		for k, zk := range field.Z {
			writeValuesBin(file, yj, zk, field.Get(i, j, k)*norm)
		}
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintDoseFieldXZBinary(normalize bool, pointOnY int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("DoseField_XZ")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Nx, field.Nz)

	j := pointOnY
	for i, xi := range field.X {
		for k, zk := range field.Z {
			writeValuesBin(file, xi, zk, field.Get(i, j, k)*norm)
		}
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}

func (field *Field) PrintDoseFieldXYBinary(normalize bool, pointOnZ int) {
	fmt.Println("Binary Field output has started")
	log.Println("Начат вывод давления в бинарные файлы")

	file := createBinFileInOutDir("DoseField_XY")
	defer file.Close()

	//norm - норма
	norm := 1.0
	if normalize {
		norm = 1.0 / (field.MaxValue())
	}

	writeSizesBin(file, field.Nx, field.Ny)

	k := pointOnZ
	for i, xi := range field.X {
		for j, yj := range field.Y {
			writeValuesBin(file, xi, yj, field.Get(i, j, k)*norm)
		}
	}

	fmt.Println("Binary Field output has finished")
	log.Println("Вывод давления в бинарные файлы завершен")
}
