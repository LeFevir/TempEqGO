# -*- coding: utf-8 -*-
"""
Module for reading binary files with field information

"""

import struct
import numpy.numarray as narray

def read_1D(path):
	# Read file with binary field data
	fid = open(path, 'rb')
	isize,  = struct.unpack('<q',fid.read(8))			# 16  - because size of variable is 8 bytes

	X = narray.zeros((isize))
	Z = narray.zeros((isize))

	for i in xrange(0,isize):
	        X[i], Z[i]=struct.unpack('<dd',fid.read(16))

	fid.close()

	X = X*1000.0 # To make it in mm

	return X, Z
	#-----------------------------------

def read_2D(path):
	# Read file with binary field data
	fid = open(path, 'rb')
	isize, jsize = struct.unpack('<qq',fid.read(16))			# 8  - because size of variable is 8 bytes

	X = narray.zeros((isize,jsize))
	Y = narray.zeros((isize,jsize))
	Z = narray.zeros((isize,jsize))


	for i in xrange(0,isize):
		for j in xrange(0,jsize):
			X[i,j], Y[i,j], Z[i,j]=struct.unpack('<ddd',fid.read(24))

	fid.close()

	X = X*1000.0 # To make it in mm
	Y = Y*1000.0

	return X, Y, Z
	#-----------------------------------

if __name__ == '__main__':
	pass