package sessionRWData

import (
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
)

func TestSessionReadWrite(t *testing.T) {
	// 定义监听IP和端口
	addr := "127.0.0.1:1024"
	// 定义传输的数据
	my_data := "hello"
	// 定义一个等待组,解决后面定义的两个协程同步问题
	wg := sync.WaitGroup{}
	// 两个协程，1个读取，1个写入。
	wg.Add(2)
	// 写数据的协程
	go func() {
		defer wg.Done()
		// 创建tcp连接
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		conn, _ := listen.Accept()
		s := Session{conn: conn}
		// 写数据
		s.Write([]byte(my_data))
	}()
	// 读数据的协程
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		s := Session{conn: conn}
		data, err := s.Read()
		if err != nil {
			t.Fatal(err)
		}
		// 这里读取的 data 结果就应该是 前面写数据传入的 my_data
		fmt.Printf("读取的结果是: %v\n", string(data))
	}()
	wg.Wait()
}
