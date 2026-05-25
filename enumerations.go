package fltk

/*
#include "enumerations.h"
*/
import "C"

// Align en: Alignment enums, zh-cn: 对齐方式枚举
type Align uint

var (
	ALIGN_CENTER             = Align(C.go_FL_ALIGN_CENTER)             // en: Center alignment, zh-cn: 中心对齐
	ALIGN_TOP                = Align(C.go_FL_ALIGN_TOP)                // en: Top alignment, zh-cn: 顶部对齐
	ALIGN_BOTTOM             = Align(C.go_FL_ALIGN_BOTTOM)             // en: Bottom alignment, zh-cn: 底部对齐
	ALIGN_LEFT               = Align(C.go_FL_ALIGN_LEFT)               // en: Left alignment, zh-cn: 左对齐
	ALIGN_RIGHT              = Align(C.go_FL_ALIGN_RIGHT)              // en: Right alignment, zh-cn: 右对齐
	ALIGN_INSIDE             = Align(C.go_FL_ALIGN_INSIDE)             // en: Align inside, zh-cn: 内部对齐
	ALIGN_TEXT_OVER_IMAGE    = Align(C.go_FL_ALIGN_TEXT_OVER_IMAGE)    // en: Text over image, zh-cn: 文本在图像上
	ALIGN_IMAGE_OVER_TEXT    = Align(C.go_FL_ALIGN_IMAGE_OVER_TEXT)    // en: Image over text, zh-cn: 图像在文本上
	ALIGN_CLIP               = Align(C.go_FL_ALIGN_CLIP)               // en: Clip alignment, zh-cn: 裁剪对齐
	ALIGN_WRAP               = Align(C.go_FL_ALIGN_WRAP)               // en: Wrap alignment, zh-cn: 自动换行对齐
	ALIGN_IMAGE_NEXT_TO_TEXT = Align(C.go_FL_ALIGN_IMAGE_NEXT_TO_TEXT) // en: Image next to text, zh-cn: 图像在文本旁
	ALIGN_TEXT_NEXT_TO_IMAGE = Align(C.go_FL_ALIGN_TEXT_NEXT_TO_IMAGE) // en: Text next to image, zh-cn: 文本在图像旁
	ALIGN_IMAGE_BACKDROP     = Align(C.go_FL_ALIGN_IMAGE_BACKDROP)     // en: Image backdrop, zh-cn: 图像背景
	ALIGN_TOP_LEFT           = Align(C.go_FL_ALIGN_TOP_LEFT)           // en: Top-left alignment, zh-cn: 左上对齐
	ALIGN_TOP_RIGHT          = Align(C.go_FL_ALIGN_TOP_RIGHT)          // en: Top-right alignment, zh-cn: 右上对齐
	ALIGN_BOTTOM_LEFT        = Align(C.go_FL_ALIGN_BOTTOM_LEFT)        // en: Bottom-left alignment, zh-cn: 左下对齐
	ALIGN_BOTTOM_RIGHT       = Align(C.go_FL_ALIGN_BOTTOM_RIGHT)       // en: Bottom-right alignment, zh-cn: 右下对齐
	ALIGN_LEFT_TOP           = Align(C.go_FL_ALIGN_LEFT_TOP)           // en: Left-top alignment, zh-cn: 左上角对齐
	ALIGN_RIGHT_TOP          = Align(C.go_FL_ALIGN_RIGHT_TOP)          // en: Right-top alignment, zh-cn: 右上角对齐
	ALIGN_LEFT_BOTTOM        = Align(C.go_FL_ALIGN_LEFT_BOTTOM)        // en: Left-bottom alignment, zh-cn: 左下角对齐
	ALIGN_RIGHT_BOTTOM       = Align(C.go_FL_ALIGN_RIGHT_BOTTOM)       // en: Right-bottom alignment, zh-cn: 右下角对齐
	ALIGN_NOWRAP             = Align(C.go_FL_ALIGN_NOWRAP)             // en: No-wrap alignment, zh-cn: 不换行对齐
	ALIGN_POSITION_MASK      = Align(C.go_FL_ALIGN_POSITION_MASK)      // en: Position mask alignment, zh-cn: 位置掩码对齐
	ALIGN_IMAGE_MASK         = Align(C.go_FL_ALIGN_IMAGE_MASK)         // en: Image mask alignment, zh-cn: 图像掩码对齐
)

type ArrowType int

var (
	ARROW_SINGLE = ArrowType(C.go_FL_ARROW_SINGLE)
	ARROW_DOUBLE = ArrowType(C.go_FL_ARROW_DOUBLE)
	ARROW_CHOICE = ArrowType(C.go_FL_ARROW_CHOICE)
	ARROW_RETURN = ArrowType(C.go_FL_ARROW_RETURN)
)

type Orientation int

var (
	ORIENT_NONE  = Orientation(C.go_FL_ORIENT_NONE)
	ORIENT_RIGHT = Orientation(C.go_FL_ORIENT_RIGHT)
	ORIENT_NE    = Orientation(C.go_FL_ORIENT_NE)
	ORIENT_UP    = Orientation(C.go_FL_ORIENT_UP)
	ORIENT_NW    = Orientation(C.go_FL_ORIENT_NW)
	ORIENT_LEFT  = Orientation(C.go_FL_ORIENT_LEFT)
	ORIENT_SW    = Orientation(C.go_FL_ORIENT_SW)
	ORIENT_DOWN  = Orientation(C.go_FL_ORIENT_DOWN)
	ORIENT_SE    = Orientation(C.go_FL_ORIENT_SE)
)

type BoxType int

const (
	NO_BOX                 = BoxType(0)
	FLAT_BOX               = BoxType(1)
	UP_BOX                 = BoxType(2)
	DOWN_BOX               = BoxType(3)
	UP_FRAME               = BoxType(4)
	DOWN_FRAME             = BoxType(5)
	THIN_UP_BOX            = BoxType(6)
	THIN_DOWN_BOX          = BoxType(7)
	THIN_UP_FRAME          = BoxType(8)
	THIN_DOWN_FRAME        = BoxType(9)
	ENGRAVED_BOX           = BoxType(10)
	EMBOSSED_BOX           = BoxType(11)
	ENGRAVED_FRAME         = BoxType(12)
	EMBOSSED_FRAME         = BoxType(13)
	BORDER_BOX             = BoxType(14)
	SHADOW_BOX             = BoxType(15)
	BORDER_FRAME           = BoxType(16)
	SHADOW_FRAME           = BoxType(17)
	ROUNDED_BOX            = BoxType(18)
	RSHADOW_BOX            = BoxType(19)
	ROUNDED_FRAME          = BoxType(20)
	RFLAT_BOX              = BoxType(21)
	ROUND_UP_BOX           = BoxType(22)
	ROUND_DOWN_BOX         = BoxType(23)
	DIAMOND_UP_BOX         = BoxType(24)
	DIAMOND_DOWN_BOX       = BoxType(25)
	OVAL_BOX               = BoxType(26)
	OSHADOW_BOX            = BoxType(27)
	OVAL_FRAME             = BoxType(28)
	OFLAT_BOX              = BoxType(29)
	PLASTIC_UP_BOX         = BoxType(30)
	PLASTIC_DOWN_BOX       = BoxType(31)
	PLASTIC_UP_FRAME       = BoxType(32)
	PLASTIC_DOWN_FRAME     = BoxType(33)
	PLASTIC_THIN_UP_BOX    = BoxType(34)
	PLASTIC_THIN_DOWN_BOX  = BoxType(35)
	PLASTIC_ROUND_UP_BOX   = BoxType(36)
	PLASTIC_ROUND_DOWN_BOX = BoxType(37)
	GTK_UP_BOX             = BoxType(38)
	GTK_DOWN_BOX           = BoxType(39)
	GTK_UP_FRAME           = BoxType(40)
	GTK_DOWN_FRAME         = BoxType(41)
	GTK_THIN_UP_BOX        = BoxType(42)
	GTK_THIN_DOWN_BOX      = BoxType(43)
	GTK_THIN_UP_FRAME      = BoxType(44)
	GTK_THIN_DOWN_FRAME    = BoxType(45)
	GTK_ROUND_UP_FRAME     = BoxType(46)
	GTK_ROUND_DOWN_FRAME   = BoxType(47)
	GLEAM_UP_BOX           = BoxType(48)
	GLEAM_DOWN_BOX         = BoxType(49)
	GLEAM_UP_FRAME         = BoxType(50)
	GLEAM_DOWN_FRAME       = BoxType(51)
	GLEAM_THIN_UP_BOX      = BoxType(52)
	GLEAM_THIN_DOWN_BOX    = BoxType(53)
	GLEAM_ROUND_UP_BOX     = BoxType(54)
	GLEAM_ROUND_DOWN_BOX   = BoxType(55)
	FREE_BOXTYPE           = BoxType(56)
)

// Font en: Font enums, zh-cn: 字体枚举
type Font int

var (
	HELVETICA             = Font(C.go_FL_HELVETICA)             // en: Helvetica, zh-cn: 黑体
	HELVETICA_BOLD        = Font(C.go_FL_HELVETICA_BOLD)        // en: Helvetica Bold, zh-cn: 黑体粗体
	HELVETICA_ITALIC      = Font(C.go_FL_HELVETICA_ITALIC)      // en: Helvetica Italic, zh-cn: 黑体斜体
	HELVETICA_BOLD_ITALIC = Font(C.go_FL_HELVETICA_BOLD_ITALIC) // en: Helvetica Bold Italic, zh-cn: 黑体粗斜体
	COURIER               = Font(C.go_FL_COURIER)               // en: Courier, zh-cn: 宋体
	COURIER_BOLD          = Font(C.go_FL_COURIER_BOLD)          // en: Courier Bold, zh-cn: 宋体粗体
	COURIER_ITALIC        = Font(C.go_FL_COURIER_ITALIC)        // en: Courier Italic, zh-cn: 宋体斜体
	COURIER_BOLD_ITALIC   = Font(C.go_FL_COURIER_BOLD_ITALIC)   // en: Courier Bold Italic, zh-cn: 宋体粗斜体
	TIMES                 = Font(C.go_FL_TIMES)                 // en: Times, zh-cn: Times 字体
	TIMES_BOLD            = Font(C.go_FL_TIMES_BOLD)            // en: Times Bold, zh-cn: Times 粗体
	TIMES_ITALIC          = Font(C.go_FL_TIMES_ITALIC)          // en: Times Italic, zh-cn: Times 斜体
	TIMES_BOLD_ITALIC     = Font(C.go_FL_TIMES_BOLD_ITALIC)     // en: Times Bold Italic, zh-cn: Times 粗斜体
	SYMBOL                = Font(C.go_FL_SYMBOL)                // en: Symbol, zh-cn: 符号
	SCREEN                = Font(C.go_FL_SCREEN)                // en: Screen, zh-cn: 屏幕字体
	SCREEN_BOLD           = Font(C.go_FL_SCREEN_BOLD)           // en: Screen Bold, zh-cn: 屏幕粗体
	ZAPF_DINGBATS         = Font(C.go_FL_ZAPF_DINGBATS)         // en: Zapf Dingbats, zh-cn: Zapf Dingbats 字体
	FREE_FONT             = Font(C.go_FL_FREE_FONT)             // en: Free font, zh-cn: 自由字体
	BOLD                  = Font(C.go_FL_BOLD)                  // en: Bold, zh-cn: 粗体
	ITALIC                = Font(C.go_FL_ITALIC)                // en: Italic, zh-cn: 斜体
	BOLD_ITALIC           = Font(C.go_FL_BOLD_ITALIC)           // en: Bold Italic, zh-cn: 粗斜体
)

type LabelType int

var (
	NORMAL_LABEL = LabelType(C.go_FL_NORMAL_LABEL)
	NO_LABEL     = LabelType(C.go_FL_NO_LABEL)
)

type WrapMode int

const (
	WRAP_NONE      = WrapMode(0)
	WRAP_AT_COLUMN = WrapMode(1)
	WRAP_AT_PIXEL  = WrapMode(2)
	WRAP_AT_BOUNDS = WrapMode(3)
)

type Event int

var (
	NO_EVENT       = Event(C.go_FL_NO_EVENT)
	PUSH           = Event(C.go_FL_PUSH)
	DRAG           = Event(C.go_FL_DRAG)
	RELEASE        = Event(C.go_FL_RELEASE)
	MOVE           = Event(C.go_FL_MOVE)
	MOUSEWHEEL     = Event(C.go_FL_MOUSEWHEEL)
	ENTER          = Event(C.go_FL_ENTER)
	LEAVE          = Event(C.go_FL_LEAVE)
	FOCUS          = Event(C.go_FL_FOCUS)
	UNFOCUS        = Event(C.go_FL_UNFOCUS)
	KEY            = Event(C.go_FL_KEYDOWN)
	KEYDOWN        = Event(C.go_FL_KEYDOWN)
	KEYUP          = Event(C.go_FL_KEYUP)
	CLOSE          = Event(C.go_FL_CLOSE)
	SHORTCUT       = Event(C.go_FL_SHORTCUT)
	DEACTIVATE     = Event(C.go_FL_DEACTIVATE)
	ACTIVATE       = Event(C.go_FL_ACTIVATE)
	HIDE           = Event(C.go_FL_HIDE)
	SHOW           = Event(C.go_FL_SHOW)
	PASTE          = Event(C.go_FL_PASTE)
	SELECTIONCLEAR = Event(C.go_FL_SELECTIONCLEAR)
	DND_ENTER      = Event(C.go_FL_DND_ENTER)
	DND_DRAG       = Event(C.go_FL_DND_DRAG)
	DND_LEAVE      = Event(C.go_FL_DND_LEAVE)
	DND_RELEASE    = Event(C.go_FL_DND_RELEASE)
)

type CallbackCondition int

var (
	WhenNever           = CallbackCondition(C.go_FL_WHEN_NEVER)
	WhenChanged         = CallbackCondition(C.go_FL_WHEN_CHANGED)
	WhenNotChanged      = CallbackCondition(C.go_FL_WHEN_NOT_CHANGED)
	WhenRelease         = CallbackCondition(C.go_FL_WHEN_RELEASE)
	WhenReleaseAlways   = CallbackCondition(C.go_FL_WHEN_RELEASE_ALWAYS)
	WhenEnterKey        = CallbackCondition(C.go_FL_WHEN_ENTER_KEY)
	WhenEnterKeyAlways  = CallbackCondition(C.go_FL_WHEN_ENTER_KEY_ALWAYS)
	WhenEnterKeyChanged = CallbackCondition(C.go_FL_WHEN_ENTER_KEY_CHANGED)
)

var (
	RGB         = int(C.go_FL_RGB)
	INDEX       = int(C.go_FL_INDEX)
	SINGLE      = int(C.go_FL_SINGLE)
	DOUBLE      = int(C.go_FL_DOUBLE)
	ACCUM       = int(C.go_FL_ACCUM)
	ALPHA       = int(C.go_FL_ALPHA)
	DEPTH       = int(C.go_FL_DEPTH)
	STENCIL     = int(C.go_FL_STENCIL)
	RGB8        = int(C.go_FL_RGB8)
	MULTISAMPLE = int(C.go_FL_MULTISAMPLE)
	STEREO      = int(C.go_FL_STEREO)
	FAKE_SINGLE = int(C.go_FL_FAKE_SINGLE)
	OPENGL3     = int(C.go_FL_OPENGL3)
)

// Color en: Color enums, zh-cn: 颜色枚举
type Color uint

var (
	FOREGROUND_COLOR  = Color(C.go_FL_FOREGROUND_COLOR)  // en: Foreground color, zh-cn: 前景色
	BACKGROUND2_COLOR = Color(C.go_FL_BACKGROUND2_COLOR) // en: Background2 color, zh-cn: 背景色2
	INACTIVE_COLOR    = Color(C.go_FL_INACTIVE_COLOR)    // en: Inactive color, zh-cn: 非活动色
	SELECTION_COLOR   = Color(C.go_FL_SELECTION_COLOR)   // en: Selection color, zh-cn: 选择色
	GRAY0             = Color(C.go_FL_GRAY0)             // en: Gray0, zh-cn: 灰色0
	DARK3             = Color(C.go_FL_DARK3)             // en: Dark3, zh-cn: 深色3
	DARK2             = Color(C.go_FL_DARK2)             // en: Dark2, zh-cn: 深色2
	DARK1             = Color(C.go_FL_DARK1)             // en: Dark1, zh-cn: 深色1
	BACKGROUND_COLOR  = Color(C.go_FL_BACKGROUND_COLOR)  // en: Background color, zh-cn: 背景色
	LIGHT1            = Color(C.go_FL_LIGHT1)            // en: Light1, zh-cn: 浅色1
	LIGHT2            = Color(C.go_FL_LIGHT2)            // en: Light2, zh-cn: 浅色2
	LIGHT3            = Color(C.go_FL_LIGHT3)            // en: Light3, zh-cn: 浅色3
	BLACK             = Color(C.go_FL_BLACK)             // en: Black, zh-cn: 黑色
	RED               = Color(C.go_FL_RED)               // en: Red, zh-cn: 红色
	GREEN             = Color(C.go_FL_GREEN)             // en: Green, zh-cn: 绿色
	YELLOW            = Color(C.go_FL_YELLOW)            // en: Yellow, zh-cn: 黄色
	BLUE              = Color(C.go_FL_BLUE)              // en: Blue, zh-cn: 蓝色
	MAGENTA           = Color(C.go_FL_MAGENTA)           // en: Magenta, zh-cn: 洋红色
	CYAN              = Color(C.go_FL_CYAN)              // en: Cyan, zh-cn: 青色
	DARK_RED          = Color(C.go_FL_DARK_RED)          // en: Dark red, zh-cn: 深红色
	DARK_GREEN        = Color(C.go_FL_DARK_GREEN)        // en: Dark green, zh-cn: 深绿色
	DARK_YELLOW       = Color(C.go_FL_DARK_YELLOW)       // en: Dark yellow, zh-cn: 深黄色
	DARK_BLUE         = Color(C.go_FL_DARK_BLUE)         // en: Dark blue, zh-cn: 深蓝色
	DARK_MAGENTA      = Color(C.go_FL_DARK_MAGENTA)      // en: Dark magenta, zh-cn: 深洋红色
	DARK_CYAN         = Color(C.go_FL_DARK_CYAN)         // en: Dark cyan, zh-cn: 深青色
	WHITE             = Color(C.go_FL_WHITE)             // en: White, zh-cn: 白色
)

func ColorFromRgb(r, g, b uint8) Color {
	r1 := uint(r)
	g1 := uint(g)
	b1 := uint(b)
	return Color(((r1 & 0xff) << 24) + ((g1 & 0xff) << 16) + ((b1 & 0xff) << 8) + 0x00)
}

type LineStyle int

var (
	SOLID        = LineStyle(0)
	DASH         = LineStyle(1)
	Dot          = LineStyle(2)
	DASH_DOT     = LineStyle(3)
	DASH_DOT_DOT = LineStyle(4)
	CAP_FLAT     = LineStyle(100)
	CAP_ROUND    = LineStyle(200)
	CAP_SQUARE   = LineStyle(300)
	JOIN_MITER   = LineStyle(1000)
	JOIN_ROUND   = LineStyle(2000)
	JOIN_BEVEL   = LineStyle(3000)
)

var (
	ESCAPE    = int(C.go_FL_ESCAPE)
	TAB       = int(C.go_FL_TAB)
	ENTER_KEY = int(C.go_FL_ENTER_KEY)
	HOME      = int(C.go_FL_HOME)
	LEFT      = int(C.go_FL_LEFT)
	UP        = int(C.go_FL_UP)
	RIGHT     = int(C.go_FL_RIGHT)
	DOWN      = int(C.go_FL_DOWN)
	PAGE_UP   = int(C.go_FL_PAGE_UP)
	PAGE_DOWN = int(C.go_FL_PAGE_DOWN)
	END       = int(C.go_FL_END)
	MENU      = int(C.go_FL_MENU)
	HELP      = int(C.go_FL_HELP)
	F1        = int(C.go_FL_F1)
	F2        = int(C.go_FL_F2)
	F3        = int(C.go_FL_F3)
	F4        = int(C.go_FL_F4)
	F5        = int(C.go_FL_F5)
	F6        = int(C.go_FL_F6)
	F7        = int(C.go_FL_F7)
	F8        = int(C.go_FL_F8)
	F9        = int(C.go_FL_F9)
	F10       = int(C.go_FL_F10)
	F11       = int(C.go_FL_F11)
	F12       = int(C.go_FL_F12)
	DELETE    = int(C.go_FL_DELETE)
	BACKSPACE = int(C.go_FL_BACKSPACE)
	INSERT    = int(C.go_FL_INSERT)
)
