package app

import (
	context "golang.org/x/net/context"

	"github.com/chemidy/goheroe-code-examples/gRPC-Service-Discovery-With-Server-Reflection-in-go/proto"
)

var heroes = []*superpower.SuperPower{
	&superpower.SuperPower{Name: "Flash power", Cat: superpower.SuperPowerCategory_SciencePowers},
	&superpower.SuperPower{Name: "Superman power", Cat: superpower.SuperPowerCategory_CosmicBasedPowers},
	&superpower.SuperPower{Name: "Spiderman power", Cat: superpower.SuperPowerCategory_SupernaturalPowers},
}

// GoHeroeServer implements GoHeroeServer.
type GoHeroeServer struct {
}

// NewGoHeroeServer returns a new GoHeroeServer.
func NewGoHeroeServer() (*GoHeroeServer, error) {
	return &GoHeroeServer{}, nil
}

// List super powers with a filter
func (server GoHeroeServer) List(ctx context.Context, filter *superpower.Filter) (*superpower.SuperPowers, error) {

	if filter.Category == 0 {
		return &superpower.SuperPowers{SuperPow: heroes}, nil
	}
	return &superpower.SuperPowers{SuperPow: filterHeores(heroes, func(heroe *superpower.SuperPower) bool { return heroe.Cat == filter.Category })}, nil
}

// Add super power
func (server GoHeroeServer) Add(ctx context.Context, heroe *superpower.SuperPower) (*superpower.SuperPowers, error) {
	heroes = append(heroes, heroe)
	return &superpower.SuperPowers{SuperPow: heroes}, nil
}

type filterOnSuperPower func(*superpower.SuperPower) bool

// FilterHeores (slice, predicate func)
func filterHeores(in []*superpower.SuperPower, fn filterOnSuperPower) []*superpower.SuperPower {
	out := make([]*superpower.SuperPower, 0)
	for _, current := range in {
		if fn(current) {
			out = append(out, current)
		}
	}
	return out
}
