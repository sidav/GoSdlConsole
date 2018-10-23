package main

import (
	"fmt"
	"time"
	cw "GoSdlConsole/GoSdlConsole"
)

func main() {
	cw.Init_console()
	defer cw.Close_console()

	for i:=0; i < 5000; i++{
		cw.SetColor(cw.RED, cw.BLACK)
		cw.SetBgColorRGB(0, 0, 0)
		cw.PutString("Look, I am", 0, 0)
		cw.SetFgColorRGB(0, 128, 32)
		cw.PutString("being rendered by ", 3, 4)
		cw.PutString("F", 79, 24)
		cw.SetFgColorRGB(255, 255, 32)
		cw.SetBgColorRGB(0, 255, 32)
		cw.PutString("OPENGL", 21, 4)
		cw.Flush_console()
		fmt.Println(cw.ReadKey())
		time.Sleep(20 * time.Millisecond)
	}
}
