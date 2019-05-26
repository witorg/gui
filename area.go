package gui

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

func makeSplashArea() *ui.Area {
	// make this button just to get the default font (but don't display the button)
	// There should be another way to do this (?)
	Data.fontButton = CreateFontButton("SplashFont", "DONE")

	makeAttributedString()
	splashArea     := ui.NewArea(myAH)

	// create a 'fake' button entry for the mouse clicks
	var newmap ButtonMap
	newmap.Action	= "AREA"
	newmap.A	= splashArea

	myAH.button	= &newmap
	Data.AllButtons = append(Data.AllButtons, newmap)
	Data.splashArea = splashArea

	if (Data.Debug) {
		spew.Dump(Data.splashArea)
		log.Println("DEBUGGING", Data.Debug)
	} else {
		log.Println("NOT DEBUGGING AREA mhAH.button =", myAH.button)
		log.Println("NOT DEBUGGING AREA mhAH.button =", myAH.button)
		log.Println("NOT DEBUGGING AREA mhAH.button =", myAH.button)
	}
	return Data.splashArea
}

func appendWithAttributes(what string, attrs ...ui.Attribute) {
	start := len(Data.attrstr.String())
	end := start + len(what)
	Data.attrstr.AppendUnattributed(what)
	for _, a := range attrs {
		Data.attrstr.SetAttribute(a, start, end)
	}
}

func makeAttributedString() {
	Data.attrstr = ui.NewAttributedString("")

	appendWithAttributes("Welcome to the Cloud Control Panel\n", ui.TextSize(16), ui.TextColor{0.0, 0.0, 0.8, .8}) // "RGBT"

	appendWithAttributes("(alpha)\n\n", ui.TextSize(10))

	appendWithAttributes("This control panel was designed to be an interface to your 'private' cloud. ", ui.TextWeightBold)
	appendWithAttributes("The concept of a private cloud means that you can use a providers system, or, seemlessly, use your own hardware in your own datacenter. ", ui.TextWeightBold)

	Data.attrstr.AppendUnattributed("\n")
	Data.attrstr.AppendUnattributed("\n")
	appendWithAttributes("This control panel requires:\n")
	Data.attrstr.AppendUnattributed("\n")
	appendWithAttributes("IPv6\n")
	appendWithAttributes("Your hostname in DNS\n")
	Data.attrstr.AppendUnattributed("\n\n\n\n\n")

	appendWithAttributes("<click or press any key>\n", ui.TextSize(10))
}

func (ah areaHandler) Draw(a *ui.Area, p *ui.AreaDrawParams) {
	tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
		String:		Data.attrstr,
		DefaultFont:	Data.fontButton.Font(),
		Width:		p.AreaWidth,
		Align:		ui.DrawTextAlign(1),
	})
	p.Context.Text(tl, 0, 0)
	defer tl.Free()
}

func (ah areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	if (Data.Debug) {
		log.Println("GOT MouseEvent()")
		spew.Dump(me)
	}
	if (me.Down == 1) {
		log.Println("GOT MOUSE DOWN")
	}
	if (me.Up == 1) {
		log.Println("GOT MOUSE UP")
		log.Println("GOT MOUSE UP")
		log.Println("GOT MOUSE UP")
		mouseClick(myAH.button)
	}
}

func (ah areaHandler) MouseCrossed(a *ui.Area, left bool) {
	log.Println("GOT MouseCrossed()")
}

func (ah areaHandler) DragBroken(a *ui.Area) {
	log.Println("GOT DragBroken()")
}

func (ah areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	log.Println("GOT KeyEvent()")
	if (ke.Key == 10) {
		log.Println("GOT ENTER")
		log.Println("GOT ENTER")
		log.Println("GOT ENTER")
	}
	if (ke.Key == 32) {
		log.Println("GOT ENTER")
		log.Println("GOT ENTER")
		log.Println("GOT ENTER")
	}
	spew.Dump(ke)
	// splashWin.Destroy()
	// ui.Quit()
	return false
}
