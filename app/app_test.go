package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/simapp"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"

	"github.com/Entangle-Protocol/entangle-blockchain/encoding"
)

func TestEthermintAppExport(t *testing.T) {
	db := dbm.NewMemDB()
	app := SetupWithDB(false, nil, db)
	app.Commit()

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewEthermintApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, encoding.MakeConfig(ModuleBasics), simapp.EmptyAppOptions{}, true)
	_, err := app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
