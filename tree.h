#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct GTree GTree;
  typedef struct Fl_Tree_Item Fl_Tree_Item;
  typedef struct Fl_Widget Fl_Widget;

  extern GTree* go_fltk_new_Tree(int x, int y, int w, int h, const char* text);

  extern void go_fltk_Tree_set_show_root(GTree* tree, int show);
  extern Fl_Tree_Item* go_fltk_Tree_add(GTree* tree, const char* path);

  extern void go_fltk_Tree_Item_set_widget(Fl_Tree_Item* item, Fl_Widget* widget);

  extern const unsigned int go_FL_TREE_ITEM_DRAW_DEFAULT;
  extern const unsigned int go_FL_TREE_ITEM_DRAW_LABEL_AND_WIDGET;
  extern const unsigned int go_FL_TREE_ITEM_HEIGHT_FROM_WIDGET;
  extern void go_fltk_Tree_set_item_draw_mode(GTree* tree, unsigned int drawMode);

  extern const int go_FL_TREE_CONNECTOR_NONE;
  extern const int go_FL_TREE_CONNECTOR_DOTTED;
  extern const int go_FL_TREE_CONNECTOR_SOLID;
  extern void go_fltk_Tree_set_connector_style(GTree* tree, int style);

#ifdef __cplusplus
}
#endif
