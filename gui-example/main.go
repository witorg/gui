package main

import (
	"log"
	"os"
	"time"

	"git.wit.org/wit/gui"
)

func customExit(gw *gui.GuiWindow) {
	log.Println("Should Exit Here")
	os.Exit(0)
}

func main() {
	log.Println("starting my Control Panel")

	gui.Config.Width = 800
	gui.Config.Height = 300
	gui.Config.Exit = customExit

	go gui.Main(initGUI)

	watchGUI()
}

func initGUI() {
	n := gui.NewWindow("WIT GUI Example Window", 640, 480)
	n.AddDemoTab("A Simple Tab Demo")
}

// This demonstrates how to properly interact with the GUI
// You can not involke the GUI from external goroutines in most cases.

func watchGUI() {
	var i = 1
	for {
		log.Println("Waiting for customExit()", i)
		i += 1
		time.Sleep(1 * time.Second)
		if i == 4 {
			log.Println("Opening a Debug Window via the gui.Queue()")
			gui.Queue(gui.DebugWindow)
		}
	}
}
