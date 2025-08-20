package fltk

/*
#include <stdlib.h>
#include "image.h"
*/
import "C"
import (
	"errors"
	"fmt"
	goimage "image"
	"unsafe"
)

type image struct {
	iPtr *C.Fl_Image
}

type Image interface {
	getImage() *image
}

var ErrImageDestroyed = errors.New("image is destroyed")

func (i *image) getImage() *image { return i }
func (i *image) ptr() *C.Fl_Image {
	if i.iPtr == nil {
		panic(ErrImageDestroyed)
	}
	return i.iPtr
}

func initImage(i Image, p unsafe.Pointer) {
	i.getImage().iPtr = (*C.Fl_Image)(p)
}

func (i *image) Destroy() {
	C.go_fltk_image_delete(i.ptr())
	i.iPtr = nil
}

func (i *image) Draw(x, y, w, h int) {
	C.go_fltk_image_draw(i.ptr(), C.int(x), C.int(y), C.int(w), C.int(h))
}

func (i *image) W() int {
	return int(C.go_fltk_image_w(i.ptr()))
}
func (i *image) H() int {
	return int(C.go_fltk_image_h(i.ptr()))
}
func (i *image) Count() int {
	return int(C.go_fltk_image_count(i.ptr()))
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
	C.go_fltk_image_scale(i.ptr(), C.int(width), C.int(height), C.int(prop), C.int(expand))
}
func (i *image) fail() C.int {
	return C.go_fltk_image_fail(i.ptr())
}
func (i *image) DataW() int {
	return int(C.go_fltk_image_data_w(i.ptr()))
}
func (i *image) DataH() int {
	return int(C.go_fltk_image_data_h(i.ptr()))
}
func (i *image) D() int {
	return int(C.go_fltk_image_d(i.ptr()))
}
func (i *image) Ld() int {
	return int(C.go_fltk_image_ld(i.ptr()))
}
func (i *image) Inactive() {
	C.go_fltk_image_inactive(i.ptr())
}

// The ColorAverage() method averages the colors in the image with the provided FLTK color value.
// The second argument specifies the amount of the original image to combine with the color,
// so a value of 1.0 results in no color blend, and a value of 0.0 results in a constant image of the specified color.

func (i *image) ColorAverage(color Color, blend float32) {
	if blend > 1.0 {
		blend = 1.0
	} else if blend < 0 {
		blend = 0
	}

	C.go_fltk_image_color_average(i.ptr(), C.uint(color), C.float(blend))
}

var ErrNoImage = errors.New("no image was found")
var ErrImageFileAccess = errors.New("image file access error")
var ErrImageDecodingFailed = errors.New("image decoding failed")
var ErrImageMemoryAccess = errors.New("invalid memory access by image decoder")

func image_error(val C.int) error {
	if val == 0 {
		return nil
	} else if val == C.go_Fl_Image_ERR_NO_IMAGE {
		return ErrNoImage
	} else if val == C.go_Fl_Image_ERR_FILE_ACCESS {
		return ErrImageFileAccess
	} else if val == C.go_Fl_Image_ERR_FORMAT {
		return ErrImageDecodingFailed
	} else if val == C.go_Fl_Image_ERR_MEMORY_ACCESS {
		return ErrImageMemoryAccess
	} else {
		return fmt.Errorf("unknown image error: %d", int(val))
	}
}

type RgbImage struct {
	image
}

var ErrImageDataTooShort = errors.New("image data too short")

func NewRgbImage(data []uint8, w, h, depth int) (*RgbImage, error) {
	if len(data) < w*h*depth {
		return nil, ErrImageDataTooShort
	}
	img := &RgbImage{}
	cData := (*C.uchar)(unsafe.Pointer(&data[0]))
	initImage(img, unsafe.Pointer(C.go_fltk_rgb_image_data(cData, C.int(len(data)), C.int(w), C.int(h), C.int(depth), 0)))
	if err := image_error(img.fail()); err != nil {
		img.Destroy()
		return nil, err
	}
	return img, nil
}

func NewRgbImageFromImage(image goimage.Image) (*RgbImage, error) {
	rgbImage := &RgbImage{}
	var w, h, stride, depth int
	var data []uint8
	var dataLen int
	var cData *C.uchar
	if grayImage, ok := image.(*goimage.Gray); ok {
		cData = (*C.uchar)(unsafe.Pointer(&grayImage.Pix[0]))
		dataLen = len(grayImage.Pix)
		w, h, stride = grayImage.Rect.Dx(), grayImage.Rect.Dy(), grayImage.Stride
		depth = 1
	} else if rgbaImage, ok := image.(*goimage.RGBA); ok {
		cData = (*C.uchar)(unsafe.Pointer(&rgbaImage.Pix[0]))
		dataLen = len(rgbaImage.Pix)
		w, h, stride = rgbaImage.Rect.Dx(), rgbaImage.Rect.Dy(), rgbaImage.Stride
		depth = 4
	} else {
		w, h = image.Bounds().Dx(), image.Bounds().Dy()
		depth = 4
		data = make([]uint8, 0, w*h*4)
		for dy := 0; dy < h; dy++ {
			y := image.Bounds().Min.Y + dy
			for dx := 0; dx < w; dx++ {
				x := image.Bounds().Min.X + dx
				r, g, b, a := image.At(x, y).RGBA()
				r8, g8, b8, a8 := uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)
				data = append(data, r8, g8, b8, a8)
			}
		}
		cData = (*C.uchar)(unsafe.Pointer(&data[0]))
		dataLen = len(data)
	}
	initImage(rgbImage, unsafe.Pointer(C.go_fltk_rgb_image_data(cData, C.int(dataLen), C.int(w), C.int(h), C.int(depth), C.int(stride))))
	if err := image_error(rgbImage.fail()); err != nil {
		rgbImage.Destroy()
		return nil, err
	}
	return rgbImage, nil
}

type SvgImage struct {
	RgbImage
}

func NewSvgImageLoad(path string) (*SvgImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &SvgImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_svg_image_load(fileStr)))
	if err := image_error(img.fail()); err != nil {
		img.Destroy()
		return nil, err
	}
	return img, nil
}

func NewSvgImageFromString(str string) (*SvgImage, error) {
	imagestr := C.CString(str)
	defer C.free(unsafe.Pointer(imagestr))
	img := &SvgImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_svg_image_data(imagestr)))
	if err := image_error(img.fail()); err != nil {
		img.Destroy()
		return nil, err
	}
	return img, nil
}

type PngImage struct {
	RgbImage
}

func NewPngImageLoad(path string) (*PngImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &PngImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_png_image_load(fileStr)))
	if err := image_error(img.fail()); err != nil {
		img.Destroy()
		return nil, err
	}
	return img, nil
}

func NewPngImageFromData(data []uint8) (*PngImage, error) {
	cData := (*C.uchar)(unsafe.Pointer(&data[0]))
	img := &PngImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_png_image_data(cData, C.int(len(data)))))
	if err := image_error(img.fail()); err != nil {
		img.Destroy()
		return nil, err
	}
	return img, nil
}

type JpegImage struct {
	RgbImage
}

func NewJpegImageLoad(path string) (*JpegImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &JpegImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_jpg_image_load(fileStr)))
	if err := image_error(img.fail()); err != nil {
		img.Destroy()
		return nil, err
	}
	return img, nil
}

func NewJpegImageFromData(data []uint8) (*JpegImage, error) {
	cData := (*C.uchar)(unsafe.Pointer(&data[0]))
	img := &JpegImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_jpg_image_data(cData, C.int(len(data)))))
	if err := image_error(img.fail()); err != nil {
		img.Destroy()
		return nil, err
	}
	return img, nil
}

type BmpImage struct {
	RgbImage
}

func NewBmpImageLoad(path string) (*BmpImage, error) {
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &BmpImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_bmp_image_load(fileStr)))
	if err := image_error(img.fail()); err != nil {
		img.Destroy()
		return nil, err
	}
	return img, nil
}

func NewBmpImageFromData(data []uint8) (*BmpImage, error) {
	cData := (*C.uchar)(unsafe.Pointer(&data[0]))
	img := &BmpImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_bmp_image_data(cData, C.long(len(data)))))
	if err := image_error(img.fail()); err != nil {
		img.Destroy()
		return nil, err
	}
	return img, nil
}

type SharedImage struct {
	image
}

var shared_init bool

func register_images() {
	C.go_fltk_register_images()
}

func NewSharedImageLoad(path string) (*SharedImage, error) {
	if !shared_init {
		register_images()
		shared_init = true
	}
	fileStr := C.CString(path)
	defer C.free(unsafe.Pointer(fileStr))
	img := &SharedImage{}
	initImage(img, unsafe.Pointer(C.go_fltk_shared_image_load(fileStr)))
	if img.iPtr == nil {
		return nil, errors.New("shared Image initialization error")
	}
	if err := image_error(img.fail()); err != nil {
		img.Destroy()
		return nil, err
	}
	return img, nil
}
