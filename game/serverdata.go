package game

type ServerData struct {
	Gid        int
	Rtype      int
	Ridx       int
	Tid        int
	Gsid       string
	Gsidtid    string
	State      int
	Initscore  int
	Motor      int
	Lid        int
	Code       int
	Quick      int
	QuickSit   int
	Scorescale float64
	Gobj       map[string]interface{}
}
