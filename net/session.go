package net

import (
	"context"
	"github.com/lackone/go-tcp/utils"
	"net"
)

type Session struct {
	id   int64           //会话ID
	conn net.Conn        //底层连接
	ctx  context.Context //上下文
}

func (s *Session) GetId() int64 {
	return s.id
}

func (s *Session) GetConn() net.Conn {
	return s.conn
}

func (s *Session) Read(b []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Session) Write(b []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func NewSession(ctx context.Context, conn net.Conn) *Session {
	snowflake, _ := utils.NewSnowflake(0)
	id := snowflake.GetId()

	return &Session{
		id:   id,
		conn: conn,
		ctx:  ctx,
	}
}
