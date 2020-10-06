# -*- coding: utf-8 -*-
import matplotlib.pyplot as plt
import matplotlib
import readbin

# Read file with binary field data
X2,Y2,Z2 = readbin.read_2D(ur'd:\Educ\АКУСТИКА\.ПРОГА\PhlipsProstate\DoseField_XZ.bin')
X1,Y1,Z1 = readbin.read_2D(ur'd:\Educ\АКУСТИКА\.ПРОГА\PhlipsProstate\TempField_XZ.bin')
#-----------------------------------


# Set Font
# plt.rc('font',**{'family':'serif','serif':['Times New Roman'],'size':14})
# Set figure size and dpi
plt.figure(figsize=(11, 5), dpi=100, facecolor='w', edgecolor='k')

plt.contourf(Y1,X1,Z1,100,cmap=plt.cm.jet)
plt.colorbar(orientation='horizontal', format='%.0f')

C = plt.contour(Y2,X2,Z2, levels=[1], colors = ['w'], linewidths = [3])
# plt.clabel(C, fmt = '%2.1f', colors = 'w')

plt.gca().set_aspect('equal')
plt.ylabel('x, mm')
plt.xlabel('z, mm')

#plt.savefig("fug_XY.png")
plt.show()