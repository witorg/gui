package gui

import (
	"log"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/davecgh/go-spew/spew"
)

var names = make([]string, 100)

func makeWindowDebug() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	/////////////////////////////////////////////////////
	vbox := addGroup(hbox, "Window")
	cbox := ui.NewCombobox()

	for name, _ := range Data.WindowMap {
		log.Println("range Data.WindowMap() name =", name)
		addName(cbox, name)
	}
	cbox.SetSelected(0)

	vbox.Append(cbox, false)

	cbox.OnSelected(func(*ui.Combobox) {
		x := cbox.Selected()
		log.Println("x =", x)
		log.Println("names[x] =", names[x])
		dumpBox(names[x])
	})

	/////////////////////////////////////////////////////
	vbox = addGroup(hbox, "Debug Window")

	b1 := addButton(vbox, "dumpBox(window)")
	b1.OnClicked(func(*ui.Button) {
		x := cbox.Selected()
		log.Println("x =", x)
		log.Println("names[x] =", names[x])
		dumpBox(names[x])
	})

	b2 := addButton(vbox, "SetMargined(tab)")
	b2.OnClicked(func(*ui.Button) {
		x := cbox.Selected()
		log.Println("x =", x)
		log.Println("FindWindow; names[x] =", names[x])
		gw := FindWindow(names[x])
		if gw == nil {
			return
		}
		if gw.UiTab == nil {
			return
		}
		if gw.TabNumber == nil {
			return
		}
		scs := spew.ConfigState{MaxDepth: 1}
		scs.Dump(gw)
		log.Println("gui.DumpBoxes()\tWindow.UiTab     =", gw.UiTab)
		log.Println("gui.DumpBoxes()\tWindow.TabNumber =", *gw.TabNumber)
		gw.UiTab.SetMargined(*gw.TabNumber, true)
	})

	b3 := addButton(vbox, "Hide(tab)")
	b3.OnClicked(func(*ui.Button) {
		x := cbox.Selected()
		log.Println("x =", x)
		log.Println("FindWindow; names[x] =", names[x])
		gw := FindWindow(names[x])
		if gw == nil {
			return
		}
		if gw.UiTab == nil {
			return
		}
		gw.UiTab.Hide()
	})

	b4 := addButton(vbox, "Show(tab)")
	b4.OnClicked(func(*ui.Button) {
		x := cbox.Selected()
		log.Println("x =", x)
		log.Println("FindWindow; names[x] =", names[x])
		gw := FindWindow(names[x])
		if gw == nil {
			return
		}
		if gw.UiTab == nil {
			return
		}
		gw.UiTab.Show()
	})

	b5 := addButton(vbox, "Delete(tab)")
	b5.OnClicked(func(*ui.Button) {
		x := cbox.Selected()
		log.Println("x =", x)
		log.Println("FindWindow; names[x] =", names[x])
		gw := FindWindow(names[x])
		if gw == nil {
			return
		}
		if gw.UiTab == nil {
			return
		}
		if gw.TabNumber == nil {
			return
		}
		gw.UiTab.Delete(*gw.TabNumber)
	})

	/////////////////////////////////////////////////////
	vbox = addGroup(hbox, "Global Debug")

	dump3 := addButton(vbox, "Dump Windows")
	dump3.OnClicked(func(*ui.Button) {
		DumpWindows()
	})

	dump2 := addButton(vbox, "Dump Boxes")
	dump2.OnClicked(func(*ui.Button) {
		DumpBoxes()
	})

	dump1 := addButton(vbox, "Dump MAP")
	dump1.OnClicked(func(*ui.Button) {
		DumpMap()
	})

	/////////////////////////////////////////////////////
	vbox = addGroup(hbox, "Node Debug")

	n1 := addButton(vbox, "DebugDataNodeMap()")
	n1.OnClicked(func(*ui.Button) {
		DebugDataNodeMap()
	})

	n2 := addButton(vbox, "DebugDataNodeChildren()")
	n2.OnClicked(func(*ui.Button) {
		DebugDataNodeChildren()
	})

	n3 := addButton(vbox, "Node.ListChildren()")
	n3.OnClicked(func(*ui.Button) {
		DebugNodeChildren()
	})

/*
	/////////////////////////////////////////////////////
	vbox = addGroup(hbox, "Numbers")
	pbar := ui.NewProgressBar()
	vbox.Append(pbar, false)
*/


	return hbox
}

var x int = 0

func addName(c *ui.Combobox, s string) {
	c.Append(s)
	names[x] = s
	x = x + 1
}

func addGroup(b *ui.Box, name string) *ui.Box {
	group := ui.NewGroup(name)
	group.SetMargined(true)
	b.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	return vbox
}

func FindWindow(s string) *GuiWindow {
	for name, window := range Data.WindowMap {
		if name == s {
			return window
		}
	}
	log.Printf("COULD NOT FIND WINDOW", s)
	return nil
}

func FindBox(s string) *GuiBox {
	for name, window := range Data.WindowMap {
		if name != s {
			continue
		}
		for name, abox := range window.BoxMap {
			log.Printf("gui.DumpBoxes() \tBOX mapname=%-12s abox.Name=%-12s", name, abox.Name)
			return abox
		}
		log.Println("gui.FindBox() NEED TO INIT WINDOW name =", name)
	}
	log.Println("gui.FindBox() COULD NOT FIND BOX", s)
	return nil
}

func dumpBox(s string) {
	for name, window := range Data.WindowMap {
		if name != s {
			continue
		}
		log.Println("gui.DumpBoxes() MAP: ", name)
		if window.TabNumber == nil {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber = nil")
		} else {
			log.Println("gui.DumpBoxes() \tWindows.TabNumber =", *window.TabNumber)
		}
		log.Println("gui.DumpBoxes()\tWindow.name =", window.Name)
		// log.Println("gui.DumpBoxes()\tWindow.UiWindow type =", reflect.TypeOf(window.UiWindow))
		log.Println("gui.DumpBoxes()\tWindow.UiWindow =", window.UiWindow)
		log.Println("gui.DumpBoxes()\tWindow.UiTab    =", window.UiTab)
		log.Println("gui.dumpBox() BoxMap START")
		for name, abox := range window.BoxMap {
			log.Printf("gui.DumpBoxes() \tBOX mapname=%-12s abox.Name=%-12s", name, abox.Name)
			if name == "MAINBOX" {
				if Config.Debug {
					scs := spew.ConfigState{MaxDepth: 1}
					scs.Dump(abox.UiBox)
				}
			}
		}
		log.Println("gui.dumpBox() BoxMap END")
		if window.UiTab != nil {
			pages := window.UiTab.NumPages()
			log.Println("gui.DumpBoxes()\tWindow.UiTab.NumPages() =", pages)
			tabSetMargined(window.UiTab)
			if Config.Debug {
				scs := spew.ConfigState{MaxDepth: 2}
				scs.Dump(window.UiTab)
			}
		}
	}
}

func addButton(box *ui.Box, name string) *ui.Button {
	button := ui.NewButton(name)

	button.OnClicked(func(*ui.Button) {
		log.Println("Should do something here")
	})

	box.Append(button, false)
	return button
}
