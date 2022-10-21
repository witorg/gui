package toolkit

import "log"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

// create a new box
func (t *Toolkit) GetBox() *ui.Box {
	return t.uiBox
}

// create a new box
func (t *Toolkit) NewBox() *Toolkit {
	log.Println("gui.Toolbox.NewBox() START create default")
	t.Dump()
	if (t.uiGroup != nil) {
		log.Println("\tgui.Toolbox.NewBox() is a Group")
		var newTK Toolkit

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiGroup.SetChild(vbox)
		newTK.uiBox = vbox

		return &newTK
	}
	if (t.uiBox != nil) {
		log.Println("\tgui.Toolbox.NewBox() is a Box")
		var newTK Toolkit

		vbox := ui.NewVerticalBox()
		vbox.SetPadded(padded)
		t.uiBox.Append(vbox, stretchy)
		newTK.uiBox = vbox
		newTK.Name = t.Name

		return &newTK
	}
	log.Println("\tgui.Toolbox.NewBox() FAILED. Couldn't figure out where to make a box")
	t.Dump()
	return nil
}
