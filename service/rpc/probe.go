package rpc

import (
	"context"
	"log"
	"time"

	"github.com/XOS/Probe/model"
	pb "github.com/XOS/Probe/proto"
	"github.com/XOS/Probe/service/dao"
)

// NezhaHandler ..
type ProbeHandler struct {
	Auth *AuthHandler
}

// ReportState ..
func (s *ProbeHandler) ReportState(c context.Context, r *pb.State) (*pb.Receipt, error) {
	var clientID uint64
	var err error
	if clientID, err = s.Auth.Check(c); err != nil {
		return nil, err
	}
	state := model.PB2State(r)
	dao.ServerLock.RLock()
	defer dao.ServerLock.RUnlock()
	dao.ServerList[clientID].LastActive = time.Now()
	dao.ServerList[clientID].State = &state
	return &pb.Receipt{Proced: true}, nil
}

// Heartbeat ..
func (s *ProbeHandler) Heartbeat(r *pb.Beat, stream pb.ProbeService_HeartbeatServer) error {
	var clientID uint64
	var err error
	defer log.Printf("Heartbeat exit server:%v err:%v", clientID, err)
	if clientID, err = s.Auth.Check(stream.Context()); err != nil {
		return err
	}
	// 放入在线服务器列表
	dao.ServerLock.RLock()
	closeCh := make(chan error)
	dao.ServerList[clientID].StreamClose = closeCh
	dao.ServerList[clientID].Stream = stream
	dao.ServerLock.RUnlock()
	select {
	case err = <-closeCh:
		return err
	}
}

// Register ..
func (s *ProbeHandler) Register(c context.Context, r *pb.Host) (*pb.Receipt, error) {
	var clientID uint64
	var err error
	if clientID, err = s.Auth.Check(c); err != nil {
		return nil, err
	}
	host := model.PB2Host(r)
	dao.ServerLock.RLock()
	defer dao.ServerLock.RUnlock()
	dao.ServerList[clientID].Host = &host
	return &pb.Receipt{Proced: true}, nil
}
