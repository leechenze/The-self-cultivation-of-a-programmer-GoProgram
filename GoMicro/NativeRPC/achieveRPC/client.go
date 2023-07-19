package achieveRPC

import (
	"GoMicro/NativeRPC/sessionRWData"
	"net"
	"reflect"
)

// 声明客户端
type Client struct {
	conn net.Conn
}

// 创建客户端对象
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

// 实现通用的RPC客户端
// 绑定RPC访问的方法
// 传入访问的函数名

// 函数具体实现在客户端，Client只有函数原型
// 使用makeFunc()完成原型到函数的调用

// fPtr是指向函数原型，调用方式如下：
// xxx.callRPC("queryUser", &query)
func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	// 通过反射，获取fPtr为初始化的函数原型
	fn := reflect.ValueOf(fPtr).Elem()
	// 定义另一个函数，作用是对第一个函数参数操作
	// 完成与Server的交互
	f := func(args []reflect.Value) []reflect.Value {
		// 处理输入的参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		// 创建连接
		clientSession := sessionRWData.NewSession(c.conn)
		// 数据编码
		requestRPC := sessionRWData.RPCData{Name: rpcName, Args: inArgs}
		byt, err := sessionRWData.Encode(requestRPC)
		if err != nil {
			panic(err)
		}
		// 写出数据
		err = clientSession.Write(byt)
		if err != nil {
			panic(err)
		}
		// 读取响应数据
		responseBytes, err := clientSession.Read()
		if err != nil {
			panic(err)
		}
		// 数据解码
		responseRPC, err := sessionRWData.Decode(responseBytes)
		if err != nil {
			panic(err)
		}
		// 处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(responseRPC.Args))
		for i, arg := range responseRPC.Args {
			// 必须进行nil转换
			if arg == nil {
				// 这一步表示 必须填充一个真正的类型不能是nil
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}

	// 参数1：一个未初始化函数的方法值，类型是reflect.Type
	// 参数2：另一个函数，作用是对第一个函数参数操作
	// 返回 reflect.Value 类型
	// makeFunc使用传入的函数原型，创建一个绑定 参数2 的新函数
	val := reflect.MakeFunc(fn.Type(), f)
	// 为函数fPtr赋值
	fn.Set(val)
}
