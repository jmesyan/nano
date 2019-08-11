package nano

type ChannelStatus int
const (
	_ ChannelStatus = iota
	ChannelCreating
	ChannelCreated
)
type GameChannel struct {
	Uid int
	ClientNid string
	GameNid string
	Status ChannelStatus
}

type GameChannelOpt func(gc *GameChannel)

func NewGameChannel(uid int, clientNid string, opts ...GameChannelOpt)*GameChannel{
	gc := &GameChannel{
		Uid:uid,
		ClientNid:clientNid,
		Status:ChannelCreating,
	}
	if len(opts) > 0 {
		for _, opt:=range opts{
			opt(gc)
		}
	}
	return gc
}

func (gc *GameChannel) SetGameNid(gameNid string) {
	gc.GameNid = gameNid
	gc.Status = ChannelCreated
}


func (gc *GameChannel)C2S(cmd, msg string){

}

func (gc *GameChannel)S2C(heart int, cmd, msg string){

}



