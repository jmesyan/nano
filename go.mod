module github.com/jmesyan/nano

go 1.12

replace (
	golang.org/x/net v0.0.0-20180724234803-3673e40ba225 => github.com/golang/net v0.0.0-20180724234803-3673e40ba225
	golang.org/x/net v0.0.0-20180826012351-8a410e7b638d => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd
	golang.org/x/net v0.0.0-20181023162649-9b4f9f5ad519 => github.com/golang/net v0.0.0-20181023162649-9b4f9f5ad519
	golang.org/x/net v0.0.0-20181201002055-351d144fa1fc => github.com/golang/net v0.0.0-20181201002055-351d144fa1fc
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd => github.com/golang/net v0.0.0-20190213061140-3a22650c66bd
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 => github.com/golang/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/net v0.0.0-20190509222800-a4d6f7feada5 => github.com/golang/net v0.0.0-20190509222800-a4d6f7feada5
)

replace (
	golang.org/x/sys v0.0.0-20180830151530-49385e6e1522 => github.com/golang/sys v0.0.0-20180830151530-49385e6e1522
	golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e => github.com/golang/sys v0.0.0-20180909124046-d0be0721c37e
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190215142949-d0b11bdaac8a
	golang.org/x/sys v0.0.0-20190412213103-97732733099d => github.com/golang/sys v0.0.0-20190412213103-97732733099d
	golang.org/x/sys v0.0.0-20190422165155-953cdadca894 => github.com/golang/sys v0.0.0-20190422165155-953cdadca894
	golang.org/x/sys v0.0.0-20190425145619-16072639606e => github.com/golang/sys v0.0.0-20190425145619-16072639606e
	golang.org/x/sys v0.0.0-20190509141414-a5b02f93d862 => github.com/golang/sys v0.0.0-20190509141414-a5b02f93d862
	golang.org/x/sys v0.0.0-20190710143415-6ec70d6a5542 => github.com/golang/sys v0.0.0-20190710143415-6ec70d6a5542
)

replace (
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sync v0.0.0-20180823144017-11551d06cbcc => github.com/golang/sync v0.0.0-20180823144017-11551d06cbcc
	golang.org/x/sync v0.0.0-20181026203630-95b1ffbd15a5 => github.com/golang/sync v0.0.0-20181026203630-95b1ffbd15a5
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4 => github.com/golang/sync v0.0.0-20181221193216-37e7f081c4d4
	golang.org/x/sync v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sync v0.0.0-20190215142949-d0b11bdaac8a
)

replace (
	golang.org/x/sys v0.0.0-20180823144017-11551d06cbcc => github.com/golang/sys v0.0.0-20180823144017-11551d06cbcc
	golang.org/x/sys v0.0.0-20181026203630-95b1ffbd15a5 => github.com/golang/sys v0.0.0-20181026203630-95b1ffbd15a5
)

replace golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0

replace (
	golang.org/x/crypto v0.0.0-20181029021203-45a5f77698d3 => github.com/golang/crypto v0.0.0-20181029021203-45a5f77698d3
	golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 => github.com/golang/crypto v0.0.0-20190123085648-057139ce5d2b
	golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9 => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/crypto v0.0.0-20190122013713-64072686203f => github.com/golang/crypto v0.0.0-20190122013713-64072686203f
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
)

require (
	github.com/go-redis/redis v0.0.0-20190803144825-742f3ccb21cd
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/consul/api v1.1.0
	github.com/nats-io/nats.go v1.8.1
	github.com/sirupsen/logrus v1.4.2
)
