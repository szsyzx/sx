package model

import (
	"fmt"
	"html/template"
	"time"

	pb "github.com/XOS/Probe/proto"
)

// Server ..
type Server struct {
	Common
	Name         string
	DisplayIndex int    // 展示权重，越大越靠前
	Secret       string `json:"-"`

	Host       *Host  `gorm:"-"`
	State      *State `gorm:"-"`
	LastActive time.Time

	Stream      pb.ProbeService_HeartbeatServer `gorm:"-" json:"-"`
	StreamClose chan<- error                    `gorm:"-" json:"-"`
}

func (s Server) Marshal() template.JS {
	return template.JS(fmt.Sprintf(`{"ID":%d,"Name":"%s","Secret":"%s"}`, s.ID, s.Name, s.Secret))
}
