package autheo_test

import (
	"testing"

	keepertest "Autheo/testutil/keeper"
	"Autheo/testutil/nullify"
	autheo "Autheo/x/autheo/module"
	"Autheo/x/autheo/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AutheoKeeper(t)
	autheo.InitGenesis(ctx, k, genesisState)
	got := autheo.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
