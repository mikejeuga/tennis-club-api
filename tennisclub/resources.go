package tennisclub

import "context"

type EntityStorage[Entity, ID any] interface {
	Create(ctx context.Context, thingy Entity) (ID, error)
	FindAll(ctx context.Context) ([]Entity, error)
	FindByID(ctx context.Context, id ID) (Entity, bool, error)
	DeleteByID(ctx context.Context, id ID) error
}

type PlayerStorage interface {
	EntityStorage[Player, string]
}
