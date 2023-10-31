// go:build windows && amd64 
 
package fltk 
 
// #cgo windows,amd64 CXXFLAGS: -std=c++11 
// #cgo windows,amd64 CPPFLAGS: -I${SRCDIR}/include -I${SRCDIR}/include/FL/images -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 
// #cgo windows,amd64 LDFLAGS: -mwindows ${SRCDIR}/lib/windows/amd64/libfltk_images.a ${SRCDIR}/lib/windows/amd64/libfltk_jpeg.a ${SRCDIR}/lib/windows/amd64/libfltk_png.a ${SRCDIR}/lib/windows/amd64/libfltk_z.a ${SRCDIR}/lib/windows/amd64/libfltk_gl.a -lglu32 -lopengl32 ${SRCDIR}/lib/windows/amd64/libfltk_forms.a ${SRCDIR}/lib/windows/amd64/libfltk.a -lgdiplus -lole32 -luuid -lcomctl32 -lws2_32 
import "C" 
