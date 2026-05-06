//go:build darwin && amd64

package fltk

// #cgo darwin,amd64 CXXFLAGS: -std=c++11
// #cgo darwin,amd64 CPPFLAGS: -I${SRCDIR}/lib/darwin/amd64 -I${SRCDIR}/include -I${SRCDIR}/include/FL/images -I${SRCDIR}/include/png -I${SRCDIR}/include/zlib -I${SRCDIR}/include/jpeg -isysroot /Library/Developer/CommandLineTools/SDKs/MacOSX.sdk -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -D_THREAD_SAFE -D_REENTRANT
// #cgo darwin,amd64 LDFLAGS: ${SRCDIR}/lib/darwin/amd64/libfltk_images.a ${SRCDIR}/lib/darwin/amd64/libfltk_jpeg.a ${SRCDIR}/lib/darwin/amd64/libfltk_png.a ${SRCDIR}/lib/darwin/amd64/libfltk_z.a ${SRCDIR}/lib/darwin/amd64/libfltk_gl.a -framework OpenGL ${SRCDIR}/lib/darwin/amd64/libfltk_forms.a ${SRCDIR}/lib/darwin/amd64/libfltk.a -lm -lpthread -framework Cocoa
import "C"
