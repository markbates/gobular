package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/gobular/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
