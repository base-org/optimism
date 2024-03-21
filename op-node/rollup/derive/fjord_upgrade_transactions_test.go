package derive

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
)

func TestFjordSourcesMatchSpec(t *testing.T) {
	for _, test := range []struct {
		source       UpgradeDepositSource
		expectedHash string
	}{
		{
			source:       deployFjordGasPriceOracleSource,
			expectedHash: "0x86122c533fdcb89b16d8713174625e44578a89751d96c098ec19ab40a51a8ea3",
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

func TestFjordNetworkTransactions(t *testing.T) {
	upgradeTxns, err := FjordNetworkUpgradeTransactions()
	require.NoError(t, err)
	require.Len(t, upgradeTxns, 3)

	deployGasPriceOracleSender, deployGasPriceOracle := toDepositTxn(t, upgradeTxns[0])
	require.Equal(t, deployGasPriceOracleSender, common.HexToAddress("0x4210000000000000000000000000000000000003"))
	require.Equal(t, deployFjordGasPriceOracleSource.SourceHash(), deployGasPriceOracle.SourceHash())
	require.Nil(t, deployGasPriceOracle.To())
	require.Equal(t, uint64(1_450_000), deployGasPriceOracle.Gas())
	require.Equal(t, bindings.GasPriceOracleMetaData.Bin, hexutil.Bytes(deployGasPriceOracle.Data()).String())

	updateGasPriceOracleSender, updateGasPriceOracle := toDepositTxn(t, upgradeTxns[1])
	require.Equal(t, updateGasPriceOracleSender, common.Address{})
	require.Equal(t, updateFjordGasPriceOracleSource.SourceHash(), updateGasPriceOracle.SourceHash())
	require.NotNil(t, updateGasPriceOracle.To())
	require.Equal(t, *updateGasPriceOracle.To(), common.HexToAddress("0x420000000000000000000000000000000000000F"))
	require.Equal(t, uint64(50_000), updateGasPriceOracle.Gas())
	require.Equal(t, common.FromHex("0x3659cfe6000000000000000000000000ff256497d61dcd71a9e9ff43967c13fde1f72d12"), updateGasPriceOracle.Data())

	gpoSetFjordSender, gpoSetFjord := toDepositTxn(t, upgradeTxns[2])
	require.Equal(t, gpoSetFjordSender, common.HexToAddress("0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001"))
	require.Equal(t, enableFjordSource.SourceHash(), gpoSetFjord.SourceHash())
	require.NotNil(t, gpoSetFjord.To())
	require.Equal(t, *gpoSetFjord.To(), common.HexToAddress("0x420000000000000000000000000000000000000F"))
	require.Equal(t, uint64(80_000), gpoSetFjord.Gas())
	require.Equal(t, common.FromHex("0x8e98b106"), gpoSetFjord.Data())
}
