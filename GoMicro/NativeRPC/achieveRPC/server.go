package achieveRPC

import (
	"GoMicro/NativeRPC/sessionRWData"
	"fmt"
	"net"
	"reflect"
)

// 声明服RPC服务端
type Server struct {
	// 地址
	addr string
	// 服务端维护的函数名到函数反射值的map
	funcs map[string]reflect.Value
}

// 创建服务端对象
func NewServer(addr string) *Server {

	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}

}

/**
 * 服务端绑定注册方法（实现将函数名与函数真正实现对应起来）
 * 第一个参数是函数名
 * 第二个参数传入真正的函数，即函数体
 */
func (s *Server) Register(rpcName string, fn interface{}) {
	_, ok := s.funcs[rpcName]
	if ok {
		// Server.funcs 是否有 rpcName，如果有的话表示已经注册过了，直接返回即可。
		return
	} else {
		// 如果map中没有的话就将映射添加进map，便于调用。
		fnVal := reflect.ValueOf(fn)
		s.funcs[rpcName] = fnVal
	}
}

// 服务端等待调用
func (s *Server) Run() {
	// 监听
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Printf("监听 %s 错误： %v \n", s.addr, err)
		return
	}

	for {
		// 拿到连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept 错误： %v \n", err)
			return
		}
		// 创建会话
		session := sessionRWData.NewSession(conn)
		// RPC读取数据
		byt, err := session.Read()
		if err != nil {
			fmt.Printf("read 错误： %v \n", err)
			return
		}
		// 对数据解码
		rpcData, err := sessionRWData.Decode(byt)
		if err != nil {
			fmt.Printf("decode 错误： %v \n", err)
			return
		}
		// 根据读取到的数据的Name，得到调用的函数名.
		fn, ok := s.funcs[rpcData.Name]
		if !ok {
			// 如果不存在的话，就是有问题的。
			fmt.Printf("函数 %v 不存在： %v \n", rpcData.Name)
			return
		}
		// 函数存在就调用该函数，还要解析客户端传来的参数，参数是一个数组（RPCData.Args)，这里定义一个数组放参数。
		// 创建一个切片是反射值类型的，长度是0，容量是rpcData.Args
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			// 遍历每一项放到 inArgs 中去。
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		// 反射调用方法，传入参数
		out := fn.Call(inArgs)
		// 解析遍历执行结果，放到一个数组中。
		outArgs := make([]interface{}, 0, len(out))
		for _, oArg := range out {
			//
			outArgs = append(outArgs, oArg.Interface())
		}
		// 包装要返回给客户端的数据
		respRPCData := sessionRWData.RPCData{rpcData.Name, outArgs}
		// 编码
		respBytes, err := sessionRWData.Encode(respRPCData)
		if err != nil {
			fmt.Printf("encode 错误： %v \n", err)
			return
		}
		// 使用RPC写出数据
		err = session.Write(respBytes)
		if err != nil {
			fmt.Printf("session write 错误： %v \n", err)
			return
		}
	}

}
