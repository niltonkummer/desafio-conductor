package pdf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"math"
	"time"

	"github.com/leekchan/accounting"
	"github.com/shopspring/decimal"
	"github.com/signintech/gopdf"
)

const (
	FontRegular = "Roboto Regular"
	FontMedium  = "Roboto Medium"
	FontBold    = "Roboto Bold"

	FontSizePeriod    float64 = 16
	FontSpacingPeriod float64 = 16
	Margin            float64 = 30
	FontSize          float64 = 12
	FontSpacing       float64 = 6

	FontSizeRow    float64 = 10
	FontSpacingRow float64 = 16
)

var formatter = accounting.Accounting{Symbol: "R$ ", Precision: 2, Thousand: ".", Decimal: ","}

var (
	ColorBlack       = []uint8{44, 44, 44}
	ColorGrey        = []uint8{142, 142, 147}
	ColorTableHeader = []uint8{44, 44, 46}
)

var (
	//go:embed static/images/logo.jpeg
	logoImage []byte

	//go:embed static/fonts/Roboto-Regular.ttf
	fontRegularTTF []byte

	//go:embed static/fonts/Roboto-Medium.ttf
	fontMediumTTF []byte

	//go:embed static/fonts/Roboto-Bold.ttf
	fontBoldTTF []byte
)

type Row interface {
	Description() string
	Amount() decimal.Decimal
	CreatedAt() time.Time
}

type Rows []Row

func GeneratePDF(w io.Writer, period time.Time, location *time.Location, rows Rows) error {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) // 595x841 = A4
	pdf.AddPage()
	pdf.SetInfo(gopdf.PdfInfo{
		Title:        "Fatura",
		Subject:      "Cliente",
		CreationDate: time.Now().In(location),
	})
	pdf.SetMargins(Margin, Margin, Margin, Margin)

	// fonts
	if err := pdf.AddTTFFontByReader(FontRegular, bytes.NewBuffer(fontRegularTTF)); err != nil {
		return fmt.Errorf("could not use Roboto Regular; %w", err)
	}

	if err := pdf.AddTTFFontByReader(FontMedium, bytes.NewBuffer(fontMediumTTF)); err != nil {
		return fmt.Errorf("could not use Roboto Medium; %w", err)
	}

	if err := pdf.AddTTFFontByReader(FontBold, bytes.NewBuffer(fontBoldTTF)); err != nil {
		return fmt.Errorf("could not use Roboto Bold; %w", err)
	}

	// reset position to margin
	pdf.SetY(Margin)
	pdf.SetX(Margin)

	// Left Header
	{
		logoImageHolder, err := gopdf.ImageHolderByBytes(logoImage)
		if err != nil {
			return fmt.Errorf("could not create image holder: %w", err)
		}

		if err := pdf.ImageByHolder(logoImageHolder, Margin, Margin, nil); err != nil {
			return fmt.Errorf("could not place logo: %w", err)
		}
		pdf.SetY(Margin + 80)
	}

	startUserRow := pdf.GetY() + FontSize*2

	// Right header
	{
		_ = pdf.SetFont(FontBold, "", int(FontSize))
		pdf.SetTextColor(ColorBlack[0], ColorBlack[1], ColorBlack[2])

		text := "Fatura "
		size := gopdf.PageSizeA4.W/2 - Margin

		pdf.SetY(Margin + 80)
		pdf.SetX(gopdf.PageSizeA4.W - Margin - size)
		_ = pdf.Cell(nil, text)

		pdf.SetY(startUserRow)

		// right
		{
			_ = pdf.SetFont(FontRegular, "", int(FontSize))
			pdf.SetTextColor(ColorBlack[0], ColorBlack[1], ColorBlack[2])
			pdf.SetX(gopdf.PageSizeA4.W - Margin - size)
			_ = pdf.Cell(nil, "Documento emitido em:")

			pdf.SetY(pdf.GetY() + FontSize + FontSpacing)
			pdf.SetX(gopdf.PageSizeA4.W - Margin - size)
			pdf.SetTextColor(ColorGrey[0], ColorGrey[1], ColorGrey[2])
			_ = pdf.Cell(nil, time.Now().In(location).Format("02/01/2006 15:04"))

			pdf.SetY(pdf.GetY() + FontSize + FontSpacing + 10)
			_ = pdf.SetFont(FontRegular, "", int(FontSize))
			pdf.SetTextColor(ColorBlack[0], ColorBlack[1], ColorBlack[2])
			pdf.SetX(gopdf.PageSizeA4.W - Margin - size)
			_ = pdf.Cell(nil, "Valor Total: ")

			pdf.SetY(pdf.GetY() + FontSize + FontSpacing)
			pdf.SetTextColor(ColorGrey[0], ColorGrey[1], ColorGrey[2])
			_ = pdf.SetFont(FontRegular, "", int(FontSize))
			pdf.SetX(gopdf.PageSizeA4.W - Margin - size)
			_ = pdf.Cell(nil, func() string {
				var sum decimal.Decimal

				for _, r := range rows {
					sum = sum.Add(r.Amount())
				}
				return formatter.FormatMoneyBigRat(sum.Rat())
			}())
		}
	}

	// Customer
	{
		pdf.SetX(Margin)
		pdf.SetY(startUserRow + FontSize)

		_ = pdf.SetFont(FontBold, "", int(FontSize))
		pdf.SetTextColor(ColorBlack[0], ColorBlack[1], ColorBlack[2])
		_ = pdf.Text("Nome: ")

		_ = pdf.SetFont(FontRegular, "", int(FontSize))
		pdf.SetTextColor(ColorGrey[0], ColorGrey[1], ColorGrey[2])
		_ = pdf.Text("John Doe")

		pdf.Br(FontSize + FontSpacing)
	}

	pdf.Br(FontSpacingPeriod * 2)

	var (
		footerSize      = (FontSize + FontSpacing) * 3
		pageCounterSize = (FontSize + FontSpacing) * 2

		PeriodSize    = FontSizePeriod + FontSpacingPeriod*2
		tableHeadSize = FontSize + FontSpacingRow
		tableRowSize  = FontSizeRow + FontSpacingRow

		emptyPageSize = gopdf.PageSizeA4.H - Margin*2
	)

	bodyLeft := gopdf.PageSizeA4.H - pdf.GetY() - Margin - pageCounterSize
	neededSize := PeriodSize + tableHeadSize + tableRowSize*float64(len(rows)) + footerSize

	page := 1
	pageCount := int(math.Ceil((neededSize-bodyLeft)/emptyPageSize)) + 1

	bodyLeft -= PeriodSize

	pdf.Br(FontSpacingPeriod)
	_ = pdf.SetFont(FontBold, "", int(FontSizePeriod))
	pdf.SetTextColor(ColorBlack[0], ColorBlack[1], ColorBlack[2])
	_ = pdf.Cell(nil, period.Format("01/2006"))
	pdf.Br(FontSizePeriod + FontSpacingPeriod)

	// table
	var start int
	for {

		split := false
		end := len(rows)
		fit := true

		bodyLeft := gopdf.PageSizeA4.H - pdf.GetY() - Margin - pageCounterSize

		if fit {
			maxRows := int(math.Ceil((bodyLeft - tableHeadSize) / tableRowSize))

			if end > start+maxRows {
				end = start + maxRows

				if end < start {
					fit = false
					end = start
				}
			}

		} else {
			// go to next loop
			end = start
		}

		if start != end {
			writeTable(&pdf, location, rows[:end][start:])
			start = end
		}

		if start == len(rows) {
			break
		}

		if split && fit {
			continue
		}

		// page count
		pdf.SetY(gopdf.PageSizeA4.H - Margin - FontSize)
		_ = pdf.SetFont(FontRegular, "", int(FontSize))
		pdf.SetTextColor(ColorGrey[0], ColorGrey[1], ColorGrey[2])
		_ = pdf.CellWithOption(
			&gopdf.Rect{
				W: gopdf.PageSizeA4.W - Margin*2,
				H: FontSize,
			},
			fmt.Sprintf("%d/%d", page, pageCount),
			gopdf.CellOption{Align: gopdf.Right},
		)

		pdf.AddPage()
		page++
	}

	bodyLeft = gopdf.PageSizeA4.H - Margin - pdf.GetY()
	if bodyLeft < footerSize {
		_ = pdf.SetFont(FontRegular, "", int(FontSize))
		pdf.SetTextColor(ColorGrey[0], ColorGrey[1], ColorGrey[2])
		pdf.SetY(gopdf.PageSizeA4.H - Margin - FontSize)
		_ = pdf.CellWithOption(
			&gopdf.Rect{
				W: gopdf.PageSizeA4.W - Margin*2,
				H: FontSize,
			},
			fmt.Sprintf("%d/%d", page, pageCount),
			gopdf.CellOption{Align: gopdf.Right},
		)

		pdf.AddPage()
		page++
	}

	// footer
	{
		_ = pdf.SetFont(FontBold, "", int(FontSize))
		pdf.SetTextColor(ColorBlack[0], ColorBlack[1], ColorBlack[2])
		pdf.SetY(gopdf.PageSizeA4.H - Margin - FontSize - FontSpacing)
		pdf.SetX(Margin)
		_ = pdf.Text("Conductor")

		pdf.SetY(gopdf.PageSizeA4.H - Margin - FontSize)
		pdf.SetX(Margin)
		_ = pdf.CellWithOption(
			&gopdf.Rect{
				W: (gopdf.PageSizeA4.W - Margin*2) / 2,
				H: FontSize,
			},
			"00-000.000/0001-01",
			gopdf.CellOption{Align: gopdf.Left},
		)

		pdf.SetTextColor(ColorGrey[0], ColorGrey[1], ColorGrey[2])
		_ = pdf.CellWithOption(
			&gopdf.Rect{
				W: (gopdf.PageSizeA4.W - Margin*2) / 2,
				H: FontSize,
			},
			fmt.Sprintf("%d/%d", page, pageCount),
			gopdf.CellOption{Align: gopdf.Right},
		)
	}

	return pdf.Write(w)
}

func writeTable(pdf *gopdf.GoPdf, timezone *time.Location, rows []Row) {
	// A4 = 595x842
	sizeColumns := []float64{40, 295, 100, 100}
	columns := []string{"Data", "Descrição", "Valor"}

	_ = pdf.SetFont(FontMedium, "", int(FontSize))
	pdf.SetTextColor(ColorTableHeader[0], ColorTableHeader[1], ColorTableHeader[2])

	dir := gopdf.Left
	for x, c := range columns {
		_ = pdf.CellWithOption(&gopdf.Rect{
			W: sizeColumns[x],
			H: FontSize,
		},
			c,
			gopdf.CellOption{Align: dir},
		)
		dir = gopdf.Right
	}
	pdf.Br(FontSize + FontSpacingRow)

	_ = pdf.SetFont(FontRegular, "", int(FontSizeRow))
	pdf.SetTextColor(ColorGrey[0], ColorGrey[1], ColorGrey[2])
	pdf.SetStrokeColor(ColorGrey[0], ColorGrey[1], ColorGrey[2])
	pdf.SetLineWidth(0.25)

	for _, row := range rows {
		dir = gopdf.Left
		for column := 0; column < len(columns); column++ {
			var text string

			switch column {
			case 0:
				text = row.CreatedAt().In(timezone).Format("02/01")
			case 1:
				text = row.Description()
				size, _ := pdf.MeasureTextWidth(" ...")
				args, _ := pdf.SplitText(text, sizeColumns[column]-size)
				if len(args) > 1 {
					text = args[0] + "..."
				}

			case 2:
				text = decimal.Zero.String()
				text = formatter.FormatMoneyBigRat(row.Amount().Rat())

			}

			_ = pdf.CellWithOption(&gopdf.Rect{
				W: sizeColumns[column],
				H: FontSizeRow,
			},
				text,
				gopdf.CellOption{Align: dir},
			)
			dir = gopdf.Right
		}

		pdf.Br(FontSizeRow + FontSpacingRow)
	}
}
