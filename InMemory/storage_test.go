package InMemory_test

import (
	"context"
	"github.com/adamluzsi/testcase/random"
	"github.com/mikejeuga/tennis-club-api/InMemory"
	"github.com/mikejeuga/tennis-club-api/tennisclub"
	"github.com/mikejeuga/tennis-club-api/tennisclub/specifications"
	"testing"
	"time"
)

func TestStorage_implPlayerStorage(t *testing.T) {

	specifications.PlayerStorage{
		Subject: func(tb testing.TB) tennisclub.PlayerStorage {
			return &InMemory.Storage{}
		},
		MakerContext: func(tb testing.TB) context.Context {
			return context.Background()
		},
		MakePlayer: func(tb testing.TB) tennisclub.Player {
			rand := random.New(random.CryptoSeed{})
			tm := rand.TimeBetween(time.Now().AddDate(-100, 0, 0), time.Now())
			return tennisclub.Player{
				FirstName:   rand.String(),
				LastName:    rand.String(),
				Nationality: rand.String(),
				Birthday: tennisclub.Birthday{
					Day:   tm.Day(),
					Month: tm.Month(),
					Year:  tm.Year(),
				},
			}
		},
	}.Test(t)
}
