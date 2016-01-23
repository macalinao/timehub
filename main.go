package main

import "github.com/nsf/termbox-go"

func switch_page(dir int) {
}

func draw_all() {
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	draw_all()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowDown, termbox.KeyArrowRight:
				switch_page(1)
				draw_all()
			case termbox.KeyArrowUp, termbox.KeyArrowLeft:
				switch_page(-1)
				draw_all()
			}
		case termbox.EventResize:
			draw_all()
		}
	}

}
