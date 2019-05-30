package gui

import "log"
import "time"
import "regexp"
import "reflect"

import "github.com/andlabs/ui"
import _ "github.com/andlabs/ui/winmanifest"

import pb "git.wit.com/wit/witProtobuf"

// import "github.com/davecgh/go-spew/spew"

func makeCloudInfoBox(wm *WindowMap) *ui.Box {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	if (Data.Debug) {
		log.Println("makeCloudInfoBox() add debugging buttons")
		vbox := ui.NewVerticalBox()
		vbox.SetPadded(true)
		hbox.Append(vbox, false)

		addDebuggingButtons(wm, vbox)

		hbox.Append(ui.NewVerticalSeparator(), false)
	}

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	hostnamebox := ui.NewHorizontalBox()
	hostnamebox.SetPadded(true)
	vbox.Append(hostnamebox, false)

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	hostnamebox.Append(entryForm, true)

	hostnameEntry :=  ui.NewEntry()
	entryForm.Append("hostname:", hostnameEntry, false)
	tmp := Data.Hostname + " (" + Data.IPv6 + ")"
	hostnameEntry.SetText(tmp)
	hostnameEntry.SetReadOnly(true)

	hostnamebox.Append(CreateButton(wm, nil, nil, "Edit", "EDIT", nil), false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	agrid := ui.NewGrid()
	agrid.SetPadded(true)

	agrid.Append(ui.NewLabel("Accounts:"),   0, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Nickname"),    1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Username"),    2, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	agrid.Append(ui.NewLabel("Domain Name"), 3, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	row := 1

	for key, a := range Data.Config.Accounts {
		log.Println("account =          ", key, a)
		log.Println("Accounts[key] =    ", Data.Config.Accounts[key])
		log.Println("account.Nick =     ", Data.Config.Accounts[key].Nick)
		log.Println("account.Username = ", Data.Config.Accounts[key].Username)
		log.Println("account.Token =    ", Data.Config.Accounts[key].Token)

		agrid.Append(ui.NewLabel(Data.Config.Accounts[key].Nick),	1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
		agrid.Append(ui.NewLabel(Data.Config.Accounts[key].Username),	2, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
		agrid.Append(ui.NewLabel(Data.Config.Accounts[key].Domain),	3, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		name := "Login " + Data.Config.Accounts[key].Nick
		l := CreateButton(wm, Data.Config.Accounts[key], nil, name, "LOGIN", nil)
		agrid.Append(l, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		name  = "Show " + Data.Config.Accounts[key].Nick
		b := CreateButton(wm, Data.Config.Accounts[key], nil, name, "SHOW", nil)
		agrid.Append(b, 5, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		row += 1
	}

	row += 1
	agrid.Append(ui.NewLabel(""),    1, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	row += 1
	a := CreateButton(wm, nil, nil, "Add Account", "ADD TAB", nil)
	agrid.Append(a, 4, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	q := CreateButton(wm, nil, nil, "Quit", "QUIT", nil)
	agrid.Append(q, 5, row, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	vbox.Append(agrid, false)
	return hbox
}

//
// THIS IS THE STANDARD VM DISPLAY TABLE
// This maps the 'human' indexed cells in the table
// to the machine's andlabs/libui values. That way
// if you want to work against column 4, then you
// can just reference 4 instead of the internal number
// which could be anything since TEXTCOLOR, TEXT, BG, etc
// fields use between 1 and 3 values internally
//
func AddVmsTab(wm *WindowMap, name string, count int, a *pb.Account) *TableData {
	var parts []TableColumnData

	human := 0

	tmp := TableColumnData{}
	tmp.CellType = "BG"
	tmp.Heading  = "background"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "name"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "hostname"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "IPv6"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "cpus"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "TEXTCOLOR"
	tmp.Heading  = "memory"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	tmp = TableColumnData{}
	tmp.CellType = "BUTTON"
	tmp.Heading  = "Details"
	tmp.Index    = human
	parts = append(parts, tmp)
	human += 1

	mh := AddTableTab(wm, wm.T, 1, name, count, parts, a)
	return mh
}

func ShowAccountQuestionTab(wm *WindowMap) {
	wm.T.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	wm.Box2 = AddAccountQuestionBox(wm)
	wm.T.InsertAt("New Account?", 0, wm.Box2)
	wm.T.SetMargined(0, true)
}

func ShowAccountTab(wm *WindowMap, i int) {
	log.Println("ShowAccountTab() START")

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	// Create the things for the Account Tab
	abox := AddAccountBox(wm)

	// Set the parents and data structure links
	// aTab.me = wm.T
	// aTab.parentWindow = Data.Window1.W
	// aTab.tabOffset = 0

	if (i >= 0) {
		log.Println("ShowAccountTab() InsertAt i=", i)
		wm.T.Delete(0)
		wm.T.InsertAt("Add Account", i, abox)
		wm.T.SetMargined(0, true)
	} else {
		// TODO: After append try to discover the tab index #
		log.Println("ShowAccountTab() Append")
		AddBoxToTab("Create New Account", wm.T, abox)
	}
}

func ShowMainTab(wm *WindowMap) {
	log.Println("ShowMainTab() wm =", wm)
	log.Println("ShowMainTab() wm.T =", wm.T)
	log.Println("ShowMainTab() wm.T =", wm.T)
	wm.T.Delete(0)

	log.Println("Sleep(200)")
	time.Sleep(200 * time.Millisecond)

	wm.Box2 = makeCloudInfoBox(wm)
	wm.T.InsertAt("Main", 0, wm.Box2)
	wm.T.SetMargined(0, true)
}

func StartNewWindow(c *pb.Config, bg bool, action string) {
	log.Println("InitNewWindow() Create a new window")
	var newWindowMap WindowMap
	newWindowMap.C = c
	newWindowMap.Action = action
	Data.Windows = append(Data.Windows, &newWindowMap)

	ui.OnShouldQuit(func() bool {
		// mouseClick(&newBM)
                ui.Quit()
		return true
	})

	for i, aWM := range(Data.Windows) {
		log.Println(aWM)
		if (aWM.W == nil) {
			log.Println("ShowWindow() THIS WINDOW IS NOT YET SHOWN i=", i)
			log.Println("ShowWindow() THIS WINDOW IS NOT YET SHOWN aWM=", reflect.TypeOf(aWM))
			// Data.NewWindow = &aWM
			// Data.NewWindow = i
			// ui.Main(InitWindow)
			time.Sleep(200 * time.Millisecond)
			if (bg) {
				ui.Main(func() {
					InitWindow(Data.Windows[i], i)
				})
			} else {
				ui.Main(func() {
					InitWindow(Data.Windows[i], i)
				})
			}
			return
		}
	}
}

func getSplashText() *ui.AttributedString {
	var aText  *ui.AttributedString
	aText = ui.NewAttributedString("Click to continue")
	return aText
}


func InitWindow(wm *WindowMap, i int) {
	log.Println("InitWindow() THIS WINDOW IS NOT YET SHOWN")
	// i := Data.NewWindow
	log.Println("InitWindow() THIS WINDOW IS NOT YET SHOWN i=", i)

	c := Data.Windows[i].C
	Data.Windows[i].W = ui.NewWindow("", int(c.Width), int(c.Height), true)
	Data.Windows[i].W.SetBorderless(false)

        // create a 'fake' button entry for the mouse clicks
	var newBM ButtonMap
	newBM.Action	= "QUIT"
	newBM.W		= Data.Windows[i].W
	newBM.WM	= Data.Windows[i]
	Data.AllButtons = append(Data.AllButtons, newBM)

	Data.Windows[i].W.OnClosing(func(*ui.Window) bool {
		log.Println("InitWindow() OnClosing() THIS WINDOW IS CLOSING i=", i)
		// mouseClick(&newBM)
                ui.Quit()
		return true
	})

	Data.Windows[i].T = ui.NewTab()
	Data.Windows[i].W.SetChild(Data.Windows[i].T)
	Data.Windows[i].W.SetMargined(true)

	log.Println("InitWindow() i =", i)
	log.Println("InitWindow() Data.Windows[i] =", Data.Windows[i])
	log.Println("InitWindow() Data.Windows[i].Action =", Data.Windows[i].Action)

	if (Data.Windows[i].Action == "SPLASH") {
		log.Println("InitWindow() TRYING SPLASH")
		tmp := getSplashText()
		log.Println("InitWindow() TRYING SPLASH tmp =", tmp)
		Data.Windows[i].Box1 = ShowSplashBox(Data.Windows[i], i, tmp)

		Data.Windows[i].T.Append("WIT Splash", Data.Windows[i].Box1)
		Data.Windows[i].T.SetMargined(0, true)
	}

	Data.Windows[i].W.Show()
	Data.State = "splash"
}

// makeEntryBox(box, "hostname:", "blah.foo.org") {
func makeEntryVbox(hbox *ui.Box, a string, startValue string, edit bool, action string) *EntryMap {
	// Start 'Nickname' vertical box
	vboxN := ui.NewVerticalBox()
	vboxN.SetPadded(true)
	vboxN.Append(ui.NewLabel(a), false)

	e := defaultMakeEntry(startValue, edit, action)

	vboxN.Append(e.E, false)
	hbox.Append(vboxN, false)
	// End 'Nickname' vertical box

	return e
}

/*
// string handling examples that might be helpful for normalizeInt()
isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

for _, username := range []string{"userone", "user2", "user-three"} {
    if !isAlpha(username) {
        fmt.Printf("%q is not valid\n", username)
    }
}

const alpha = "abcdefghijklmnopqrstuvwxyz"

func alphaOnly(s string) bool {
   for _, char := range s {
      if !strings.Contains(alpha, strings.ToLower(string(char))) {
         return false
      }
   }
   return true
}
*/

func normalizeInt(s string) string {
	// reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Println("normalizeInt() regexp.Compile() ERROR =", err)
		return s
	}
	clean := reg.ReplaceAllString(s, "")
	log.Println("normalizeInt() s =", clean)
	return clean
}

func defaultEntryChange(e *ui.Entry) {
	for key, em := range Data.AllEntries {
		if (Data.Debug) {
			log.Println("\tdefaultEntryChange() Data.AllEntries =", key, em)
		}
		if Data.AllEntries[key].E == e {
			log.Println("defaultEntryChange() FOUND", 
				"action =", Data.AllEntries[key].Action,
				"Last =", Data.AllEntries[key].Last,
				"e.Text() =", e.Text())
			Data.AllEntries[key].Last = e.Text()
			if Data.AllEntries[key].Normalize != nil {
				fixed := Data.AllEntries[key].Normalize(e.Text())
				e.SetText(fixed)
			}
			return
		}
	}
	log.Println("defaultEntryChange() ERROR. MISSING ENTRY MAP. e.Text() =", e.Text())
}

func defaultMakeEntry(startValue string, edit bool, action string) *EntryMap {
	e := ui.NewEntry()
	e.SetText(startValue)
	if (edit == false) {
		e.SetReadOnly(true)
	}
	e.OnChanged(defaultEntryChange)

	// add the entry field to the global map
	var newEntryMap EntryMap
	newEntryMap.E      = e
	newEntryMap.Edit   = edit
	newEntryMap.Action = action
	if (action == "Memory") {
		newEntryMap.Normalize = normalizeInt
	}
	Data.AllEntries = append(Data.AllEntries, newEntryMap)

	return &newEntryMap
}

func makeEntryHbox(hbox *ui.Box, a string, startValue string, edit bool, action string) *EntryMap {
	// Start 'Nickname' vertical box
	hboxN := ui.NewHorizontalBox()
	hboxN.SetPadded(true)
	hboxN.Append(ui.NewLabel(a), false)

	e := defaultMakeEntry(startValue, edit, action)
	hboxN.Append(e.E, false)

	hbox.Append(hboxN, false)
	// End 'Nickname' vertical box

	return e
}

func AddBoxToTab(name string, tab *ui.Tab, box *ui.Box) {
	tab.Append(name, box)
	tab.SetMargined(0, true)
}
