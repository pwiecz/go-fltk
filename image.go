package fltk

/*
#include <stdlib.h>
#include "image.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type image struct {
	ptr *C.Fl_Image
}

type Image interface {
	getImage() *image
}

func (i *image) getImage() *image { return i }

func initImage(i Image, p unsafe.Pointer) {
	i.getImage().ptr = (*C.Fl_Image)(p)
}

func (i *image) Destroy() {
	if i.ptr != nil {
		C.go_fltk_image_delete(i.ptr)
	}
	i.ptr = nil
}

func (i *image) W() int {
	return int(C.go_fltk_image_w(i.ptr))
}
func (i *image) H() int {
	return int(C.go_fltk_image_h(i.ptr))
}
func (i *image) Count() int {
	return int(C.go_fltk_image_count(i.ptr))
}
func (i *image) Scale(width int, height int, proportional bool, can_expand bool) {
	prop := 0
	expand := 0
	if proportional {
		prop = 1
	}
	if can_expand {
		expand = 1
	}
	C.go_fltk_image_scale(i.ptr, C.int(width), C.int(height), C.int(prop), C.int(expand))
}
func (i *image) fail() int {
	return int(C.go_fltk_image_fail(i.ptr))
}
func (i *image) DataW() int {
	return int(C.go_fltk_image_data_w(i.ptr))
}
func (i *image) DataH() int {
	return int(C.go_fltk_image_data_h(i.ptr))
}
func (i *image) D() int {
	return int(C.go_fltk_image_d(i.ptr))
}
func (i *image) Ld() int {
	return int(C.go_fltk_image_ld(i.ptr))
}
func (i *image) Inactive() {
	C.go_fltk_image_inactive(i.ptr)
}

func image_error(val int) error {
	if val == -1 {
		return errors.New("No Image was found")
	} else if val == -2 {
		return errors.New("File access error")
	} else if val == -3 {
		return errors.New("Image format error")
	} else {
		return nil
	}
}

type SvgImage struct {
	image
}

func SvgImageLoad(path string) (*SvgImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &SvgImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_svg_image_load(fileStr)))
	return img, image_error(img.fail())
}

func SvgImageString(str string) (*SvgImage, error) {
	imagestr := C.CString(str)
	defer C.free(unsafe.Pointer(imagestr))
	img := &SvgImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_svg_image_data(imagestr)))
	return img, image_error(img.fail())
}

type PngImage struct {
	image
}

func PngImageLoad(path string) (*PngImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &PngImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_png_image_load(fileStr)))
	return img, image_error(img.fail())
}

type JpegImage struct {
	image
}

func JpegImageLoad(path string) (*JpegImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &JpegImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_jpg_image_load(fileStr)))
	return img, image_error(img.fail())
}

type RgbImage struct {
	image
}

func RgbImageData(data []uint8, w, h, depth int) (*RgbImage, error) {
	buf := (*C.uchar)(unsafe.Pointer(&data[0]))

	img := &RgbImage{}

	initImage(img, unsafe.Pointer(C.go_fltk_rgb_image_data(buf, C.int(w), C.int(h), C.int(depth), C.int(0))))
	return img, image_error(img.fail())
}

type SharedImage struct {
	image
}

func SharedImageLoad(path string) (*SharedImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &SharedImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_shared_image_load(fileStr)))
	return img, image_error(img.fail())
}
