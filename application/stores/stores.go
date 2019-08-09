package stores

import "github.com/jmesyan/nano/application/models/structure"

var (
	StoresHandler *StoreDatas
)

type StoreSys struct {
	SYS_MAINTENANCE   bool
	MAINTENANCE_TIME  int64
	MAINTENANCE_TIME2 int
}

type StoreGds struct {
	Configs map[int]*structure.GameGoldsType
}
type StoreDatas struct {
	Sys *StoreSys
	Gds *StoreGds
}

func NewStoreDatas() *StoreDatas {
	return &StoreDatas{
		Sys: &StoreSys{},
		Gds: &StoreGds{
			Configs: make(map[int]*structure.GameGoldsType),
		},
	}
}

func init() {
	StoresHandler = NewStoreDatas()
}
