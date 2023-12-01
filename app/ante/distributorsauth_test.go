package ante_test

import (
	"math/big"

	"github.com/Entangle-Protocol/entangle-blockchain/app/ante"
	"github.com/Entangle-Protocol/entangle-blockchain/tests"
	distributorsauth "github.com/Entangle-Protocol/entangle-blockchain/x/distributorsauth/types"
	evmtypes "github.com/Entangle-Protocol/entangle-blockchain/x/evm/types"
)

func (suite AnteTestSuite) TestDeploySmartContracts() {
	suite.SetupTest()

	dec := ante.NewDistributorsAuthDecorator(suite.app.DistributorsAuthKeeper)

	addr, privKey := tests.NewAddrKey()

	signedTx := evmtypes.NewTxContract(suite.app.EvmKeeper.ChainID(), 1, big.NewInt(10), 1000, big.NewInt(1), nil, nil, nil, nil)
	signedTx.From = addr.Hex()
	err := signedTx.Sign(suite.ethSigner, tests.NewSigner(privKey))
	suite.Require().NoError(err)

	// Testing mode succes in all cases
	_, err = dec.AnteHandle(suite.ctx, signedTx, false, NextFn)
	suite.Require().NoError(err)

	_, err = dec.AnteHandle(suite.ctx, signedTx, true, NextFn)
	suite.Require().NoError(err)

	suite.app.DistributorsAuthKeeper.SetTestingMode(false)
	dec = ante.NewDistributorsAuthDecorator(suite.app.DistributorsAuthKeeper)

	_, err = dec.AnteHandle(suite.ctx, signedTx, true, NextFn)
	suite.Require().NoError(err)

	// Error for not distributor in not testing mode
	_, err = dec.AnteHandle(suite.ctx, signedTx, false, NextFn)
	suite.Require().Error(err)

	distr := distributorsauth.DistributorInfo{
		Address: addr.Hex(),
		EndDate: 12345678901234,
	}
	// Add distributor address
	suite.app.DistributorsAuthKeeper.AddDistributor(suite.ctx, distr)

	// Check is ok for correcrt distributor
	_, err = dec.AnteHandle(suite.ctx, signedTx, false, NextFn)
	suite.Require().NoError(err)

	_, err = dec.AnteHandle(suite.ctx, signedTx, false, NextFn)
	suite.Require().NoError(err)
}
