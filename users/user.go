package users

type User struct {
	Sid           int64  `json:"sid"`            // session_id
	Uid           int    `json:"uid"`            //用户ID
	Channel       string `json:"channel"`        //通道
	ConnectorNid  string `json:"connector_nid"`  //客户端地址
	GameserverNid string `json:"gameserver_nid"` //服务端地址
}
