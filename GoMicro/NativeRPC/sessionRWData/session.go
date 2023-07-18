package sessionRWData

import (
	"encoding/binary"
	"io"
	"net"
)

// 会话中数据读写
// 会话连接的结构体
type Session struct {
	conn net.Conn
}

// 创建新的连接
func newSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

// 向连接中写数据
func (s *Session) Write(data []byte) error {
	// 4个字节头 + 数据长度的切片
	buf := make([]byte, 4+len(data))
	// 写入头部数据，记录数据长度
	// binary只认固定长度的类型，所以使用了uint32而不是直接写入。
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	// 写入数据
	copy(buf[:4], data)
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

// 从连接中读数据
func (s *Session) Read() ([]byte, error) {
	// 读取头部长度
	header := make([]byte, 4)
	// 按头部长度，读取头部数据
	_, err := io.ReadFull(s.conn, header)
	if err != nil {
		return nil, err
	}
	// 读取数据长度
	dataLen := binary.BigEndian.Uint32(header)
	// 根据数据长度读取数据
	data := make([]byte, dataLen)
	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
