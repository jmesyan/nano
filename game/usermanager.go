package game

type UserManager struct {
	users map[int]*GamePlayer
}

type UserManagerOpt func(um *UserManager)

func NewUserManager(opts ...UserManagerOpt) *UserManager {
	um := &UserManager{}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(um)
		}
	}
	return um
}

func (um *UserManager) AddUserSort() {

}
