package main

// #cgo LDFLAGS: -L /usr/local/lib -lgeotiff
// #cgo CFLAGS: -I /usr/local/include
// #include <geotiff.h>
// #include <xtiffio.h>
// #include <geo_normalize.h>
import "C"
import (
	"errors"
	"unsafe"
)

func GeoTIFFProj4Def(file string) (string, error) {
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
