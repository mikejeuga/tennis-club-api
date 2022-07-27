package tennisclub

import (
	"context"
)

type Registrar struct {
	PlayerStorage
}

func (r *Registrar) Register(ctx context.Context, player Player) error {
	//TODO implement me
	panic("implement me")
}
