package main

import (
	"fmt"
	"github.com/jmesyan/nano/connectors"
	"github.com/jmesyan/nano/utils"
	"log"
	"net/http"
	"time"

	"github.com/jmesyan/nano"
	"github.com/jmesyan/nano/component"
	"github.com/jmesyan/nano/serialize/json"
	"github.com/jmesyan/nano/session"
	"strings"
)

type (
	Room struct {
	}

	// RoomManager represents a component that contains a bundle of room
	RoomManager struct {
		component.Base
		timer *utils.Timer
		rooms map[int]*Room
	}

	// UserMessage represents a message that user sent
	UserMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	// NewUser message will be received when new user join room
	NewUser struct {
		Content string `json:"content"`
	}

	// AllMembers contains all members uid
	AllMembers struct {
		Members []int64 `json:"members"`
	}

	// JoinResponse represents the result of joining room
	JoinResponse struct {
		Code   int    `json:"code"`
		Result string `json:"result"`
	}

	stats struct {
		component.Base
		timer         *utils.Timer
		outboundBytes int
		inboundBytes  int
	}
)

func (stats *stats) outbound(s *session.Session, msg nano.Message) error {
	stats.outboundBytes += len(msg.Data)
	return nil
}

func (stats *stats) inbound(s *session.Session, msg nano.Message) error {
	stats.inboundBytes += len(msg.Data)
	return nil
}

func (stats *stats) AfterInit() {
	stats.timer = utils.NewTimer(time.Minute, func() {
		println("OutboundBytes", stats.outboundBytes)
		println("InboundBytes", stats.outboundBytes)
	})
}

const (
	testRoomID = 1
	roomIDKey  = "ROOM_ID"
)

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: map[int]*Room{},
	}
}

// AfterInit component lifetime callback
func (mgr *RoomManager) AfterInit() {
	session.Lifetime.OnClosed(func(s *session.Session) {
		if !s.HasKey(roomIDKey) {
			return
		}
	})
	mgr.timer = utils.NewTimer(time.Minute, func() {
		for roomId, _ := range mgr.rooms {
			println(fmt.Sprintf("UserCount: RoomID=%d, Time=%s, Count=%d",
				roomId, time.Now().String(), assist.Count()))
		}
	})
}

var (
	assist = connectors.ConnectorHandler
)

// Join room
func (mgr *RoomManager) Join(s *session.Session, msg []byte) error {
	// NOTE: join test room only in demo
	room, found := mgr.rooms[testRoomID]
	if !found {
		room = &Room{
			//group: nano.NewGroup(fmt.Sprintf("room-%d", testRoomID)),
		}
		mgr.rooms[testRoomID] = room
	}

	fakeUID := s.ID() //just use s.ID as uid !!!
	s.Bind(fakeUID)   // binding session uids.Set(roomIDKey, room)
	s.Set(roomIDKey, room)
	s.Push("onMembers", &AllMembers{Members: assist.Members()})
	// notify others
	assist.Broadcast("onNewUser", &NewUser{Content: fmt.Sprintf("New user: %d", s.ID())})
	// new user join group
	assist.Add(s) // add session to group
	return s.Response(&JoinResponse{Result: "success"})
}

// Message sync last message to all members
func (mgr *RoomManager) Message(s *session.Session, msg *UserMessage) error {
	if !s.HasKey(roomIDKey) {
		return fmt.Errorf("not join room yet")
	}
	//room := s.Value(roomIDKey).(*Room)
	return nil
}

func main() {
	// override default serializer
	nano.SetSerializer(json.NewSerializer())

	// rewrite component and handler name
	room := NewRoomManager()
	nano.Register(room,
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower),
	)

	// traffic stats
	pipeline := nano.NewPipeline()
	var stats = &stats{}
	pipeline.Outbound().PushBack(stats.outbound)
	pipeline.Inbound().PushBack(stats.inbound)

	nano.EnableDebug()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	nano.SetWSPath("/nano")

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	nano.SetCheckOriginFunc(func(_ *http.Request) bool { return true })
	nano.ListenWS(":3250", nano.WithPipeline(pipeline))
}
