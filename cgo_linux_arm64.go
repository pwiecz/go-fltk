//go:build linux && arm64

package fltk

// #cgo linux,arm64 CXXFLAGS: -std=c++11
// #cgo linux,arm64 CPPFLAGS: -I${SRCDIR}/lib/linux/arm64 -I${SRCDIR}/include -I${SRCDIR}/include/FL/images -I${SRCDIR}/include/png -I${SRCDIR}/include/zlib -I${SRCDIR}/include/jpeg -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -D_THREAD_SAFE -D_REENTRANT
// #cgo linux,arm64 LDFLAGS: ${SRCDIR}/lib/linux/arm64/libfltk_images.a ${SRCDIR}/lib/linux/arm64/libfltk_jpeg.a ${SRCDIR}/lib/linux/arm64/libfltk_png.a ${SRCDIR}/lib/linux/arm64/libfltk_z.a ${SRCDIR}/lib/linux/arm64/libfltk_gl.a ${SRCDIR}/lib/linux/arm64/libfltk_forms.a ${SRCDIR}/lib/linux/arm64/libfltk.a -lm -lX11 -lpthread -lm -ldl
import "C"
