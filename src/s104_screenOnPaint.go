package main

import (
	//"math/rand"
	"time"

	//"golang.org/x/mobile/app"
	//"golang.org/x/mobile/event/key"
	//"golang.org/x/mobile/event/lifecycle"
	//"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	//"golang.org/x/mobile/event/touch"
	//"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/gl"
)

func _screenOnPaint(___glCtx gl.Context, ___sz size.Event) {
	___glCtx.ClearColor(1, 1, 1, 1)
	___glCtx.Clear(gl.COLOR_BUFFER_BIT)
	__now := clock.Time(time.Since(_startTime) * 60 / time.Second)
	_GamE.Update(__now)
	_GamE._screenRender(___sz, ___glCtx, _glImageS)
} // _screenOnPaint

