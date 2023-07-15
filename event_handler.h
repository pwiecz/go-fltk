#pragma once

#include "_cgo_export.h"

#include <vector>


class WidgetWithEventHandler {
public:
  virtual void set_event_handler(int handlerId) = 0;
};
class WidgetWithResizeHandler {
public:
  virtual void set_resize_handler(uintptr_t handlerId) = 0;
};
class WidgetWithDrawHandler {
public:
  virtual void set_draw_handler(uintptr_t handlerId) = 0;
  virtual void basedraw() = 0;  
};
class WidgetWithDeletionHandler {
public:
  virtual void add_deletion_handler(uintptr_t handlerId) = 0;
};

template<class BaseWidget>
class EventHandler : public BaseWidget, public WidgetWithEventHandler, public WidgetWithDrawHandler, public WidgetWithResizeHandler, public WidgetWithDeletionHandler {
public:
  template<class... Arg>
  EventHandler(Arg... args)
    : BaseWidget(args...) {}

  virtual ~EventHandler() {
    for (uintptr_t deletionHandlerId : m_deletionHandlerIds) {
      _go_callbackHandler(deletionHandlerId);
    }
  }

  int handle(int event) final {
    if (m_eventHandlerId >= 0) {
      const int ret = _go_eventHandler(m_eventHandlerId, event);
      if (ret != 0) {
        return ret;
      }
    }
    return BaseWidget::handle(event);
  }

  void draw() override {
    if (m_drawHandlerId != 0) {
      _go_drawHandler(m_drawHandlerId, this);
    } else {
      BaseWidget::draw();
    }
  }

  void basedraw() final {
    BaseWidget::draw();
  }    
    
  void resize(int x, int y, int w, int h) final {
    BaseWidget::resize(x, y, w, h);
    if (m_resizeHandlerId != 0) {
      _go_callbackHandler(m_resizeHandlerId);
    }
  }


  void set_event_handler(int handlerId) final {
    m_eventHandlerId = handlerId;
  }

  void set_draw_handler(uintptr_t handlerId) final {
    m_drawHandlerId = handlerId;
  }

  void set_resize_handler(uintptr_t handlerId) final {
    m_resizeHandlerId = handlerId;
  }

  void add_deletion_handler(uintptr_t handlerId) final {
    m_deletionHandlerIds.push_back(handlerId);
  }

protected:
  int m_eventHandlerId = -1;
  uintptr_t m_drawHandlerId = 0;
  uintptr_t m_resizeHandlerId = 0;
  std::vector<uintptr_t> m_deletionHandlerIds;
};
