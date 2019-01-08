// +build darwin linux

package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"strings"
	"time"

	"github.com/golang/freetype/truetype"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
)

type Game struct {
	lastCalc   clock.Time // when we last calculated a frame
	touchCount uint64
	font       *truetype.Font
}

func NewGame() *Game {
	var __g Game
	__g.reset()
	return &__g
}

func (___g2 *Game) reset() {
	var err error
	___g2.font, err = LoadCustomFont()
	if err != nil {
		log.Fatalf("error parsing font: %v", err)
	}
}

func (___g3 *Game) Touch(down bool) {
	if down {
		___g3.touchCount++
	}
}

func (___g3 *Game) Update(now clock.Time) {
	// Compute game states up to now.
	for ; ___g3.lastCalc < now; ___g3.lastCalc++ {
		___g3.calcFrame()
	}
}

func (___g4 *Game) calcFrame() {

}

func (___g5 *Game) _screenRender(___sz3 size.Event, ___glCtx3 gl.Context, ___images3 *glutil.Images) {
	headerHeightPx, footerHeightPx := 100, 100

	header := &_textLineT{
		text:            fmt.Sprintf("%vpx * %vpx", ___sz3.WidthPx, ___sz3.HeightPx),
		font:            ___g5.font,
		widthPx:         ___sz3.WidthPx,
		heightPx:        headerHeightPx,
		textColor:       image.White,
		backgroundColor: image.NewUniform(color.RGBA{0x31, 0xA6, 0xA2, 0xFF}),
		fontSize:        24,
		xPt:             0,
		yPt:             0,
		align:           Left,
	}
	header._lineRender(___sz3)

	loading := &_textLineT{
		placeholder:     "Loading...",
		text:            "Loading" + strings.Repeat(".", int(time.Now().Unix()%4)),
		font:            ___g5.font,
		widthPx:         ___sz3.WidthPx,
		heightPx:        ___sz3.HeightPx - headerHeightPx - footerHeightPx,
		textColor:       image.White,
		backgroundColor: image.NewUniform(color.RGBA{0x35, 0x67, 0x99, 0xFF}),
		fontSize:        96,
		xPt:             0,
		yPt:             PxToPt(___sz3, headerHeightPx),
	}
	loading._lineRender(___sz3)

	footer := &_textLineT{
		text:            fmt.Sprintf("%d", ___g5.touchCount),
		font:            ___g5.font,
		widthPx:         ___sz3.WidthPx,
		heightPx:        footerHeightPx,
		textColor:       image.White,
		backgroundColor: image.NewUniform(color.RGBA{0x31, 0xA6, 0xA2, 0xFF}),
		fontSize:        24,
		xPt:             0,
		yPt:             PxToPt(___sz3, ___sz3.HeightPx-footerHeightPx),
		align:           Right,
	}
	footer._lineRender(___sz3)

	// TODO: think about using Pt for everything?

}// _screenRender

// PxToPt convert a size from pixels to points (based on screen PixelsPerPt)
func PxToPt(___sz2 size.Event, ___sizePx int) geom.Pt {
	return geom.Pt(float32(___sizePx) / ___sz2.PixelsPerPt)
}
