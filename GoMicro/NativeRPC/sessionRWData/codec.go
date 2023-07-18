package sessionRWData

import (
	"bytes"
	"encoding/gob"
)

// 定义RPC交互的数据格式
// 向rpcClient和rpcServer示例中那样 客户端想向服务端发起调用 首先需要知道调用的函数名和函数的参数。
// 所以这块这个结构体应该基本有访问的函数名和访问函数的参数，Args就给一个 []interface{} 任意类型，因为参数类型是不确定的
type RPCData struct {
	Name string
	Args []interface{}
}

// 编码
func Encode(data RPCData) ([]byte, error) {
	var buf bytes.Buffer
	// 得到字节数组的编码器
	bufEnc := gob.NewEncoder(&buf)

	// 对数据编码
	err := bufEnc.Encode(data)
	if err != nil {
		return nil, err
	}
	// 因为结果是直接数组（[]byte），所以不能直接返回 buf，应该转为直接数组返回 buf.Bytes()
	return buf.Bytes(), nil
}

// 解码
func Decode(b []byte) (RPCData, error) {
	buf := bytes.NewBuffer(b)
	// 返回字节数组的解码器
	bufDec := gob.NewDecoder(buf)
	var data RPCData
	// 对数据解码
	err := bufDec.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
