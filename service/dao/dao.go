package dao

import (
	"sync"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/XOS/Probe/model"
	pb "github.com/XOS/Probe/proto"
)

const (
	SnapshotDelay = 3
	ReportDelay   = 2
)

// Conf ..
var Conf *model.Config

// Cache ..
var Cache *cache.Cache

// DB ..
var DB *gorm.DB

// ServerList ..
var ServerList map[string]*model.Server

// ServerLock ..
var ServerLock sync.RWMutex

// Version ..
var Version = "debug"

func init() {
	if len(Version) > 7 {
		Version = Version[:7]
	}
}

// SendCommand ..
func SendCommand(cmd *pb.Command) {
	ServerLock.RLock()
	defer ServerLock.RUnlock()
	var err error
	for _, server := range ServerList {
		if server.Stream != nil {
			err = server.Stream.Send(cmd)
			if err != nil {
				close(server.StreamClose)
				server.Stream = nil
			}
		}
	}
}
