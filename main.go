package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

//	一个mac状态栏的番茄时间小程序
func main() {
	runtime.LockOSThread()

	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)

		//	先统一创建组件的方式比分类创建要更优（考虑组件间可能有功能依赖）
		itemNext := cocoa.NSMenuItem_New()
		itemStop := cocoa.NSMenuItem_New()
		itemQuit := cocoa.NSMenuItem_New()
		menu := cocoa.NSMenu_New()

		obj.Retain()
		obj.Button().SetTitle("▶️ Ready")

		nextClicked := make(chan bool)
		countdown := false
		isStop := false

		//	计时程序
		go func() {
			timer := 1500 // 	默认时间值
			state := -1
			for {
				select {
				case <-time.After(1 * time.Second):
					if timer > 0 && countdown {
						timer = timer - 1
					}
					//	当处于工作或者休息时才会在结束时更新状态
					if timer <= 0 && state%2 == 1 {
						state = (state + 1) % 4
						itemStop.SetHidden(state%2 == 0)
					}
				case <-nextClicked:
					state = (state + 1) % 4
					if state%2 == 1 {
						countdown = true
					} else {
						countdown = false
					}
					timer = map[int]int{
						0: 0,
						1: 2700,
						2: 0,
						3: 300,
					}[state]
				}
				labels := map[int]string{
					0: "▶️ Ready 2 go ?",
					1: "✴️ Working %02d:%02d",
					2: "✅ Finished! Just take a break!",
					3: "⏸️ Break %02d:%02d",
				}
				if state%2 == 1 {
					obj.Button().SetTitle(fmt.Sprintf(labels[state], timer/60, timer%60))
					itemStop.SetHidden(false)
				} else {
					obj.Button().SetTitle(fmt.Sprintf(labels[state]))
					itemStop.SetHidden(true)
				}

			}
		}()
		//	触发更新，即将state设置0
		nextClicked <- true

		itemNext.SetTitle("Next")
		itemNext.SetAction(objc.Sel("nextClicked:"))
		cocoa.DefaultDelegateClass.AddMethod("nextClicked:", func(_ objc.Object) {
			nextClicked <- true
			isStop = false
			itemStop.SetTitle("Stop")
		})

		itemStop.SetTitle("Stop")
		itemStop.SetAction(objc.Sel("toStop:"))
		cocoa.DefaultDelegateClass.AddMethod("toStop:", func(_ objc.Object) {
			countdown, isStop = !countdown, !isStop
			switch isStop {
			case true:
				itemStop.SetTitle("Start")
			case false:
				itemStop.SetTitle("Stop")
			}
		})

		itemQuit.SetTitle("Quit")
		itemQuit.SetAction(objc.Sel("terminate:"))

		menu.AddItem(itemNext)
		menu.AddItem(itemStop)
		menu.AddItem(itemQuit)
		obj.SetMenu(menu)

	})
	app.Run()
}
