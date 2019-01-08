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
	"golang.org/x/mobile/exp/gl/glutil"
	//"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/gl"
)

var (
	_startTime = time.Now()
	_glImageS    *glutil.Images
	_GamE      *Game
)

func _screenMain() {
	rand.Seed(time.Now().UnixNano())

	app.Main(func(___a7 app.App) {
		var __glCtx7 gl.Context
		var __sz7 size.Event
		for __e7 := range ___a7.Events() {
			switch __e8 := ___a7.Filter(__e7).(type) {
			case lifecycle.Event:
				switch __e8.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					__glCtx7, _ = __e8.DrawContext.(gl.Context)
					onStart(__glCtx7)
					___a7.Send(paint.Event{})
				case lifecycle.CrossOff:
					onStop()
					__glCtx7 = nil
				}
				// switch __e8.Crosses(lifecycle.StageFocused) {
				// case lifecycle.CrossOn:
				// 	__glCtx7, _ = __e8.DrawContext.(gl.Context)
				// 	onStart(__glCtx7)
				// 	___a7.Send(paint.Event{})
				// case lifecycle.CrossOff:
				// 	__glCtx7.ClearColor(1, 1, 1, 1)
				// 	__glCtx7.Clear(gl.COLOR_BUFFER_BIT)
				// 	___a7.Publish()
				// 	onStop()
				// 	__glCtx7 = nil
				// }
			case size.Event:
				__sz7 = __e8
			case paint.Event:
				if __glCtx7 == nil || __e8.External {
					continue
				}
				_screenOnPaint(__glCtx7, __sz7)
				___a7.Publish()
				___a7.Send(paint.Event{}) // keep animating
			case touch.Event:
				if __down7 := __e8.Type == touch.TypeBegin; __down7 || __e8.Type == touch.TypeEnd {
					_GamE.Touch(__down7)
				}
			case key.Event:
				if __e8.Code != key.CodeSpacebar {
					break
				}
				if __down7 := __e8.Direction == key.DirPress; __down7 || __e8.Direction == key.DirRelease {
					_GamE.Touch(__down7)
				}
			}
		}
	})
} // _screenMain() 
