// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1BlockStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"number\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1001,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"timestamp\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1002,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"basefee\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"hash\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_bytes32\"},{\"astId\":1004,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"sequenceNumber\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint64\"},{\"astId\":1005,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"blobBaseFeeScalar\",\"offset\":8,\"slot\":\"3\",\"type\":\"t_uint32\"},{\"astId\":1006,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"baseFeeScalar\",\"offset\":12,\"slot\":\"3\",\"type\":\"t_uint32\"},{\"astId\":1007,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"l1FeeOverhead\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"l1FeeScalar\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_uint256\"},{\"astId\":1010,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"blobBaseFee\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_uint256\"}],\"types\":{\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"}}}"

var L1BlockStorageLayout = new(solc.StorageLayout)

var L1BlockDeployedBin = "0x608060405234801561001057600080fd5b50600436106100f55760003560e01c80638381f58a11610097578063c598591811610066578063c598591814610229578063e591b28214610249578063e81b2c6d14610289578063f82061401461029257600080fd5b80638381f58a146101e35780638b239f73146101f75780639e8c496614610200578063b80777ea1461020957600080fd5b806354fd4d50116100d357806354fd4d50146101335780635cf249691461017c57806364ca23ef1461018557806368d5dca6146101b257600080fd5b8063015d8eb9146100fa57806309bd5a601461010f578063440a5e201461012b575b600080fd5b61010d61010836600461044c565b61029b565b005b61011860025481565b6040519081526020015b60405180910390f35b61010d6103da565b61016f6040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b60405161012291906104be565b61011860015481565b6003546101999067ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610122565b6003546101ce9068010000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610122565b6000546101999067ffffffffffffffff1681565b61011860055481565b61011860065481565b6000546101999068010000000000000000900467ffffffffffffffff1681565b6003546101ce906c01000000000000000000000000900463ffffffff1681565b61026473deaddeaddeaddeaddeaddeaddeaddeaddead000181565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610122565b61011860045481565b61011860075481565b3373deaddeaddeaddeaddeaddeaddeaddeaddead000114610342576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603b60248201527f4c31426c6f636b3a206f6e6c7920746865206465706f7369746f72206163636f60448201527f756e742063616e20736574204c3120626c6f636b2076616c7565730000000000606482015260840160405180910390fd5b6000805467ffffffffffffffff98891668010000000000000000027fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116998916999099179890981790975560019490945560029290925560038054919094167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009190911617909255600491909155600555600655565b3373deaddeaddeaddeaddeaddeaddeaddeaddead00011461040357633cc50b456000526004601cfd5b60043560801c60035560143560801c600055602435600155604435600755606435600255608435600455565b803567ffffffffffffffff8116811461044757600080fd5b919050565b600080600080600080600080610100898b03121561046957600080fd5b6104728961042f565b975061048060208a0161042f565b9650604089013595506060890135945061049c60808a0161042f565b979a969950949793969560a0850135955060c08501359460e001359350915050565b600060208083528351808285015260005b818110156104eb578581018301518582016040015282016104cf565b818111156104fd576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01692909201604001939250505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L1BlockStorageLayoutJSON), L1BlockStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1Block"] = L1BlockStorageLayout
	deployedBytecodes["L1Block"] = L1BlockDeployedBin
	immutableReferences["L1Block"] = false
}
