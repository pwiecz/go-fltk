package fltk

/*
#include <stdlib.h>
#include "dialogs.h"
*/
import "C"
import "unsafe"

func MessageBox(title, message string) {
	titleStr := C.CString(title)
	defer C.free(unsafe.Pointer(titleStr))
	messageStr := C.CString(message)
	defer C.free(unsafe.Pointer(messageStr))
	C.go_fltk_message_box(titleStr, messageStr)
}

func ChoiceDialog(message string, options ...string) int {
	if len(options) == 0 || len(options) >= 3 {
		panic("Unsupported number of ChoiceDialog options")
	}
	messageStr := C.CString(message)
	defer C.free(unsafe.Pointer(messageStr))
	option0 := C.CString(options[0])
	defer C.free(unsafe.Pointer(option0))
	option1 := (*C.char)(nil)
	if len(options) > 1 {
		option1 = C.CString(options[1])
		defer C.free(unsafe.Pointer(option1))
	}
	option2 := (*C.char)(nil)
	if len(options) > 2 {
		option2 = C.CString(options[2])
		defer C.free(unsafe.Pointer(option2))
	}
	return int(C.go_fltk_choice_dialog(messageStr, option0, option1, option2))
}

// AlertDialog - Display an alert dialog (fltk_alert)
func AlertDialog(message string) {
	messageStr := C.CString(message)
	defer C.free(unsafe.Pointer(messageStr))
	C.go_fltk_alert_dialog(messageStr)
}

// AskDialog - Display confirmation dialog (fltk_ask). Noted that this is marked s depricated in fltk
func AskDialog(message string) int {
	messageStr := C.CString(message)
	defer C.free(unsafe.Pointer(messageStr))
	return int(C.go_fltk_ask_dialog(messageStr))
}

// InputDialog - Display input dialog and get the return value as string. If the dialog is cancelled return empty string. maxchar - maximum chars for the taking the input. message: prompt message to user.
func InputDialog(maxchar int, message string) string {
	messageStr := C.CString(message)
	defer C.free(unsafe.Pointer(messageStr))
	output := C.go_fltk_input_dialog(C.int(maxchar), messageStr)
	defer C.free(unsafe.Pointer(output))
	return C.GoString(output)
}
