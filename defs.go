package main

// #cgo LDFLAGS: -lgeotiff -ltiff
// #include <geotiff.h>
// #include <tiffio.h>
// #include <xtiffio.h>
// #include <geo_normalize.h>
import "C"
import (
	"errors"
	"unsafe"
)

func GeoTIFFProj4Def(file string) (string, error) {
 	C.TIFFSetErrorHandler(nil)
	C.TIFFSetWarningHandler(nil)
	
	tif := C.XTIFFOpen(C.CString(file), C.CString("r"))
	if tif == nil {
		return "", errors.New("Could not open file")
	}
	defer C.XTIFFClose(tif)

	gtif := C.GTIFNew(unsafe.Pointer(tif))
	if gtif == nil {
		return "", errors.New("Could not open file as GeoTIFF")
	}
	defer C.GTIFFree(gtif)

	defn := new(C.GTIFDefn)

	if C.GTIFGetDefn(gtif, defn) == 0 {
		return "", errors.New("Unable to get coordinate system defintion")
	}

	proj4 := C.GoString(C.GTIFGetProj4Defn(defn))
	return proj4, nil
}
