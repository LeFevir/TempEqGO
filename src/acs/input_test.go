// Copyright 2012. Sergey Ilyin. All rights reserved.
// Lab 366, Acoustic Department, Faculty of Physics
// Lomonosov Moscow State University

package acs

import (
	"testing"
)

func TestParseParamFile(t *testing.T) {
	s := parseParamFile("test_files\\water_medium.txt")
	if len(s) != 5 {
		t.Log("Wrong number of values in medium file!")
		t.Fail()
	}

	s = parseParamFile("test_files\\hand_array.txt")
	if len(s) == 0 {
		t.Log("No values for array!")
		t.Fail()
	}

	s = parseParamFile("test_files\\hand_array_BOM.txt")
	if len(s) == 0 {
		t.Log("No values for array!")
		t.Fail()
	}

	s = parseParamFile("test_files\\hand_array_elements.txt")
	if len(s) == 0 {
		t.Log("No values for array elements!")
		t.Fail()
	}

	s = parseParamFile("test_files\\grid_calc.txt")
	if len(s) == 0 {
		t.Log("No values for grid calc!")
		t.Fail()
	}

	s = parseParamFile("test_files\\grid_scan.txt")
	if len(s) == 0 {
		t.Log("No values for grid scan!")
		t.Fail()
	}
}

func TestNewArrayFromFile(t *testing.T) {
	s := NewArrayFromFile("test_files\\array.txt")
	if s.Name != "ProstateProbe1" {
		t.Log("A mistake at reading array name!")
		t.Fail()
	}
	if s.ElementWidth != 0.004 {
		t.Log("A mistake at reading element width!")
		t.Fail()
	}
	if s.ElementLength != 0.005 {
		t.Log("A mistake at reading element length!")
		t.Fail()
	}
	if s.Frequency != 6000000.0 {
		t.Log("A mistake at reading array frequency!")
		t.Fail()
	}
}

func TestAddElementsFromFile(t *testing.T) {
	array := NewArrayFromFile("test_files\\hand_array.txt")
	array.AddElementsFromFile("test_files\\hand_array_elements.txt")
	if array.NumberOfElements() != 256 {
		t.Log("A mistake at reading array elements!")
		t.Fail()
	}
}

func TestNewMediumFromFile(t *testing.T) {
	s := NewMediumFromFile("test_files\\water_medium.txt")
	if s.Name != "Water" {
		t.Log("A mistake at reading medium name!")
		t.Fail()
	}
	if s.SpeedOfSound != 1500.0 {
		t.Log("A mistake at reading speed of sound!")
		t.Fail()
	}
	if s.Density != 1000.0 {
		t.Log("A mistake at reading density!")
		t.Fail()
	}
}
func TestSetFocusFromFile(t *testing.T) {
	array := NewArrayFromFile("test_files\\hand_array.txt")
	array.AddElementsFromFile("test_files\\hand_array_elements.txt")
	array.SetFocusFromFile("test_files\\focus.txt")
}

func TestNewYZGridFromFile(t *testing.T) {
	field := NewGridFromFile("test_files\\grid_yz.txt")
	if field.Nx != 1 || field.Ny != 51 || field.Nz != 121 {
		t.Log("A mistake at reading grid number of points!")
		t.Fail()
	}
	if field.MinX != 0.0 || field.MinY != -0.025 || field.MinZ != 0.070 {
		t.Log("A mistake at reading grid Bottom Border!")
		t.Fail()
	}
	if field.Dx != 1.0 || field.Dy != 0.001 || field.Dz != 0.001 {
		t.Log("A mistake at reading grid step!")
		t.Fail()
	}
}

func TestNewXZGridFromFile(t *testing.T) {
	field := NewGridFromFile("test_files\\grid_xz.txt")
	if field.Nx != 51 || field.Ny != 1 || field.Nz != 121 {
		t.Log("A mistake at reading grid number of points!")
		t.Fail()
	}
	if field.MinX != -0.025 || field.MinY != 0.0 || field.MinZ != 0.070 {
		t.Log("A mistake at reading grid Bottom Border!")
		t.Fail()
	}
	if field.Dx != 0.001 || field.Dy != 1.0 || field.Dz != 0.001 {
		t.Log("A mistake at reading grid step!")
		t.Fail()
	}
}

func TestNewXYGridFromFile(t *testing.T) {
	field := NewGridFromFile("test_files\\grid_xy.txt")
	if field.Nx != 51 || field.Ny != 51 || field.Nz != 1 {
		t.Log("A mistake at reading grid number of points!")
		t.Fail()
	}
	if field.MinX != -0.025 || field.MinY != -0.025 || field.MinZ != 0.130 {
		t.Log("A mistake at reading grid Bottom Border!")
		t.Fail()
	}
	if field.Dx != 0.001 || field.Dy != 0.001 || field.Dz != 1.0 {
		t.Log("A mistake at reading grid step!")
		t.Fail()
	}
}

func TestNewZGridFromFile(t *testing.T) {
	field := NewGridFromFile("test_files\\grid_z.txt")
	if field.Nx != 1 || field.Ny != 1 || field.Nz != 121 {
		t.Log("A mistake at reading grid number of points!")
		t.Fail()
	}
	if field.MinX != 0.0 || field.MinY != 0.0 || field.MinZ != 0.070 {
		t.Log("A mistake at reading grid Bottom Border!")
		t.Fail()
	}
	if field.Dx != 1.0 || field.Dy != 1.0 || field.Dz != 0.001 {
		t.Log("A mistake at reading grid step!")
		t.Fail()
	}
}

func TestNewYGridFromFile(t *testing.T) {
	field := NewGridFromFile("test_files\\grid_y.txt")
	if field.Nx != 1 || field.Ny != 51 || field.Nz != 1 {
		t.Log("A mistake at reading grid number of points!")
		t.Fail()
	}
	if field.MinX != 0.0 || field.MinY != -0.025 || field.MinZ != 0.130 {
		t.Log("A mistake at reading grid Bottom Border!")
		t.Fail()
	}
	if field.Dx != 1.0 || field.Dy != 0.001 || field.Dz != 1.0 {
		t.Log("A mistake at reading grid step!")
		t.Fail()
	}
}

func TestNewXGridFromFile(t *testing.T) {
	field := NewGridFromFile("test_files\\grid_x.txt")
	if field.Nx != 51 || field.Ny != 1 || field.Nz != 1 {
		t.Log("A mistake at reading grid number of points!")
		t.Fail()
	}
	if field.MinX != -0.025 || field.MinY != 0.0 || field.MinZ != 0.130 {
		t.Log("A mistake at reading grid Bottom Border!")
		t.Fail()
	}
	if field.Dx != 0.001 || field.Dy != 1.0 || field.Dz != 1.0 {
		t.Log("A mistake at reading grid step!")
		t.Fail()
	}
}
