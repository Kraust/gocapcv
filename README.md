This was a quick and dirty method of capturing the screen to feed into opencv.
This seems to be much more performant than other solutions as it's exclusively written with X11 in mind.
Only supports X11 on Linux. Not tested on other platforms that may support X11.
Recycled from a different project.

Depends:
 - gocv dependencies (opencv, hdf5, vtk)
 - libx11
