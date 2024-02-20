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

/*
func toDepositTxn(t *testing.T, data hexutil.Bytes) (common.Address, *types.Transaction) {
	txn := new(types.Transaction)
	err := txn.UnmarshalBinary(data)
	require.NoError(t, err)
	require.Truef(t, txn.IsDepositTx(), "expected deposit txn, got %v", txn.Type())
	require.False(t, txn.IsSystemTx())

	signer := types.NewLondonSigner(big.NewInt(420))
	from, err := signer.Sender(txn)
	require.NoError(t, err)

	return from, txn
}

func TestEcotoneNetworkTransactions(t *testing.T) {
	upgradeTxns, err := EcotoneNetworkUpgradeTransactions()
	require.NoError(t, err)
	require.Len(t, upgradeTxns, 6)

	deployL1BlockSender, deployL1Block := toDepositTxn(t, upgradeTxns[0])
	require.Equal(t, deployL1BlockSender, common.HexToAddress("0x4210000000000000000000000000000000000000"))
	require.Equal(t, deployL1BlockSource.SourceHash(), deployL1Block.SourceHash())
	require.Nil(t, deployL1Block.To())
	require.Equal(t, uint64(375_000), deployL1Block.Gas())
	require.Equal(t, bindings.L1BlockMetaData.Bin, hexutil.Bytes(deployL1Block.Data()).String())

	deployGasPriceOracleSender, deployGasPriceOracle := toDepositTxn(t, upgradeTxns[1])
	require.Equal(t, deployGasPriceOracleSender, common.HexToAddress("0x4210000000000000000000000000000000000001"))
	require.Equal(t, deployGasPriceOracleSource.SourceHash(), deployGasPriceOracle.SourceHash())
	require.Nil(t, deployGasPriceOracle.To())
	require.Equal(t, uint64(1_000_000), deployGasPriceOracle.Gas())
	require.Equal(t, bindings.GasPriceOracleMetaData.Bin, hexutil.Bytes(deployGasPriceOracle.Data()).String())

	updateL1BlockProxySender, updateL1BlockProxy := toDepositTxn(t, upgradeTxns[2])
	require.Equal(t, updateL1BlockProxySender, common.Address{})
	require.Equal(t, updateL1BlockProxySource.SourceHash(), updateL1BlockProxy.SourceHash())
	require.NotNil(t, updateL1BlockProxy.To())
	require.Equal(t, *updateL1BlockProxy.To(), common.HexToAddress("0x4200000000000000000000000000000000000015"))
	require.Equal(t, uint64(50_000), updateL1BlockProxy.Gas())
	require.Equal(t, common.FromHex("0x3659cfe600000000000000000000000007dbe8500fc591d1852b76fee44d5a05e13097ff"), updateL1BlockProxy.Data())

	updateGasPriceOracleSender, updateGasPriceOracle := toDepositTxn(t, upgradeTxns[3])
	require.Equal(t, updateGasPriceOracleSender, common.Address{})
	require.Equal(t, updateGasPriceOracleSource.SourceHash(), updateGasPriceOracle.SourceHash())
	require.NotNil(t, updateGasPriceOracle.To())
	require.Equal(t, *updateGasPriceOracle.To(), common.HexToAddress("0x420000000000000000000000000000000000000F"))
	require.Equal(t, uint64(50_000), updateGasPriceOracle.Gas())
	require.Equal(t, common.FromHex("0x3659cfe6000000000000000000000000b528d11cc114e026f138fe568744c6d45ce6da7a"), updateGasPriceOracle.Data())

	gpoSetEcotoneSender, gpoSetEcotone := toDepositTxn(t, upgradeTxns[4])
	require.Equal(t, gpoSetEcotoneSender, common.HexToAddress("0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001"))
	require.Equal(t, enableEcotoneSource.SourceHash(), gpoSetEcotone.SourceHash())
	require.NotNil(t, gpoSetEcotone.To())
	require.Equal(t, *gpoSetEcotone.To(), common.HexToAddress("0x420000000000000000000000000000000000000F"))
	require.Equal(t, uint64(80_000), gpoSetEcotone.Gas())
	require.Equal(t, common.FromHex("0x22b90ab3"), gpoSetEcotone.Data())

	beaconRootsSender, beaconRoots := toDepositTxn(t, upgradeTxns[5])
	require.Equal(t, beaconRootsSender, common.HexToAddress("0x0B799C86a49DEeb90402691F1041aa3AF2d3C875"))
	require.Equal(t, beaconRootsSource.SourceHash(), beaconRoots.SourceHash())
	require.Nil(t, beaconRoots.To())
	require.Equal(t, uint64(250_000), beaconRoots.Gas())
	require.Equal(t, eip4788CreationData, beaconRoots.Data())
	require.NotEmpty(t, beaconRoots.Data())
}

func TestEip4788Params(t *testing.T) {
	require.Equal(t, EIP4788From, common.HexToAddress("0x0B799C86a49DEeb90402691F1041aa3AF2d3C875"))
	require.Equal(t, eip4788CreationData, common.FromHex("0x60618060095f395ff33373fffffffffffffffffffffffffffffffffffffffe14604d57602036146024575f5ffd5b5f35801560495762001fff810690815414603c575f5ffd5b62001fff01545f5260205ff35b5f5ffd5b62001fff42064281555f359062001fff015500"))
	require.NotEmpty(t, eip4788CreationData)
}

*/
