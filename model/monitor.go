package model

import (
	pb "github.com/XOS/Probe/proto"
)

const (
	_ = iota
	TaskTypeHTTPGET
	TaskTypeICMPPing
	TaskTypeTCPPing
	TaskTypeCommand
)

type Monitor struct {
	Common
	Name   string
	Type   uint8
	Target string
}

const defaultUserAgent = "Mozilla/5.0 (compatible; NGBot/2.1; +https://server.nange.cn/)"

func (m *Monitor) PB() *pb.Task {
	return &pb.Task{
		Id:   m.ID,
		Type: uint64(m.Type),
		Data: m.Target,
	}
}
