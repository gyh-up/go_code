package main

import (
	"fmt"
)
// 我们可以使用 「errors.New」 这个工厂函数来生成错误信息
func main() {
	res ,err := test(2,1)
	// 通过error断言来获取返回的错误信息，断言可以将error接口转为自己定义的错误类型：
	if e,ok := err.(*testError);ok {
		fmt.Println("错误码：",e.errorCode,",错误信息：",e.errorMsg)
	} else {
		fmt.Println(res)
	}
}
func test(m,n int) (int, error) {
	if m > n {
		return m, &testError {
			errorCode: 1,
			errorMsg:  "m大于n",
		}
	} else {
		return n, nil
	}
}
// 上面工厂函数只能传递一个字符串来返回，要想携带更多信息，这时候可以使用自定义error:
type testError struct {
	errorCode int //错误码
	errorMsg string //错误信息
}
func (t *testError) Error() string {
	return t.errorMsg
}