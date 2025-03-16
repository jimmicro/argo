# argo - Auto recover goroutine

argo 是一个轻量级的 Go panic 处理库，提供了优雅的方式来处理 goroutine 中的 panic 和错误。

## 特性

- 自动捕获并记录 goroutine 中的 panic
- 支持自定义 panic 处理器
- 提供详细的堆栈跟踪信息

## 安装

```bash
go get github.com/jimmicro/argo@latest
```

## 使用示例

```go
package main

import "github.com/jimmicro/argo"

func main() {
    // 使用 Go 函数启动一个安全的 goroutine
    rgo.Go(func() {
        // 这里的 panic 会被捕获并记录
        panic("test panic")
    })
}
 ```
