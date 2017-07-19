package actions_test

import (
	"testing"

	"github.com/gobuffalo/suite"
	"github.com/markbates/gobular/actions"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	as := &ActionSuite{suite.NewAction(actions.App())}
	suite.Run(t, as)
}
