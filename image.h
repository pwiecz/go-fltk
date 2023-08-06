#pragma once

#include <stdint.h>


#ifdef __cplusplus
extern "C" {
#endif

    typedef struct Fl_Image Fl_Image;
    typedef struct Fl_SVG_Image Fl_SVG_Image;
    typedef struct Fl_PNG_Image Fl_PNG_Image;
    typedef struct Fl_JPEG_Image Fl_JPEG_Image;
    typedef struct Fl_BMP_Image Fl_BMP_Image;
    typedef struct Fl_Shared_Image Fl_Shared_Image;
    typedef struct Fl_RGB_Image Fl_RGB_Image;

    extern void go_fltk_image_draw(Fl_Image *, int X, int Y, int W, int H);                                        
    extern void go_fltk_image_draw_ext(Fl_Image *, int X, int Y, int W, int H, int cx, int cy);                    
    extern int go_fltk_image_w(Fl_Image *);                                                                    
    extern int go_fltk_image_h(Fl_Image *);                                                                   
    extern void go_fltk_image_delete(Fl_Image *);                                                                  
    extern int go_fltk_image_count(Fl_Image *self);                                                                
    extern const char *const *go_fltk_image_data(Fl_Image *self);                                                  
    extern Fl_Image *go_fltk_image_copy(Fl_Image *self);                                                              
    extern void go_fltk_image_scale(Fl_Image *self, int width, int height, int proportional, int can_expand);      
    extern int go_fltk_image_fail(Fl_Image *self);                                                                 
    extern int go_fltk_image_data_w(const Fl_Image *self);                                                         
    extern int go_fltk_image_data_h(const Fl_Image *self);                                                         
    extern int go_fltk_image_d(const Fl_Image *self);                                                              
    extern int go_fltk_image_ld(const Fl_Image *self);                                                             
    extern void go_fltk_image_inactive(Fl_Image *self);

    extern Fl_SVG_Image *go_fltk_svg_image_load(const char *file);
    extern Fl_SVG_Image *go_fltk_svg_image_data(const char *data);
    extern Fl_PNG_Image *go_fltk_png_image_load(const char *file);
    extern Fl_PNG_Image *go_fltk_png_image_data(const unsigned char *data, int size);
    extern Fl_JPEG_Image *go_fltk_jpg_image_load(const char *file);
    extern Fl_JPEG_Image *go_fltk_jpg_image_data(const unsigned char *data, int len);
    extern Fl_BMP_Image *go_fltk_bmp_image_load(const char *file);
    extern Fl_BMP_Image *go_fltk_bmp_image_data(const unsigned char *data, long size);
    extern Fl_Shared_Image *go_fltk_shared_image_load(const char *file);
    extern Fl_RGB_Image *go_fltk_rgb_image_data(const unsigned char *bits, int bitsLen, int W, int H, int depth, int ld);

    extern void go_fltk_register_images(void);

    extern const int go_Fl_Image_ERR_NO_IMAGE;
    extern const int go_Fl_Image_ERR_FILE_ACCESS;
    extern const int go_Fl_Image_ERR_FORMAT;
    extern const int go_Fl_Image_ERR_MEMORY_ACCESS;
  
#ifdef __cplusplus
}
#endif
