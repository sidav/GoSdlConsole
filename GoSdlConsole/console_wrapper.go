package GoSdlConsole

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"strings"
)

const ( // for the great compatibility with default console color codes
	BLACK        = 0
	DARK_RED     = 1
	DARK_GREEN   = 2
	DARK_YELLOW  = 3
	DARK_BLUE    = 4
	DARK_MAGENTA = 5
	DARK_CYAN    = 6
	BEIGE        = 7
	DARK_GRAY    = 8
	RED          = 9
	GREEN        = 10
	YELLOW       = 11
	BLUE         = 12
	MAGENTA      = 13
	CYAN         = 14
	WHITE        = 15
)

var (
	winTitle                  = "Go-SDL2 Texture"
	chrW, chrH          int32 = 10, 16
	termW, termH        int32 = 80, 25
	winWidth, winHeight       = termW*chrW, termH*chrH
	FontPngFileName           = "assets/font_10x16.png"
	window              *sdl.Window
	renderer            *sdl.Renderer
	texture             *sdl.Texture
	fontImg             *sdl.Surface
	src, dst            sdl.Rect
	err                 error

	fgColor = []uint8{255, 255, 255}
	bgColor = []uint8{0, 0, 0}

	compatColorTable = map[int][]uint8{
		BLACK:        {0, 0, 0},
		DARK_RED:     {128, 0, 0},
		DARK_GREEN:   {0, 128, 0},
		DARK_YELLOW:  {128, 128, 0},
		DARK_BLUE:    {0, 0, 128},
		DARK_MAGENTA: {128, 0, 128},
		DARK_CYAN:    {0, 128, 128},
		BEIGE:        {128, 128, 96},
		DARK_GRAY:    {96, 96, 96},
		RED:          {255, 0, 0},
		GREEN:        {0, 255, 0},
		YELLOW:       {255, 255, 0},
		BLUE:         {0, 0, 255},
		MAGENTA:      {255, 0, 255},
		CYAN:         {0, 255, 255},
		WHITE:        {255, 255, 255},
	}

	isShiftBeingHeld bool
)

func prepareFont() {
	// pix := fontImg.Pixels()
	fontImg.SetColorKey(true, 0xff00ff)
	//fmt.Printf("SIZE: %d with %d/%d; ", len(fontImg.Pixels()), fontImg.PixelNum(), fontImg.BytesPerPixel())
	//fmt.Printf("pix: %d ", fontImg.Pixels()[0])
	//fmt.Printf("%d ", fontImg.Pixels()[1])
	//fmt.Printf("%d ", fontImg.Pixels()[2])
	//fmt.Printf("%d ", fontImg.Pixels()[3])
	//for i := 0; i < len(pix); i+= 4 {
	//
	//}
}

func Init_console() {

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return
	}

	fontImg, err = img.Load(FontPngFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		return
	}
	prepareFont()

	texture, err = renderer.CreateTextureFromSurface(fontImg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return
	}
	renderer.Clear()
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.FillRect(&sdl.Rect{0, 0, int32(winWidth), int32(winHeight)})
	renderer.Copy(texture, &src, &dst)
	renderer.Present()
}

func Close_console() { //should be deferred!
	window.Destroy()
	renderer.Destroy()
	texture.Destroy()
	fontImg.Free()
}

func Clear_console() {
	renderer.Clear()
}

func Flush_console() {
	renderer.Present()
}

func GetConsoleSize() (int, int) {
	return int(termW), int(termH)
}

func SetFgColorRGB(r, g, b uint8) {
	fgColor[0] = r
	fgColor[1] = g
	fgColor[2] = b
}

func SetBgColorRGB(r, g, b uint8) {
	bgColor[0] = r
	bgColor[1] = g
	bgColor[2] = b
}

func SetColor(fg int, bg int) {
	SetFgColor(fg)
	SetBgColor(bg)
}

func SetFgColor(fg int) {
	rgb := compatColorTable[fg]
	SetFgColorRGB(rgb[0], rgb[1], rgb[2])
}

func SetBgColor(bg int) {
	rgb := compatColorTable[bg]
	SetBgColorRGB(rgb[0], rgb[1], rgb[2])
}

func PutChar(c rune, x, y int) {
	code := int32(c)
	if code < 256 {
		row := code / 16
		col := code % 16
		src = sdl.Rect{chrW * col, chrH * row, chrW, chrH}
		dst = sdl.Rect{chrW * int32(x), chrH * int32(y), chrW, chrH}
		renderer.SetDrawColor(bgColor[0], bgColor[1], bgColor[2], 255)
		renderer.FillRect(&sdl.Rect{chrW * int32(x), chrH * int32(y), chrW, chrH})
		texture.SetColorMod(fgColor[0], fgColor[1], fgColor[2])
		renderer.Copy(texture, &src, &dst)
	}
}

func PutString(s string, x, y int) {
	length := len([]rune(s))
	for i := 0; i < length; i++ {
		PutChar([]rune(s)[i], x+i, y)
	}
}

func ReadKey() string {
	for {
		event := sdl.WaitEvent() // wait here until an event is in the event queue
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			if t.State == 1 {
				keyString := sdl.GetScancodeName(t.Keysym.Scancode)
				if (t.Keysym.Mod & sdl.KMOD_SHIFT) != 1 && len(keyString) == 1 {
					return strings.ToLower(keyString)
				}
				return strings.ToUpper(keyString)
			}
		}
	}
}

// DEPRECATED CRAP BELOW

//func OldTestReadKeyDeleteMePlease() string {
//	break_loop := false
//	sdl.StartTextInput()
//	for !break_loop {
//		event := sdl.WaitEvent() // wait here until an event is in the event queue
//		switch t := event.(type) {
//		case *sdl.TextInputEvent:
//			return string(t.Text[:])
//		case *sdl.KeyboardEvent:
//			fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tcode:%d\tmodifiers:%d\tstate:%d\trepeat:%d\n",
//				t.Timestamp, t.Type, sdl.GetScancodeName(t.Keysym.Scancode), t.Keysym.Scancode, t.Keysym.Mod, t.State, t.Repeat)
//			// return t.Keysym.Sym
//			break_loop = true
//		}
//	}
//	return "WTF"
//}

//func workEvents() {
//	event := sdl.WaitEvent() // wait here until an event is in the event queue
//	switch t := event.(type) {
//	case *sdl.MouseMotionEvent:
//		fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
//			t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
//	case *sdl.MouseButtonEvent:
//		fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
//			t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
//	case *sdl.MouseWheelEvent:
//		fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
//			t.Timestamp, t.Type, t.Which, t.X, t.Y)
//	case *sdl.KeyboardEvent:
//		fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
//			t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
//	}
//}
