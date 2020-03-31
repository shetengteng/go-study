package rpcdemo // 如果是rpc会与标准库冲突
import "errors"

type DemoService struct{}

type Args struct {
	A, B int
}

// Div方法时除法功能
// rpc 对参数有要求，result一定要是指针类型
// 第一个参数是输入，第二个参数是输出
// 要求符合rpc的要求
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
