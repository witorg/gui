package gui

import (
	"log"
)

var names = make([]string, 100)
var nodeNames = make([]string, 100)

var bugWin *Node
/*
	Creates a window helpful for debugging this package
*/
func DebugWindow() {
	Config.Title = "git.wit.org/wit/gui debug fixme"
	Config.Width = 300
	Config.Height = 200
	Config.Exit = StandardClose
	bugWin = NewWindow()
	bugWin.DebugTab("WIT GUI Debug Tab")
}

// this function is used by the examples to add a tab
// dynamically to the bugWin node
// TODO: make this smarter once this uses toolkit/
func DebugTab() {
	if (bugWin == nil) {
		log.Println("Not sure what window to add this to? Use node.DebugTab() instead")
		return;
	}
	bugWin.DebugTab("does this work?")
}

var checkd, checkdn, checkdt, checkdtk *Node

func (n *Node) DebugTab(title string) *Node {
	var newN, gog, g1, g2, g3, dd, gf *Node

	// time.Sleep(1 * time.Second)
	newN = n.NewTab(title)
	newN.Dump()

	gog = newN.NewGroup("GOLANG")
	gog.NewLabel("go language")
	gog.NewButton("GO Language Debug", func () {
		GolangDebugWindow()
	})

	gf = newN.NewGroup("Debug Flags")
	gf.NewLabel("flags to control debugging output")

	checkd = gf.NewCheckbox("Debug")
	checkd.OnChanged = func(*Node) {
		checkd.checked = checkd.toolkit.Checked()
		Config.Debug = checkd.checked
		if (Config.Debug) {
		} else {
		}
	}

	checkdn = gf.NewCheckbox("Debug Node")
	checkdn.OnChanged = func(*Node) {
		checkdn.checked = checkdn.toolkit.Checked()
		Config.DebugNode = checkdn.checked
	}

	checkdd := gf.NewCheckbox("Debug node.Dump()")
	checkdd.OnChanged = func(*Node) {
		Config.DebugDump = checkdd.toolkit.Checked()
	}

	checkdt = gf.NewCheckbox("Debug Tabs")
	checkdtk = gf.NewCheckbox("Debug Toolkit")

//	Debug        bool
//	DebugNode    bool
//	DebugTabs    bool
//	DebugTable   bool
//	DebugWindow  bool
//	DebugToolkit bool

	gog.NewLabel("wit/gui package")
	gog.NewButton("WIT/GUI Package Debug", func () {
		Config.Width = 640
		Config.Height = 480
		Queue(DebugWindow)
	})
	gog.NewButton("Demo wit/gui", func () {
		DemoWindow()
	})
	gog.NewButton("Demo toolkit andlabs/ui", func () {
		DemoToolkitWindow()
	})

	g1 = newN.NewGroup("Current Windows")
	dd = g1.NewDropdown("Window Dropdown")
	log.Println("dd =", dd)

	var dump = false
	for _, child := range Config.master.children {
		log.Println("\t\t", child.id, child.Width, child.Height, child.Name)
		if (child.parent != nil) {
			log.Println("\t\t\tparent =",child.parent.id)
		} else {
			log.Println("\t\t\tno parent")
			panic("no parent")
		}
		if (dump == true) {
			child.Dump()
		}
		dd.AddDropdown(child.Name)
	}
	dd.SetDropdown(0)

	g2 = newN.NewGroup("Debug Window")
	g2.NewButton("SetMargined(tab)", func () {
		log.Println("\tSTART")
		name := dd.GetText()
		log.Println("\tENDed with", name)
		// gw.UiTab.SetMargined(*gw.TabNumber, true)
	})
	g2.NewButton("Hide(tab)", func () {
		// gw.UiTab.Hide()
	})
	g2.NewButton("Show(tab)", func () {
		// gw.UiTab.Show()
	})
	g2.NewButton("Delete(tab)", func () {
		// gw.UiTab.Delete(*gw.TabNumber)
	})
	g2.NewButton("change Title", func () {
		// mainWindow.SetText("hello world")
	})

	/////////////////////////////////////////////////////
	g3 = newN.NewGroup("Node Debug")

	g3.NewButton("Node.Dump()", func () {
		bugWin.Dump()
	})
	g3.NewButton("Node.ListChildren(false)", func () {
		bugWin.ListChildren(false)
	})
	g3.NewButton("Node.ListChildren(true)", func () {
		bugWin.ListChildren(true)
	})
	g3.NewButton("AddDebugTab()", func () {
		if (bugWin != nil) {
			bugWin.DebugTab("added this DebugTab")
		}
	})

	return newN
}
