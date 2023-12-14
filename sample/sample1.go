package main

import (
	"fmt"
	errors "github.com/wuntsong/wterrors"
)

var ErrClass1 = errors.NewClass("class 1")

func main() {
	_ = errors.New() // 创建错误
	_ = errors.New("简短Message")

	_ = errors.Errorf("支持printf格式：%d", 10)

	_ = errors.WarpQuick(fmt.Errorf("go内置错误")) // 快速封装go内置错误

	test := errors.New().Warp("包含一层原因")

	_ = test.Message()
	_ = test.Code()
	_ = test.Cause()
	_ = test.Class()
	_ = test.Stack()

	_ = ErrClass1.New() // class也能做上述操作
	_ = ErrClass1.Errorf("")
	_ = ErrClass1.WarpQuick(fmt.Errorf(""))

	_ = errors.Is(test, ErrClass1) // 可以比较class和err
}
