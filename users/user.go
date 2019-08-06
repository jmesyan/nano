package users

type User struct {
	Uid            int    //用户ID
	Channel        string //通道
	ConnectorAddr  string //客户端地址
	GameServerAddr string //服务端地址
}
