package argo

// Go 自动捕获崩溃并记录错误。通过 defer 调用。
func Go(fn func()) {
	go func() {
		defer handleCrash()
		fn()
	}()
}
