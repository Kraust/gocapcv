package gocapcv

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
// void DestroyImage(XImage *ximage) {
//    ((*((ximage)->f.destroy_image))((ximage)));
// }
import "C"
import (
	"unsafe"

	"gocv.io/x/gocv"
)

var display *C.Display
var root C.Window

func init() {
	display = C.XOpenDisplay(nil)
	root = C.XDefaultRootWindow(display)
}

func CaptureScreen(x int, y int, height int, width int) (gocv.Mat, error) {

	ximg := C.XGetImage(display, root, C.int(x), C.int(y), C.uint(width), C.uint(height),
		C.AllPlanes, C.ZPixmap)
	bytes := C.GoBytes(unsafe.Pointer(ximg.data), ximg.height*ximg.width*(ximg.bits_per_pixel/8))
	res, err := gocv.NewMatFromBytes(height, width, gocv.MatTypeCV8UC4, bytes)
	C.DestroyImage(ximg)
	mss := gocv.NewMat()
	gocv.CvtColor(res, &mss, gocv.ColorBGRAToGray)
	res.Close()
	return mss, err
}
