# TempEqGO
Temperature and Dose Equation calculation on Go language

This program can calculate the acoustic field (Pressure or Intensity) of that type of transducers.

The parameters of transducer, medium and grid can be set in corresponding TXT files.

**Transducer's parameters TXT**
default txt name:
*array.txt*

Structure of file:
Name:					Probe1
Number of elements:		8
Element's width, m:		0.004
Element's length, m:	0.005
Frequency, Hz:			6000000.00

also parameters can be set without text titles:

Probe1
8
0.004
0.005
6000000.00

Transducer's active elements TXT
default txt name:
**array_elements.txt**
Structure of file:

It should contain only numbers of the elements, which are active, starting with 0,
divided by paragraphs

For example, all 8 elements active:
	0
	1
	2
	3
	4
	5
	6
	7

Only 1 element active:
	2

Medium's parameters TXT
default txt name:
**medium.txt**

Structure of file:
Name:						Water
Density, kg/m3:				1000.0
Speed Of Sound, m/s:		1500.0
Heat capacity, J/kg/K:		4200.0
Absorption, dB/cm/MHz:		0.0

also parameters can be set without text titles:
Water
1000.0
1500.0
4200.0
0.0


Grid's parameters TXT
default txt name:
**grid_calc.txt**

Structure of file:

!For example, for XY plane calculation

type of area. possible values: volume, x, y, z, yz, xz, xy
Area: xy
If we need to calc a plane XY, we should set Z, if we need to calc a Z axis, we should set X,Y coords
Z coordinate, m: 0.001
Number of points on X: 13
Number of points on Y: 73
Min value on X, m:	-0.003
Min value on Y, m:	-0.003
Step of the grid on X, m: 0.0005
Step of the grid on Y, m: 0.0005

also parameters can be set without text titles:

xy
0.001
13
73
-0.003
-0.003
0.0005
0.0005

The program works through the console and has following flags of executing:

-proc		"number of threads to use. 0 for autodetection"
-array 		"path to txt file with array params"
-els 		"path to txt file with array's active elements list"
-medium 	"path to txt file with medium params"
-gridcalc 	"path to txt file with calculation grid params"
-dir 		"optional path to output folder"
-gendir 	"generate or not output folder with time stamp in title"
-genbin 	"generate or not binary field files"
-gengob 	"dump or not field binary gob-file"
-field 		"optional path to field gob-file to restore the field without calculation"
-tris 		"number of triangles on one line to calculate the field numerical with Raleigh's integral"
