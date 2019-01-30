package txt2img // import "github.com/po3rin/txt2img"

import (
	"image"

	"image/color"
	"image/draw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// Drawer impliments set params & exec draw.
type Drawer interface {
	SetImg(img image.Image)
	SetColors(textColor color.RGBA)
	SetFont(font *truetype.Font)
	SetFontSize(fontSize float64)
	SetTextPos(textPosVertical, textPosHorizontal int)
	Draw(text string) (img *image.RGBA, err error)
}

// Drawer has data to create images.
type drawer struct {
	Img               image.Image
	Font              *truetype.Font
	FontSize          float64
	TextColor         *image.Uniform
	TextPosVertical   int
	TextPosHorizontal int
}

// Params is parameters for NewDrawer function
type Params struct {
	Img               image.Image
	Font              *truetype.Font
	FontSize          float64
	TextColor         color.RGBA
	TextPosVertical   int
	TextPosHorizontal int
}

// NewDrawer init drawer
func NewDrawer(p Params) (Drawer, error) {
	d := &drawer{}
	d.SetImg(p.Img)
	d.SetColors(p.TextColor)
	d.SetFontSize(p.FontSize)
	d.SetFont(p.Font)
	d.SetTextPos(p.TextPosVertical, p.TextPosHorizontal)
	return d, nil
}

// SetBackgroundImage sets the backgroundImage.
func (d *drawer) SetImg(img image.Image) {
	d.Img = img
}

// SetColors sets the textColor.
func (d *drawer) SetColors(textColor color.RGBA) {
	d.TextColor = image.NewUniform(textColor)
}

// SetFont sets the font.
func (d *drawer) SetFont(font *truetype.Font) {
	d.Font = font
}

// SetFontSize sets the fontSize.
func (d *drawer) SetFontSize(fontSize float64) {
	d.FontSize = fontSize
}

// SetTextPos sets the textPosition.
func (d *drawer) SetTextPos(textPosVertical, textPosHorizontal int) {
	d.TextPosVertical = textPosVertical
	d.TextPosHorizontal = textPosHorizontal
}

// Draw draw text to images.
func (d *drawer) Draw(text string) (img *image.RGBA, err error) {
	imgRect := image.Rectangle{image.Pt(0, 0), d.Img.Bounds().Size()}
	img = image.NewRGBA(imgRect)
	draw.Draw(img, img.Bounds(), d.Img, image.ZP, draw.Src)

	if d.Font != nil {
		c := freetype.NewContext()
		c.SetDPI(72)
		c.SetFont(d.Font)
		c.SetFontSize(d.FontSize)
		c.SetClip(img.Bounds())
		c.SetDst(img)
		c.SetSrc(d.TextColor)
		c.SetHinting(font.HintingNone)
		pt := freetype.Pt(d.TextPosHorizontal, d.TextPosVertical)
		_, err = c.DrawString(text, pt)
		if err != nil {
			return nil, err
		}
	}
	return img, nil
}
