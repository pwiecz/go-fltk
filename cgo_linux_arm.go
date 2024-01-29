//go:build linux && arm

package fltk

// #cgo linux,arm CXXFLAGS: -std=c++11
// #cgo linux,arm CPPFLAGS: -I${SRCDIR}/lib/linux/arm -I${SRCDIR}/include -I${SRCDIR}/include/FL/images -I${SRCDIR}/include/png -I${SRCDIR}/include/zlib -I${SRCDIR}/include/jpeg -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -D_THREAD_SAFE -D_REENTRANT
// #cgo linux,arm LDFLAGS: ${SRCDIR}/lib/linux/arm/libfltk_images.a ${SRCDIR}/lib/linux/arm/libfltk_jpeg.a ${SRCDIR}/lib/linux/arm/libfltk_png.a ${SRCDIR}/lib/linux/arm/libfltk_z.a ${SRCDIR}/lib/linux/arm/libfltk_gl.a -lGLU -lGL ${SRCDIR}/lib/linux/arm/libfltk_forms.a ${SRCDIR}/lib/linux/arm/libfltk.a -lm -lX11 -lXext -lpthread -lXinerama -lXfixes -lXcursor -lXft -lXrender -lm -lfontconfig -ldl
import "C"
