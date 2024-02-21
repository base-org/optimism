package derive

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestFjordSourcesMatchSpec(t *testing.T) {
	for _, test := range []struct {
		source       UpgradeDepositSource
		expectedHash string
	}{
		{
			source:       deployFjordL1BlockSource,
			expectedHash: "0x402f75bf100f605f36c2e2b8d5544a483159e26f467a9a555c87c125e7ab09f3",
		},
		{
			source:       deployFjordGasPriceOracleSource,
			expectedHash: "0x86122c533fdcb89b16d8713174625e44578a89751d96c098ec19ab40a51a8ea3",
		},
		{
			source:       updateFjordL1BlockProxySource,
			expectedHash: "0x0fefb8cb7f44b866e21a59f647424cee3096de3475e252eb3b79fa3f733cee2d",
		},
		{
			source:       updateFjordGasPriceOracleSource,
			expectedHash: "0x1e6bb0c28bfab3dc9b36ffb0f721f00d6937f33577606325692db0965a7d58c6",
		},
		{
			source:       enableFjordSource,
			expectedHash: "0xbac7bb0d5961cad209a345408b0280a0d4686b1b20665e1b0f9cdafd73b19b6b",
		},
	} {
		require.Equal(t, common.HexToHash(test.expectedHash), test.source.SourceHash())
	}
}

// TODO: Add unit tests around the Fjord upgrade transactions to make sure they match the spec.
