package specifications

import (
	"context"
	"github.com/adamluzsi/testcase/assert"
	"github.com/mikejeuga/tennis-club-api/tennisclub"
	"testing"
)

type PlayerStorage struct {
	Subject      func(tb testing.TB) tennisclub.PlayerStorage
	MakerContext func(tb testing.TB) context.Context
	MakePlayer   func(tb testing.TB) tennisclub.Player
}

func (c PlayerStorage) Test(t *testing.T) {
	t.Run("smoke test", func(t *testing.T) {
		storage := c.Subject(t)
		player := c.MakePlayer(t)
		ctx := c.MakerContext(t)
		_, found, err := storage.FindByID(ctx, "thePlayerID")
		assert.NoError(t, err)
		assert.False(t, found)

		registeredPlayerID, err := storage.Create(ctx, player)
		assert.NoError(t, err)
		player.ID = registeredPlayerID

		registeredPlayer, Found, err := storage.FindByID(ctx, registeredPlayerID)
		assert.NoError(t, err)
		assert.True(t, Found)
		assert.Equal(t, player, registeredPlayer)

		err = storage.DeleteByID(ctx, player.ID)
		assert.NoError(t, err)

		_, found, err = storage.FindByID(ctx, player.ID)
		assert.NoError(t, err)
		assert.False(t, found)

	})

}
