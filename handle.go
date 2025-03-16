package argo

import (
	"log"
	"runtime"
)

// ReallyCrash 控制 HandleCrash 的行为，默认为 false
// 它被暴露出来以便组件可以选择将其设置为 false 以恢复之前的行为。
// 这个标志主要用于测试崩溃条件。
var ReallyCrash = false

type PanicHandlerFunc func(interface{})

func handleCrash() {
	if r := recover(); r != nil {
		for _, fn := range panicHandlers {
			fn(r)
		}
		if ReallyCrash {
			panic(r)
		}
	}
}

// AddPanicHandler 添加一个新的 panic 处理程序。
func AddPanicHandler(fn PanicHandlerFunc) {
	panicHandlers = append(panicHandlers, fn)
}

// panicHandlers 是一个函数列表，当发生 panic 时会被调用。
var panicHandlers = []PanicHandlerFunc{logPanic}

// logPanic 在发生 panic 时记录调用者树（除了 http.ErrAbortHandler 的特殊情况）。
func logPanic(r interface{}) {
	const size = 64 << 10
	stacktrace := make([]byte, size)
	stacktrace = stacktrace[:runtime.Stack(stacktrace, false)]

	if _, ok := r.(string); ok {
		log.Printf("[PANIC] %v\nStacktrace:\n%s", r, stacktrace)
	} else {
		log.Printf("[PANIC] %v (%#v)\nStacktrace:\n%s", r, r, stacktrace)
	}
}

// GetCaller 返回调用它的函数的调用者。
func GetCaller() string {
	var pc [1]uintptr
	runtime.Callers(3, pc[:])
	f := runtime.FuncForPC(pc[0])
	if f == nil {
		return "?"
	}
	return f.Name()
}
