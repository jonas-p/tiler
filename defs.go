package main

// #cgo LDFLAGS: -lgeotiff -ltiff
// #include <stdlib.h>
// #include <geotiff.h>
// #include <tiffio.h>
// #include <xtiffio.h>
// #include <geo_normalize.h>
//
// int TIFFGetIntField(TIFF *tif, ttag_t tag) {
//     int v; TIFFGetField(tif, tag, &v);
//     return v;
// }
import "C"
import (
	"errors"
	"unsafe"
)

func GeoTIFFOffsetToCoords(gtif *C.GTIF, x, y int) (float64, float64, error) {
	cx := C.double(x)
	cy := C.double(y)

	if C.GTIFImageToPCS(gtif, &cx, &cy) == 0 {
		return 0, 0, errors.New("Could not transform offset")
	}

	return float64(cx), float64(cy), nil
}

func GeoTIFFRepresentation(file string) (*ImageRepresentation, error) {
	C.TIFFSetErrorHandler(nil)
	C.TIFFSetWarningHandler(nil)

	fstr := C.CString(file)
	defer C.free(unsafe.Pointer(fstr))

	rstr := C.CString("r")
	defer C.free(unsafe.Pointer(rstr))

	tif := C.XTIFFOpen(fstr, rstr)
	if tif == nil {
		return nil, errors.New("Could not open file")
	}
	defer C.XTIFFClose(tif)

	gtif := C.GTIFNew(unsafe.Pointer(tif))
	if gtif == nil {
		return nil, errors.New("Could not open file as GeoTIFF")
	}
	defer C.GTIFFree(gtif)

	defn := new(C.GTIFDefn)

	if C.GTIFGetDefn(gtif, defn) == 0 {
		return nil, errors.New("Unable to get coordinate system defintion")
	}

	// Bounding box
	xsize := int(C.TIFFGetIntField(tif, C.TIFFTAG_IMAGEWIDTH))
	ysize := int(C.TIFFGetIntField(tif, C.TIFFTAG_IMAGELENGTH))

	// Upper left
	ulx, uly, err := GeoTIFFOffsetToCoords(gtif, 0, 0)
	if err != nil {
		return nil, errors.New("Unable to get bounding box")
	}

	// Lower left
	llx, lly, _ := GeoTIFFOffsetToCoords(gtif, 0, ysize)
	// Upper right
	urx, ury, _ := GeoTIFFOffsetToCoords(gtif, xsize, 0)
	// Lower right
	lrx, lry, _ := GeoTIFFOffsetToCoords(gtif, xsize, ysize)

	rep := ImageRepresentation{
		BoundingBox: BoundingBox{
			UpperLeft:  Point{ulx, uly},
			LowerLeft:  Point{llx, lly},
			UpperRight: Point{urx, ury},
			LowerRight: Point{lrx, lry},
		},
		Proj4:    C.GoString(C.GTIFGetProj4Defn(defn)),
		EpsgCode: int(defn.Projection),
	}

	return &rep, nil
}
