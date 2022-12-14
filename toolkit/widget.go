package toolkit

// passes information between the toolkit library (plugin)

// All Toolkit interactions should be done via a channel or Queue()

// This is the only thing that is passed between the toolkit plugin

// what names should be used? This is not part of [[Graphical Widget]]
// Event() seems like a good name.
// Could a protobuf be used here? (Can functions be passed?)
type Widget struct {
	i     int
	s     string

	Name   string
	Width  int
	Height int
	X      int
	Y      int

	Custom    func()
	Event     func(*Widget) *Widget

	// Probably deprecate these
//	OnChanged func(*Widget)
//	Custom    func(*Widget)
//	OnExit    func(*Widget)
}

type Blah struct {
	i     int
	s     string

	Name   string
	Width  int
	Height int
}

/*
type Widget int

// https://ieftimov.com/post/golang-datastructures-trees/
const (
	Unknown Widget = iota
	Window
	Tab
	Frame
	Dropbox
	Spinner
	Label
)

func (s Widget) String() string {
	switch s {
	case Window:
		return "Window"
	case Tab:
		return "Tab"
	case Frame:
		return "Frame"
	case Label:
		return "Label"
	case Dropbox:
		return "Dropbox"
	}
	return "unknown"
}
*/
