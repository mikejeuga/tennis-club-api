package tennisclub_test

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/tennis-club-api/tennisclub"
	"testing"
)

func TestRegistrar_Register(t *testing.T) {
	s := testcase.NewSpec(t)
	s.NoSideEffect()

	var (
		registrar = testcase.Let(s, func(t *testcase.T) *tennisclub.Registrar {
			return &tennisclub.Registrar{
				PlayerStorage: nil,
			}
		})
		ctx = testcase.Let(s, func(t *testcase.T) context.Context {
			return context.Background()
		})
		player = testcase.Let[tennisclub.Player](s, nil)
	)
	act := func(t *testcase.T) error {
		return registrar.Get(t).Register(ctx.Get(t), player.Get(t))
	}

	s.When("The player has first name and last name, nationality and the date of birth", func(s *testcase.Spec) {
		player.Let(s, func(t *testcase.T) tennisclub.Player {
			bday := t.Random.Time()
			return tennisclub.Player{
				FirstName:   t.Random.String(),
				LastName:    t.Random.String(),
				Nationality: t.Random.String(),
				Birthday: tennisclub.Birthday{
					Day:   bday.Day(),
					Month: bday.Month(),
					Year:  bday.Year(),
				},
			}
		})

		s.Then("The player is registered with the club", func(t *testcase.T) {
			t.Must.NoError(act(t))
		})

		s.And("Someone is already registered with the same firstname and lastname", func(s *testcase.Spec) {
			twinPlayer := testcase.Let[tennisclub.Player](s, func(t *testcase.T) tennisclub.Player {
				bday := t.Random.Time()
				return tennisclub.Player{
					FirstName:   player.Get(t).FirstName,
					LastName:    player.Get(t).LastName,
					Nationality: t.Random.String(),
					Birthday: tennisclub.Birthday{
						Day:   bday.Day(),
						Month: bday.Month(),
						Year:  bday.Year(),
					},
				}
			})

			s.Before(func(t *testcase.T) {
				t.Must.NoError(registrar.Get(t).Register(context.Background(), twinPlayer.Get(t)))
			})

			s.Then("The player cannot be registered", func(t *testcase.T) {
				err := act(t)
				t.Must.NotNil(err)
			})
		})
	})

}
