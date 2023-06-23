package fltk

/*
#include "tile.h"
*/
import "C"
import "unsafe"

type Tile struct {
	Group
}

func NewTile(x, y, w, h int, text ...string) *Tile {
	t := &Tile{}
	initWidget(t, unsafe.Pointer(C.go_fltk_new_Tile(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return t
}
