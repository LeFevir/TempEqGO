# -*- coding: utf-8 -*-
import matplotlib.pyplot as plt
import matplotlib
import readbin

# Read file with binary field data
X,Y,Z = readbin.read_2D(ur'd:\Educ\АКУСТИКА\.ПРОГА\PhlipsProstate\DoseField_XZ.bin')
#-----------------------------------


# Set Font
# plt.rc('font',**{'family':'serif','serif':['Times New Roman'],'size':14})
# Set figure size and dpi
plt.figure(figsize=(10, 4), dpi=100, facecolor='w', edgecolor='k')

C = plt.contour(Y,X,Z, levels=[1], colors = ['k'])
plt.clabel(C, fmt = '%2.1f', colors = 'k')

plt.gca().set_aspect('equal')
plt.ylabel('x, mm')
plt.xlabel('z, mm')

#plt.savefig("fug_XY.png")
plt.show()