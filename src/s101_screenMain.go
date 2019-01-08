package main

import (
	"math/rand"
	"time"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	//"golang.org/x/mobile/exp/gl/glutil"
	//"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/gl"
)


func _screenMain() {
	rand.Seed(time.Now().UnixNano())

	app.Main(func(a app.App) {
		var glctx gl.Context
		var sz size.Event
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
					onStart(glctx)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					onStop()
					glctx = nil
				}
				// switch e.Crosses(lifecycle.StageFocused) {
				// case lifecycle.CrossOn:
				// 	glctx, _ = e.DrawContext.(gl.Context)
				// 	onStart(glctx)
				// 	a.Send(paint.Event{})
				// case lifecycle.CrossOff:
				// 	glctx.ClearColor(1, 1, 1, 1)
				// 	glctx.Clear(gl.COLOR_BUFFER_BIT)
				// 	a.Publish()
				// 	onStop()
				// 	glctx = nil
				// }
			case size.Event:
				sz = e
			case paint.Event:
				if glctx == nil || e.External {
					continue
				}
				onPaint(glctx, sz)
				a.Publish()
				a.Send(paint.Event{}) // keep animating
			case touch.Event:
				if down := e.Type == touch.TypeBegin; down || e.Type == touch.TypeEnd {
					game.Touch(down)
				}
			case key.Event:
				if e.Code != key.CodeSpacebar {
					break
				}
				if down := e.Direction == key.DirPress; down || e.Direction == key.DirRelease {
					game.Touch(down)
				}
			}
		}
	})
} // _screenMain() 
