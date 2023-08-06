#include "image.h"

#include <cstring>

#include <FL/Fl_Image.H>
#include <FL/Fl_BMP_Image.H>
#include <FL/Fl_SVG_Image.H>  
#include <FL/Fl_PNG_Image.H>  
#include <FL/Fl_JPEG_Image.H> 
#include <FL/Fl_RGB_Image.H> 
#include <FL/Fl_Shared_Image.H>

void go_fltk_image_draw(Fl_Image *i, int X, int Y, int W, int H) {
    return i->draw(X, Y, W, H);
}                                     
void go_fltk_image_draw_ext(Fl_Image *i, int X, int Y, int W, int H, int cx, int cy) {
    return i->draw(X, Y, W, H, cx, cy);
}                 
int go_fltk_image_w(Fl_Image *i) {
    return i->w();
}                                                                 
int go_fltk_image_h(Fl_Image *i) {
    return i->h();
}                                                                
void go_fltk_image_delete(Fl_Image *i) {
    i->release();
}                                                               
int go_fltk_image_count(Fl_Image *i) {
    return i->count();
}                                                             
const char *const *go_fltk_image_data(Fl_Image *i) {
    return i->data();
}                                               
Fl_Image *go_fltk_image_copy(Fl_Image *i) {
    return i->copy();
}                                                           
void go_fltk_image_scale(Fl_Image *i, int width, int height, int proportional, int can_expand) {
    return i->scale(width, height, proportional, can_expand);
}   
int go_fltk_image_fail(Fl_Image *i) {
    return i->fail();
}                                                              
int go_fltk_image_data_w(const Fl_Image *i) {
    return i->data_w();
}                                                      
int go_fltk_image_data_h(const Fl_Image *i) {
    return i->data_h();
}                                                      
int go_fltk_image_d(const Fl_Image *i) {
    return i->d();
}                                                           
int go_fltk_image_ld(const Fl_Image *i) {
    return i->ld();
}                                                          
void go_fltk_image_inactive(Fl_Image *i) {
    return i->inactive();
}
  
Fl_SVG_Image *go_fltk_svg_image_load(const char *file) {
    return new Fl_SVG_Image(file);
}

Fl_SVG_Image *go_fltk_svg_image_data(const char *data) {
    return new Fl_SVG_Image(NULL, data);
}

Fl_PNG_Image *go_fltk_png_image_load(const char *file) {
    return new Fl_PNG_Image(file);
}

Fl_PNG_Image *go_fltk_png_image_data(const unsigned char *data, int size) {
    return new Fl_PNG_Image(NULL, data, size);
}

Fl_JPEG_Image *go_fltk_jpg_image_load(const char *file) {
    return new Fl_JPEG_Image(file);
}

Fl_JPEG_Image *go_fltk_jpg_image_data(const unsigned char *data, int len) {
    return new Fl_JPEG_Image(NULL, data, len);
}

Fl_BMP_Image *go_fltk_bmp_image_load(const char *file) {
    return new Fl_BMP_Image(file);
}

Fl_BMP_Image *go_fltk_bmp_image_data(const unsigned char *data, long size) {
    return new Fl_BMP_Image(NULL, data, size);
}

Fl_Shared_Image *go_fltk_shared_image_load(const char *file) {
    return Fl_Shared_Image::get(file, 0, 0);
}

void go_fltk_register_images(void) {
  fl_register_images();
}

Fl_RGB_Image *go_fltk_rgb_image_data(const unsigned char *bits, int bitsLen, int W, int H,
                                     int depth, int ld) {
    // Create copy of bits and pass its ownership to the image object
    unsigned char *const buf = new unsigned char[bitsLen];
    std::memcpy(buf, bits, bitsLen);
    Fl_RGB_Image *img = new Fl_RGB_Image(buf, W, H, depth, ld); 
    img->alloc_array = 1;
    return img;
}

const int go_Fl_Image_ERR_NO_IMAGE = Fl_Image::ERR_NO_IMAGE;
const int go_Fl_Image_ERR_FILE_ACCESS = Fl_Image::ERR_FILE_ACCESS;
const int go_Fl_Image_ERR_FORMAT = Fl_Image::ERR_FORMAT;
const int go_Fl_Image_ERR_MEMORY_ACCESS = Fl_Image::ERR_MEMORY_ACCESS;
