package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "Autheo/testutil/keeper"
	"Autheo/x/autheo/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AutheoKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
