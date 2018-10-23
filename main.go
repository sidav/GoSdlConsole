package main

import (
	cw "GoSdlConsole/GoSdlConsole"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cw.Init_console()
	defer cw.Close_console()

	start := time.Now()

	frames := 0
	for {
		cw.Clear_console()
		for x := 0; x < 80; x++ {
			for y := 0; y < 25; y++ {
				cw.SetColor(rand.Int() % 16, rand.Int() % 16)
				cw.PutChar(rune(rand.Int() % 255), x, y)
			}
		}
		cw.Flush_console()
		frames++
		elapsed := time.Since(start)
		if elapsed >= time.Second {
			break
		}
	}
	fmt.Printf("FPS: %d", frames)
	time.Sleep(20000 * time.Millisecond)

}
