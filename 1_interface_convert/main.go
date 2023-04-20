package main
// 当一个函数的形参是 1_interface{} 时，意味着这个参数被自动的转为interface{} 类型，
//在函数中，如果想得到参数的真实类型，就需要对形参进行断言。
//类型断言就是将接口类型的值x，转换成类型T，格式为：x.(T)
//类型断言x必须为接口类型
//T可以是非接口类型，若想断言合法，则T必须实现x的接口
import "fmt"

func whoAmi(a interface{}) {
	//1.不断言
	//程序报错：cannot convert a (type 1_interface{}) to type string: need type assertion
	//fmt.Println(string(a))

	//2.非安全类型断言
	//fmt.Println(a.(string)) //无尘

	//3.安全类型断言
	value, ok := a.(string) //安全，断言失败，也不会panic，只是ok的值为false
	if !ok {
	fmt.Println("断言失败")
	return
	}
	fmt.Println(value)  //无尘
}
// 断言还有一种形式，就是使用「switch语句」判断接口的类型：
func whoAmi2(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Printf("boolean: %t\n", a) // a has type bool
	case int:
		fmt.Printf("integer: %d\n", a) // a has type int
	case string:
		fmt.Printf("string: %s\n", a) // a has type string
	default:
		fmt.Printf("unexpected type %T", a) // %T prints whatever type a has
	}
}
// 反射
//https://mp.weixin.qq.com/s?__biz=Mzk0NzI3Mjk1Mg==&mid=2247484506&idx=1&sn=bbce6d455af7ee6ea8be738abdb18dba&chksm=c37829cdf40fa0dbb0a1b087d8bb789087c7881e1cd3a00ee8b99a94f1d72d72d1b81a774478&cur_album_id=2001814334884085762&scene=189#wechat_redirect

func main() {
	str := "guo"
	whoAmi(str)
}
