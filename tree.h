#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Tree Fl_Tree;
  typedef struct GTree GTree;
  typedef struct Fl_Tree_Item Fl_Tree_Item;
  typedef struct Fl_Widget Fl_Widget;

  extern GTree* go_fltk_new_Tree(int x, int y, int w, int h, const char* text);

  extern void go_fltk_Tree_set_show_root(Fl_Tree* tree, int show);
  extern Fl_Tree_Item* go_fltk_Tree_add(Fl_Tree* tree, const char* path);
  extern int go_fltk_Tree_remove(Fl_Tree *tree, Fl_Tree_Item *item);
  extern void go_fltk_Tree_clear(Fl_Tree *tree);
  extern void go_fltk_Tree_clear_children(Fl_Tree* tree, Fl_Tree_Item* item);  

  extern void go_fltk_Tree_Item_set_widget(Fl_Tree_Item* item, Fl_Widget* widget);

  extern const unsigned int go_FL_TREE_ITEM_DRAW_DEFAULT;
  extern const unsigned int go_FL_TREE_ITEM_DRAW_LABEL_AND_WIDGET;
  extern const unsigned int go_FL_TREE_ITEM_HEIGHT_FROM_WIDGET;
  extern void go_fltk_Tree_set_item_draw_mode(Fl_Tree* tree, unsigned int drawMode);

  extern const int go_FL_TREE_CONNECTOR_NONE;
  extern const int go_FL_TREE_CONNECTOR_DOTTED;
  extern const int go_FL_TREE_CONNECTOR_SOLID;
  extern void go_fltk_Tree_set_connector_style(Fl_Tree *tree, int style);

  extern const int go_FL_TREE_SELECT_NONE;
  extern const int go_FL_TREE_SELECT_SINGLE;
  extern const int go_FL_TREE_SELECT_MULTI;
  extern const int go_FL_TREE_SELECT_SINGLE_DRAGGABLE;
  extern void go_fltk_Tree_set_select_mode(Fl_Tree* tree, int selectMode);

#ifdef __cplusplus
}
#endif
