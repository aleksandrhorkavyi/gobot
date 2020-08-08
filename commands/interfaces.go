package commands

import (
	"gobot/models"
)

type Command interface {
	Execute(u models.Update) (bool, error)
}
