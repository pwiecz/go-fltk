#include "tree.h"

#include <FL/Fl_Tree.H>

#include "event_handler.h"

class GTree : public EventHandler<Fl_Tree> {
public:
  GTree(int x, int y, int w, int h, const char *label)
      : EventHandler<Fl_Tree>(x, y, w, h, label) {}
};

    
GTree *go_fltk_new_Tree(int x, int y, int w, int h, const char *label) {
  return new GTree(x, y, w, h, label);
}

void go_fltk_Tree_set_show_root(Fl_Tree *tree, int show) {
  tree->showroot(show);
}  

Fl_Tree_Item* go_fltk_Tree_add(Fl_Tree *tree, const char *path) {
  return tree->add(path);
}

int go_fltk_Tree_remove(Fl_Tree *tree, Fl_Tree_Item *item) {
  return tree->remove(item);
}

void go_fltk_Tree_clear(Fl_Tree *tree) {
  tree->clear();
}

void go_fltk_Tree_clear_children(Fl_Tree *tree, Fl_Tree_Item *item) {
  tree->clear_children(item);
}  

void go_fltk_Tree_Item_set_widget(Fl_Tree_Item *item, Fl_Widget *widget) {
  item->widget(widget);
}

const unsigned int go_FL_TREE_ITEM_DRAW_DEFAULT = (unsigned int)FL_TREE_ITEM_DRAW_DEFAULT;
const unsigned int go_FL_TREE_ITEM_DRAW_LABEL_AND_WIDGET = (unsigned int)FL_TREE_ITEM_DRAW_LABEL_AND_WIDGET;
const unsigned int go_FL_TREE_ITEM_HEIGHT_FROM_WIDGET = (unsigned int)FL_TREE_ITEM_HEIGHT_FROM_WIDGET;

void go_fltk_Tree_set_item_draw_mode(Fl_Tree *tree, unsigned int drawMode) {
  tree->item_draw_mode((Fl_Tree_Item_Draw_Mode)drawMode);
}

const int go_FL_TREE_CONNECTOR_NONE = (int)FL_TREE_CONNECTOR_NONE;
const int go_FL_TREE_CONNECTOR_DOTTED = (int)FL_TREE_CONNECTOR_DOTTED;
const int go_FL_TREE_CONNECTOR_SOLID = (int)FL_TREE_CONNECTOR_SOLID;
void go_fltk_Tree_set_connector_style(Fl_Tree *tree, int style) {
  tree->connectorstyle((Fl_Tree_Connector)style);
}

const int go_FL_TREE_SELECT_NONE = (int)FL_TREE_SELECT_NONE;
const int go_FL_TREE_SELECT_SINGLE = (int)FL_TREE_SELECT_SINGLE;
const int go_FL_TREE_SELECT_MULTI = (int)FL_TREE_SELECT_MULTI;
const int go_FL_TREE_SELECT_SINGLE_DRAGGABLE = (int)FL_TREE_SELECT_SINGLE_DRAGGABLE;

void go_fltk_Tree_set_select_mode(Fl_Tree *tree, int selectMode) {
  tree->selectmode((Fl_Tree_Select)selectMode);
}  
