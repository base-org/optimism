package derive

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

func TestFjordNetworkTransactions(t *testing.T) {
	upgradeTxns, err := FjordNetworkUpgradeTransactions()
	require.NoError(t, err)
	require.Len(t, upgradeTxns, 5)

	deployL1BlockSender, deployL1Block := toDepositTxn(t, upgradeTxns[0])
	require.Equal(t, common.HexToAddress("0x4210000000000000000000000000000000000002"), deployL1BlockSender)
	require.Equal(t, deployFjordL1BlockSource.SourceHash(), deployL1Block.SourceHash())
	require.Nil(t, deployL1Block.To())
	require.Equal(t, uint64(420_000), deployL1Block.Gas())
	require.Equal(t, bindings.L1BlockMetaData.Bin, hexutil.Bytes(deployL1Block.Data()).String())

	deployGasPriceOracleSender, deployGasPriceOracle := toDepositTxn(t, upgradeTxns[1])
	require.Equal(t, deployGasPriceOracleSender, common.HexToAddress("0x4210000000000000000000000000000000000003"))
	require.Equal(t, deployFjordGasPriceOracleSource.SourceHash(), deployGasPriceOracle.SourceHash())
	require.Nil(t, deployGasPriceOracle.To())
	require.Equal(t, uint64(1_450_000), deployGasPriceOracle.Gas())
	require.Equal(t, bindings.GasPriceOracleMetaData.Bin, hexutil.Bytes(deployGasPriceOracle.Data()).String())

	updateL1BlockProxySender, updateL1BlockProxy := toDepositTxn(t, upgradeTxns[2])
	require.Equal(t, updateL1BlockProxySender, common.Address{})
	require.Equal(t, updateFjordL1BlockProxySource.SourceHash(), updateL1BlockProxy.SourceHash())
	require.NotNil(t, updateL1BlockProxy.To())
	require.Equal(t, *updateL1BlockProxy.To(), common.HexToAddress("0x4200000000000000000000000000000000000015"))
	require.Equal(t, uint64(50_000), updateL1BlockProxy.Gas())
	require.Equal(t, common.FromHex("0x3659cfe6000000000000000000000000a919894851548179a0750865e7974da599c0fac7"), updateL1BlockProxy.Data())

	updateGasPriceOracleSender, updateGasPriceOracle := toDepositTxn(t, upgradeTxns[3])
	require.Equal(t, updateGasPriceOracleSender, common.Address{})
	require.Equal(t, updateFjordGasPriceOracleSource.SourceHash(), updateGasPriceOracle.SourceHash())
	require.NotNil(t, updateGasPriceOracle.To())
	require.Equal(t, *updateGasPriceOracle.To(), common.HexToAddress("0x420000000000000000000000000000000000000F"))
	require.Equal(t, uint64(50_000), updateGasPriceOracle.Gas())
	require.Equal(t, common.FromHex("0x3659cfe6000000000000000000000000ff256497d61dcd71a9e9ff43967c13fde1f72d12"), updateGasPriceOracle.Data())

	gpoSetFjordSender, gpoSetFjord := toDepositTxn(t, upgradeTxns[4])
	require.Equal(t, gpoSetFjordSender, common.HexToAddress("0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001"))
	require.Equal(t, enableFjordSource.SourceHash(), gpoSetFjord.SourceHash())
	require.NotNil(t, gpoSetFjord.To())
	require.Equal(t, *gpoSetFjord.To(), common.HexToAddress("0x420000000000000000000000000000000000000F"))
	require.Equal(t, uint64(80_000), gpoSetFjord.Gas())
	require.Equal(t, common.FromHex("0x8e98b106"), gpoSetFjord.Data())
}
