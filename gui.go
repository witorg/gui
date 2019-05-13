package gui

import "log"
import "time"
// import "fmt"

import "github.com/gookit/config"
import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import "github.com/davecgh/go-spew/spew"

var mainwin *ui.Window
var maintab *ui.Tab
var tabcount int

var jcarrButton *ui.Button
var jcarrEntry  *ui.MultilineEntry

type InputData struct {
	Index		int
	CellType	string
	Heading		string
	Color		string
}

func buttonClick(button *ui.Button) {
	log.Println("hostname =", config.String("hostname"), button)
	spew.Dump(button)
	if (jcarrButton == button) {
		log.Println("This is the jcarrButton")
		cur := jcarrEntry.Text()
		jcarrEntry.SetText(cur + "THIS IS A GREAT IDEA\n")
	} else {
		log.Println("This is NOT the jcarrButton")
	}
}

func hostnameButton(hostname string) ui.Control {
	tmpbox := ui.NewHorizontalBox()
	tmpbox.SetPadded(true)

	tmpButton := ui.NewButton(hostname)
	tmpbox.Append(tmpButton, false)
	tmpButton.OnClicked(buttonClick)

	jcarrButton = tmpButton

	return tmpbox
}

func makeGroupEntries() ui.Control {
	group := ui.NewGroup("Entries")
	group.SetMargined(true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	jcarrEntry = ui.NewMultilineEntry()
	entryForm.Append("Entry", ui.NewEntry(), false)
	entryForm.Append("Password Entry", ui.NewPasswordEntry(), false)
	entryForm.Append("Search Entry", ui.NewSearchEntry(), false)
	entryForm.Append("Multiline Entry", jcarrEntry, true)
	entryForm.Append("Multiline Entry No Wrap", ui.NewNonWrappingMultilineEntry(), true)

	return group
}

func makeNumbersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	spinbox := ui.NewSpinbox(47, 100)
	slider  := ui.NewSlider(21, 100)
	pbar    := ui.NewProgressBar()

	spinbox.OnChanged(func(*ui.Spinbox) {
		slider.SetValue(spinbox.Value())
		pbar.SetValue(spinbox.Value())
	})
	slider.OnChanged(func(*ui.Slider) {
		spinbox.SetValue(slider.Value())
		pbar.SetValue(slider.Value())
	})
	vbox.Append(spinbox, false)
	vbox.Append(slider, false)
	vbox.Append(pbar, false)
	vbox.Append(hostnameButton("jcarrtest"), false)

	ip := ui.NewProgressBar()
	ip.SetValue(-1)
	vbox.Append(ip, false)

	group = ui.NewGroup("Lists")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	cbox := ui.NewCombobox()
	cbox.Append("Combobox Item 1")
	cbox.Append("Combobox Item 2")
	cbox.Append("Combobox Item 3")
	vbox.Append(cbox, false)

	ecbox := ui.NewEditableCombobox()
	ecbox.Append("Editable Item 1")
	ecbox.Append("Editable Item 2")
	ecbox.Append("Editable Item 3")
	vbox.Append(ecbox, false)

	rb := ui.NewRadioButtons()
	rb.Append("Radio Button 1")
	rb.Append("Radio Button 2")
	rb.Append("Radio Button 3")
	vbox.Append(rb, false)

	return hbox
}

func makeDataChoosersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	vbox.Append(ui.NewDatePicker(), false)
	vbox.Append(ui.NewTimePicker(), false)
	vbox.Append(ui.NewDateTimePicker(), false)
	vbox.Append(ui.NewFontButton(), false)
	vbox.Append(ui.NewColorButton(), false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	button := ui.NewButton("Open File")
	entry := ui.NewEntry()
	entry.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry.SetText(filename)
	})
	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	button = ui.NewButton("Save File")
	entry2 := ui.NewEntry()
	entry2.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.SaveFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry2.SetText(filename)
	})
	grid.Append(button,
		0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry2,
		1, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	msggrid := ui.NewGrid()
	msggrid.SetPadded(true)
	grid.Append(msggrid,
		0, 2, 2, 1,
		false, ui.AlignCenter, false, ui.AlignStart)

	button = ui.NewButton("Message Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBox(mainwin,
			"This is a normal message box.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	button = ui.NewButton("Error Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBoxError(mainwin,
			"This message box describes an error.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		1, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	return hbox
}

func setupUI() {
	mainwin = ui.NewWindow("Cloud Control Panel", config.Int("width"), config.Int("height"), false)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	maintab = ui.NewTab()
	mainwin.SetChild(maintab)
	mainwin.SetMargined(true)

	// maintab.Append("List examples", makeNumbersPage())
	tabcount = 0
	// maintab.SetMargined(tabcount, true)

/*
	maintab.Append("Choosers examples", makeDataChoosersPage())
	tabcount += 1
	maintab.SetMargined(tabcount, true)

	maintab.Append("Group examples", makeGroupEntries())
	tabcount += 1
	maintab.SetMargined(tabcount, true)
*/

	mainwin.Show()
}

func AddChoosersDemo() {
	maintab.Append("Choosers examples", makeDataChoosersPage())
	maintab.SetMargined(tabcount, true)
	tabcount += 1
}

func AddNewTab(mytab *ui.Tab, newbox ui.Control, tabOffset int) {
	mytab.Append("Cloud Info", newbox)
	mytab.SetMargined(tabOffset, true)
}

// This hangs on GTK
func AddEntriesDemo() {
	maintab.Append("Group examples", makeGroupEntries())
	tabcount += 1
	maintab.SetMargined(tabcount, true)
}

func initColumnNames(mh *TableData, cellJWC string, junk string) {
	if (cellJWC == "BG") {
		mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableColor{})
	} else if (cellJWC == "BUTTON") {
		mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableString(""))
	} else if (cellJWC == "TEXTCOLOR") {
		mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableString(""))
		mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableColor{})
	} else if (cellJWC == "TEXT") {
		mh.generatedColumnTypes = append(mh.generatedColumnTypes, ui.TableString(""))
	} else {
		panic("I don't know what this is in initColumnNames")
	}
}

func initRow(mh *TableData, row int, parts []InputData) {
	tmpBTindex := 0
	for key, foo := range parts {
		log.Println(key, foo)
		if (foo.CellType == "BG") {
			initRowBTcolor        (mh, row, tmpBTindex, parts[key])
			tmpBTindex += 1
		} else if (foo.CellType == "BUTTON") {
			initRowButtonColumn   (mh, row, tmpBTindex,    parts[key].Heading, parts[key])
			tmpBTindex += 1
		} else if (foo.CellType == "TEXTCOLOR") {
			initRowTextColorColumn(mh, row, tmpBTindex, tmpBTindex + 1, parts[key].Heading, ui.TableColor{0.0, 0, 0.9, 1}, parts[key])
			tmpBTindex += 2
		} else if (foo.CellType == "TEXT") {
			initRowTextColumn     (mh, row, tmpBTindex,    parts[key].Heading, parts[key])
			tmpBTindex += 1
		} else {
			panic("I don't know what this is in initColumnNames")
		}
	}
}

func AddSampleTableTab(mytab *ui.Tab, mytabcount int, name string, rowcount int, parts []InputData) {
	mh := new(TableData)

	mh.RowCount    = rowcount
	mh.Rows        = make([]RowData, mh.RowCount)

	// This is the standard callback function from libUI when the user does something
	mh.libUIevent      = defaultSetCellValue

	tmpBTindex := 0

	for key, foo := range parts {
		log.Println(key, foo)
		initColumnNames(mh, foo.CellType, foo.Heading)
	}

	time.Sleep(1 * 1000 * 1000 * 1000)

	for row := 0; row < mh.RowCount; row++ {
		initRow(mh, row, parts)
	}
	log.Println(mh)

	model := ui.NewTableModel(mh)
	table := ui.NewTable(
		&ui.TableParams{
			Model:	model,
			RowBackgroundColorModelColumn:	tmpBTindex,
	})

	for key, foo := range parts {
		log.Println(key, foo)
		initColumnNames(mh, foo.CellType, foo.Heading)
		if (foo.CellType == "BG") {
		} else if (foo.CellType == "BUTTON") {
			tmpBTindex += 1
			table.AppendButtonColumn("button3", tmpBTindex, ui.TableModelColumnAlwaysEditable)
		} else if (foo.CellType == "TEXTCOLOR") {
			tmpBTindex += 1
			appendTextColorColumn   (mh, table, tmpBTindex, tmpBTindex + 1, "testcolor")
			tmpBTindex += 1
		} else if (foo.CellType == "TEXT") {
			tmpBTindex += 1
			appendTextColumn        (mh, table, tmpBTindex, "jwc1col")
		} else {
			panic("I don't know what this is in initColumnNames")
		}
	}

	mytab.Append(name, table)
	mytab.SetMargined(mytabcount, true)
}

func AddTableTab(mytab *ui.Tab, mytabcount int, name string, rowcount int, parts []InputData) *TableData {
	mh := new(TableData)

	mh.RowCount    = rowcount
	mh.Rows        = make([]RowData, mh.RowCount)

	// This is the standard callback function from libUI when the user does something
	mh.libUIevent      = defaultSetCellValue

	tmpBTindex := 0

	for key, foo := range parts {
		log.Println(key, foo)
		initColumnNames(mh, foo.CellType, foo.Heading)
	}

	for row := 0; row < mh.RowCount; row++ {
		initRow(mh, row, parts)
	}
	log.Println(mh)

	model := ui.NewTableModel(mh)
	table := ui.NewTable(
		&ui.TableParams{
			Model:	model,
			RowBackgroundColorModelColumn:	tmpBTindex,
	})

	for key, foo := range parts {
		log.Println(key, foo)
		initColumnNames(mh, foo.CellType, foo.Heading)
		if (foo.CellType == "BG") {
		} else if (foo.CellType == "BUTTON") {
			tmpBTindex += 1
			table.AppendButtonColumn(foo.Heading, tmpBTindex, ui.TableModelColumnAlwaysEditable)
		} else if (foo.CellType == "TEXTCOLOR") {
			tmpBTindex += 1
			appendTextColorColumn   (mh, table, tmpBTindex, tmpBTindex + 1, foo.Heading)
			tmpBTindex += 1
		} else if (foo.CellType == "TEXT") {
			tmpBTindex += 1
			appendTextColumn        (mh, table, tmpBTindex, foo.Heading)
		} else {
			panic("I don't know what this is in initColumnNames")
		}
	}

	mytab.Append(name, table)
	mytab.SetMargined(mytabcount, true)

	return mh
}

func DoGUI() {
	ui.Main(setupUI)

	log.Println("GUI exited. Not sure what to do here. os.Exit() ?")

	// not sure how to pass this back to the main program
	// onExit()
}
