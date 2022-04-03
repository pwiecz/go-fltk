#pragma once

#include "_cgo_export.h"


class WidgetWithEventHandler {
public:
  virtual void set_event_handler(int handlerId) = 0;
};

template<class BaseWidget>
class EventHandler : public BaseWidget, public WidgetWithEventHandler {
public:
  template<class... Arg>
  EventHandler(Arg... args)
    : BaseWidget(args...) {}

  int handle(int event) final {
    if (m_eventHandlerId >= 0) {
      const int ret = _go_eventHandler(m_eventHandlerId, event);
      if (ret != 0) {
	return ret;
      }
    }
    return BaseWidget::handle(event);
  }

  void set_event_handler(int handlerId) final {
    m_eventHandlerId = handlerId;
  }

protected:
  int m_eventHandlerId = -1;

};
