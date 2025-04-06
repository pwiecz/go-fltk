//go:build linux && amd64

package fltk

// #cgo linux,amd64 CXXFLAGS: -std=c++11
// #cgo linux,amd64 CPPFLAGS: -I${SRCDIR}/include/linux/amd64 -I${SRCDIR}/include/linux/amd64/FL/images -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -D_THREAD_SAFE -D_REENTRANT
// #cgo linux,amd64 LDFLAGS: ${SRCDIR}/lib/linux/amd64/libfltk_images.a ${SRCDIR}/lib/linux/amd64/libfltk_jpeg.a ${SRCDIR}/lib/linux/amd64/libfltk_png.a ${SRCDIR}/lib/linux/amd64/libfltk_z.a ${SRCDIR}/lib/linux/amd64/libfltk_gl.a -L/usr/lib/x86_64-linux-gnu -lwayland-egl -lwayland-client -lEGL -lGLU -lGL ${SRCDIR}/lib/linux/amd64/libfltk_forms.a ${SRCDIR}/lib/linux/amd64/libfltk.a -lm -lpthread -lXinerama -lXfixes -lXcursor -L/usr/lib/x86_64-linux-gnu -lpangoxft-1.0 -lpangoft2-1.0 -lpango-1.0 -lgobject-2.0 -lglib-2.0 -lharfbuzz -lfontconfig -lfreetype -lXft -lpangocairo-1.0 -lcairo -lgtk-3 -lgdk-3 -lgio-2.0 -lX11 -lXrender -lwayland-cursor -lwayland-client -lxkbcommon -ldbus-1 -ldl
import "C"
