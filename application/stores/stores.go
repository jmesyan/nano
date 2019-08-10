package stores

import (
	"fmt"
	"github.com/jmesyan/nano/application/cache"
	"github.com/jmesyan/nano/application/models/structure"
	"github.com/jmesyan/nano/utils"
	"math"
	"math/rand"
)

var (
	StoresHandler *StoreDatas
)

type StoreSys struct {
	SYS_MAINTENANCE   bool
	MAINTENANCE_TIME  int64
	MAINTENANCE_TIME2 int
	MAINTEN_SERVERS   map[string]bool
}

type StoreGds struct {
	Configs map[int]*structure.GameGoldsType
	Gcsu    map[string]int                 //中心用户数量
	Gucs    map[int]string                 //用户报名中心服务器
	Gcsp    map[int]map[string]interface{} //用户中心服务器数据
}

type Temptable struct {
	T    int
	Code int32
}

type StoreMj struct {
	Use       []string
	Nouse     map[int][]string
	Codelen   int
	Code      map[int32]string
	CodeSort  map[string]int32
	Gsidtid   map[string][]int
	AllCode   []int32
	Temptable map[string]*Temptable
}

type StoreP2p struct {
	Mj *StoreMj
}

func (p2p *StoreP2p) initCode() {
	p2p.Mj.AllCode = []int32{}
	begin := 1
	for i := 1; i < p2p.Mj.Codelen; i++ {
		begin = begin * 10
	}
	end := begin * 10
	for j := begin; j < end; j++ {
		p2p.Mj.AllCode = append(p2p.Mj.AllCode, int32(j))
	}
	array := p2p.Mj.AllCode
	m, i := len(array), 0
	for k := m; k > 0; {
		i = int(math.Floor(float64(rand.Intn(k))))
		k -= 1
		array[k], array[i] = array[i], array[k]
	}
	fmt.Println("mj init code:", begin, end-1, m)
}

type StoreDatas struct {
	Sys *StoreSys
	Gds *StoreGds
	P2p *StoreP2p
}

func (sd *StoreDatas) initSys() {
	sd.Sys = &StoreSys{}
	sys := sd.Sys
	sys.SYS_MAINTENANCE = true
	sys.MAINTEN_SERVERS = make(map[string]bool)
	//维护开关
	maintence := cache.CacheManager.GetMaintence()
	if maintence != nil {
		if maintence.Type == 1 {
			sys.SYS_MAINTENANCE = true
		}
		time := utils.Time()
		if maintence.Type == 1 && time < maintence.Time {
			sys.MAINTENANCE_TIME2 = maintence.Time
		}
	}
}

func (sd *StoreDatas) initGds() {
	sd.Gds = &StoreGds{
		Configs: cache.CacheManager.GetGameGoldsType(),
		Gcsu:    make(map[string]int),
		Gucs:    make(map[int]string),
		Gcsp:    make(map[int]map[string]interface{}),
	}
}

func (sd *StoreDatas) initP2p() {
	sd.P2p = &StoreP2p{
		Mj: &StoreMj{
			Use:       nil,
			Nouse:     make(map[int][]string),
			Codelen:   6,
			Code:      nil,
			CodeSort:  nil,
			AllCode:   nil,
			Temptable: nil,
			Gsidtid:   make(map[string][]int),
		},
	}
	sd.P2p.initCode()
	//临时桌子10分钟回收

}

func NewStoreDatas() *StoreDatas {
	sd := &StoreDatas{}
	sd.initSys()
	sd.initGds()
	sd.initP2p()
	return sd
}

func init() {
	StoresHandler = NewStoreDatas()
}
