package main

// #cgo LDFLAGS: -lproj
// #include <proj_api.h>
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

type Proj struct {
	p _Ctype_projPJ
}

func NewProj(def string) (*Proj, error) {
	cdef := C.CString(def)
	defer C.free(unsafe.Pointer(cdef))

	proj := &Proj{p: C.pj_init_plus(cdef)}

	if proj.p == nil {
		return nil, errors.New(projError())
	}

	// Set finalizer so the garbace collector
	// can release the object properly
	runtime.SetFinalizer(proj, (*Proj).Close)

	return proj, nil
}

func (proj *Proj) Close() {
	if proj.p != nil {
		C.pj_free(proj.p)
	}
}

func projError() string {
	if C.pj_errno == 0 {
		return ""
	}
	
	err := C.pj_strerrno(C.pj_errno)
	return C.GoString(err)
}
