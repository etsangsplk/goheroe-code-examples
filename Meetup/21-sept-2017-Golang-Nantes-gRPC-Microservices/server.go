// GoHeroeServer implements GoHeroeServer.
type GoHeroeServer struct {
}

// NewGoHeroeServer returns a new GoHeroeServer.
func NewGoHeroeServer() (*GoHeroeServer, error) {
	return &GoHeroeServer{}, nil
}

// List super powers with a filter
func (server GoHeroeServer) List(ctx context.Context, filter *superpower.Filter) 
(*superpower.SuperPowers, error) {

	if filter.Category == 0 {
		return &superpower.SuperPowers{SuperPow: heroes}, nil
	}
	return &superpower.SuperPowers{SuperPow: filterHeores(heroes, 
		func(heroe *superpower.SuperPower) bool { return heroe.Cat == filter.Category })}, nil
}