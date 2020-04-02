package grifts

import (
	"github.com/larrymjordan/tasks/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
