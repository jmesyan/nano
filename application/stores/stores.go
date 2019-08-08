package stores

var (
	StoresHandler *StoreDatas
)

type StoreSys struct {
	SYS_MAINTENANCE   bool
	MAINTENANCE_TIME  int
	MAINTENANCE_TIME2 int
}
type StoreDatas struct {
	Sys *StoreSys
}
