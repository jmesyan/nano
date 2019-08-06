package users

type User struct {
	Sid            int64  `json:"sid"`             // session_id
	Uid            int    `json:"uid"`             //用户ID
	Channel        string `json:"channel"`         //通道
	ConnectorAddr  string `json:"connector_addr"`  //客户端地址
	GameserverAddr string `json:"gameserver_addr"` //服务端地址
}
