package achieveRPC

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

// 用于测试的结构体
type User struct {
	Name string
	Age  int
}

// 用于测试的查询用户的方法
func queryUser(userId int) (User, error) {
	user := make(map[int]User)
	user[0] = User{"trump", 30}
	user[1] = User{"clinton", 35}
	user[2] = User{"obama", 22}
	// 模拟查询用户
	u, ok := user[userId]
	if ok {
		return u, nil
	}
	return User{}, fmt.Errorf("id %d not in user db", userId)
}

// 测试方法
func TestRPC(t *testing.T) {
	// 需要对interface{}可能产生的类型进行注册
	gob.Register(User{})
	addr := "127.0.0.1:8080"
	// 创建服务端
	server := NewServer(addr)
	// 将方法注册到服务端
	server.Register("queryUser", queryUser)
	// 服务端等待调用
	go server.Run()
	// 客户端获取连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Error(err)
	}
	// 创建客户端
	client := NewClient(conn)
	// 声明函数原型
	var query func(int) (User, error)
	client.callRPC("queryUser", &query)
	// 得到查询结果
	user, err := query(1)
	if err != nil {
		t.Fatal(err)
	}
	// 这里query(1)得到的应该是：User{"clinton", 35}，如果正确，那么这个程序就没有问题。
	fmt.Printf("user: %v \n", user)
}
