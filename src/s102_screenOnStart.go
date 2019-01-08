package main

import (
	//"math/rand"
	//"time"

	//"golang.org/x/mobile/app"
	//"golang.org/x/mobile/event/key"
	//"golang.org/x/mobile/event/lifecycle"
	//"golang.org/x/mobile/event/paint"
	//"golang.org/x/mobile/event/size"
	//"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/gl/glutil"
	//"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/gl"
)

func onStart(glctx gl.Context) {
	_glImageS = glutil.NewImages(glctx)
	_GamE = NewGame()
} // onStart

