package pdf

import (
	"github.com/signintech/gopdf"
)

const borderAll int = gopdf.Top | gopdf.Right | gopdf.Bottom | gopdf.Left

type Cell interface {
	Text() string
	Width() float64
	Height() float64
	Border() int
}

type cell struct {
	text   string
	width  float64
	height float64
	border int
}

func (c cell) Text() string {
	return c.text
}

func (c cell) Width() float64 {
	return c.width
}

func (c cell) Height() float64 {
	return c.height
}

func (c cell) Border() int {
	return c.border
}

func newCell(text string, width float64, height float64) Cell {
	return cell{text: text, width: width, height: height}
}

func NewHalfWidthPageCell(text string) Cell {
	return cell{text: text, width: gopdf.PageSizeA4.W / 2, height: gopdf.PageSizeA4.H / float64(20), border: borderAll}
}

func NewHalfWidthPageCellWithoutBorder(text string) Cell {
	return cell{text: text, width: gopdf.PageSizeA4.W / 2, height: gopdf.PageSizeA4.H / float64(20), border: 0}
}

func NewFullWidthPageCell(text string) Cell {
	return cell{text: text, width: gopdf.PageSizeA4.W, height: gopdf.PageSizeA4.H / float64(20), border: borderAll}
}
