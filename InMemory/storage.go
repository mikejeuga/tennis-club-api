package InMemory

import (
	"context"
	"github.com/mikejeuga/tennis-club-api/tennisclub"
)

type Storage struct {
}

func (s *Storage) Create(ctx context.Context, player tennisclub.Player) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) FindByID(ctx context.Context, id string) (tennisclub.Player, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) DeleteByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
