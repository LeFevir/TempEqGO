# -*- coding: utf-8 -*-
import matplotlib.pyplot as plt
import matplotlib
import readbin

# Read file with binary field data
X,Y,Z = readbin.read_2D(ur'd:\Educ\АКУСТИКА\.ПРОГА\PhlipsProstate\TempField_YZ.bin')
#-----------------------------------


# Set Font
# plt.rc('font',**{'family':'serif','serif':['Times New Roman'],'size':14})
# Set figure size and dpi
plt.figure(figsize=(5, 5), dpi=100, facecolor='w', edgecolor='k')

plt.contourf(Y,X,Z,100,cmap=plt.cm.jet)
plt.colorbar(orientation='horizontal', format='%.0f')


plt.gca().set_aspect('equal')
plt.ylabel('y, mm')
plt.xlabel('z, mm')

#plt.savefig("fug_XY.png")
plt.show()