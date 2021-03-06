# :tomato: Mac状态栏番茄时间钟

原型： [progrium/macdriver/examples/pomodoro](https://github.com/progrium/macdriver/blob/main/examples/pomodoro/main.go#L1)  
在此基础上添加了暂停的功能

运行方法：将`main.go` 编译后直接运行即可，或者使用`go run main.go`即可运行

运行界面展示：

* [![准备](https://s3.ax1x.com/2021/03/05/6mkzsP.png)](https://imgtu.com/i/6mkzsP)
* [![6mAcOP.png](https://s3.ax1x.com/2021/03/05/6mAcOP.png)](https://imgtu.com/i/6mAcOP)

PS：在状态栏中提示信息较多时，程序容易**被折叠隐藏**

## update更新信息：
【2021-03-06】 修改番茄时间为标准的25分钟，完成4个番茄时间后进入25分钟的大休息状态。

## :construction: TODO-List：

1. 工作、休息阶段结束后进行小磁贴的通知
2. 程序运行过程中进行专注时长统计，程序退出前进行自动总结
3. 相关持久化操作，生成周月统计报表