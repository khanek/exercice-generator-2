package pdf

import (
	"bytes"
	"io"

	"khanek/exercise-generator/config/bindatafs"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

// Writer interface allow export to PDF
type Writer interface {
	ToPDF() (*gopdf.GoPdf, error)
}

func getFontReader() (io.Reader, error) {
	assetFS := bindatafs.AssetFS.NameSpace("admin")
	data, err := assetFS.Asset("assets/fonts/Roboto-Regular.ttf")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

// NewPDF returns initialized instance of pdf
func NewPDF() (*gopdf.GoPdf, error) {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	fontReader, err := getFontReader()
	if err != nil {
		return nil, err
	}
	err = pdf.AddTTFFontByReader("roboto", fontReader)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't add font")
	}
	err = pdf.SetFont("roboto", "", 14)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't set font")
	}
	return pdf, nil
}

// AddCellsPage adds new page with cells
func AddCellsPage(pdf *gopdf.GoPdf, cells []Cell) error {
	pdf.AddPage()
	const margin float64 = 0
	var positionX, positionY float64
	pdf.SetX(margin)
	pdf.SetY(margin)
	align := gopdf.Center | gopdf.Middle
	for _, cell := range cells {
		width := cell.Width()
		height := cell.Height()
		err := pdf.CellWithOption(&gopdf.Rect{W: width, H: height}, cell.Text(), gopdf.CellOption{Align: align, Border: cell.Border()})
		if err != nil {
			return errors.Wrap(err, "Error when calling pdf.CellWithOption")
		}
		if positionX+width >= gopdf.PageSizeA4.W {
			positionX = margin
			positionY += height
		} else {
			positionX += width
		}
		pdf.SetX(positionX)
		pdf.SetY(positionY)
	}
	return nil
}
