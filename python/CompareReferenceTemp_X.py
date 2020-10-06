# -*- coding: utf-8 -*-
import matplotlib.pyplot as plt
import readbin

# Read file with binary field data
X, Z0 = readbin.read_1D(ur'd:\Educ\АКУСТИКА\.ПРОГА\PhlipsProstate\0\TempField_X.bin')
X, Z1 = readbin.read_1D(ur'd:\Educ\АКУСТИКА\.ПРОГА\PhlipsProstate\1\TempField_X.bin')
X, Z2 = readbin.read_1D(ur'd:\Educ\АКУСТИКА\.ПРОГА\PhlipsProstate\2\TempField_X.bin')


# Set Font
plt.rc('font',**{'family':'serif','serif':['Times New Roman'],'size':16})
# Set figure size and dpi
plt.figure(facecolor='w', edgecolor='k')
# plt.figure(figsize=(5, 5), dpi=100, facecolor='w', edgecolor='k')

plt.plot(X,Z0,'k',linewidth = 2,label = 'Initial Temp distr')
plt.plot(X,Z1,'go',linewidth = 2,label = 'Reference temp after 1 s')
plt.plot(X,Z2,'r',linewidth = 2,label = 'Simulated temp after 1 s')

#plt.gca().set_aspect('equal')
plt.xlabel('x, $\it{mm}$')
plt.ylabel(r'T,$^{\circ}\it{C}$')
plt.title(u'Comparison of simulation of temp diffusion and reference distr',fontsize=16)
plt.legend(loc=2,prop={'size':16})

#plt.savefig("fug_XY.png")
plt.show()