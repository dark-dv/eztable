package eztable

// Style contains the data for a table style.
type Style struct {
	Header *Header
	Body   *Body
}

// Header contains the style data for a table header.
type Header struct {
	TopLeftCorner      string
	TopRightCorner     string
	TopIntersection    string
	TopRow             string
	HeaderLeft         string
	HeaderSplitter     string
	HeaderRight        string
	BottomLeft         string
	BottomRight        string
	BottomIntersection string
	BottomRow          string
}

// Body contains the style data for a table body.
type Body struct {
	BodyLeft           string
	BodyRight          string
	BodySplitter       string
	BottomLeftCorner   string
	BottomRightCorner  string
	BottomIntersection string
	BottomRow          string
}

// HasTop will check if a header has a top item above it.
func (h *Header) HasTop() bool {
	return h.TopLeftCorner != "" && h.TopRightCorner != "" && h.TopRow != "" && h.TopIntersection != ""
}

// HasBottom will check if a header has a top item below it.
func (h *Header) HasBottom() bool {
	return h.BottomLeft != "" && h.BottomRight != "" && h.BottomRow != "" && h.BottomIntersection != ""
}

// HasBottom will check if a body has a base item below it.
func (b *Body) HasBottom() bool {
	return b.BottomLeftCorner != "" && b.BottomRightCorner != "" && b.BottomRow != "" && b.BottomIntersection != ""
}

var (
	Lite *Style = &Style{
		Header: &Header{
			BottomLeft:         " ",
			BottomRight:        " ",
			BottomIntersection: " ",
			BottomRow:          "-",
		},
		Body: &Body{
			BodyLeft:     " ",
			BodyRight:    " ",
			BodySplitter: " ",
		},
	}

	Unicode *Style = &Style{
		Header: &Header{
			TopLeftCorner:      "╔",
			TopRightCorner:     "╗",
			TopIntersection:    "╤",
			TopRow:             "═",
			HeaderLeft:         "║",
			HeaderSplitter:     "│",
			HeaderRight:        "║",
			BottomLeft:         "╟",
			BottomRight:        "╢",
			BottomIntersection: "┼",
			BottomRow:          "━",
		},

		Body: &Body{
			BodyLeft:           "║",
			BodyRight:          "║",
			BodySplitter:       "│",
			BottomLeftCorner:   "╚",
			BottomRightCorner:  "╝",
			BottomIntersection: "╧",
			BottomRow:          "═",
		},
	}
)
