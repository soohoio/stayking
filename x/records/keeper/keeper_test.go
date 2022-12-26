package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/soohoio/stayking/app/apptesting"
	"github.com/soohoio/stayking/x/records/keeper"
	"github.com/soohoio/stayking/x/records/types"
)

type KeeperTestSuite struct {
	apptesting.AppTestHelper
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
}

func (s *KeeperTestSuite) GetMsgServer() types.MsgServer {
	return keeper.NewMsgServerImpl(s.App.RecordsKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
