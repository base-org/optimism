package derive

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
)

var (
	// Gas Price Oracle Parameters
	deployFjordGasPriceOracleSource       = UpgradeDepositSource{Intent: "Fjord: Gas Price Oracle Deployment"}
	GasPriceOracleFjordDeployerAddress    = common.HexToAddress("0x4210000000000000000000000000000000000002")
	gasPriceOracleFjordDeploymentBytecode = common.FromHex("0x608060405234801561001057600080fd5b5061187d806100206000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c80636ef25c3a116100b2578063de26c4a111610081578063f45e65d811610066578063f45e65d81461025b578063f820614014610263578063fe173b971461020d57600080fd5b8063de26c4a114610235578063f1c7a58b1461024857600080fd5b80636ef25c3a1461020d5780638e98b10614610213578063960e3a231461021b578063c59859181461022d57600080fd5b806349948e0e11610109578063519b4bd3116100ee578063519b4bd31461019f57806354fd4d50146101a757806368d5dca6146101f057600080fd5b806349948e0e1461016f5780634ef6e2241461018257600080fd5b80630c18c1621461013b57806322b90ab3146101565780632e0f262514610160578063313ce56714610168575b600080fd5b61014361026b565b6040519081526020015b60405180910390f35b61015e61038c565b005b610143600681565b6006610143565b61014361017d3660046113ae565b6105af565b60005461018f9060ff1681565b604051901515815260200161014d565b610143610600565b6101e36040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b60405161014d919061147d565b6101f8610661565b60405163ffffffff909116815260200161014d565b48610143565b61015e6106e6565b60005461018f90610100900460ff1681565b6101f861097a565b6101436102433660046113ae565b6109db565b6101436102563660046114f0565b610ac9565b610143610b99565b610143610c8c565b6000805460ff1615610304576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f47617350726963654f7261636c653a206f76657268656164282920697320646560448201527f707265636174656400000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103879190611509565b905090565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663e591b2826040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061040f9190611522565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f47617350726963654f7261636c653a206f6e6c7920746865206465706f73697460448201527f6f72206163636f756e742063616e2073657420697345636f746f6e6520666c6160648201527f6700000000000000000000000000000000000000000000000000000000000000608482015260a4016102fb565b60005460ff1615610582576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f47617350726963654f7261636c653a2045636f746f6e6520616c72656164792060448201527f616374697665000000000000000000000000000000000000000000000000000060648201526084016102fb565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60008054610100900460ff16156105e3576105dd6105cc83610ced565b516105d8906044611587565b611012565b92915050565b60005460ff16156105f7576105dd826110e0565b6105dd82611184565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff166368d5dca66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106c2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610387919061159f565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663e591b2826040518163ffffffff1660e01b8152600401602060405180830381865afa158015610745573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107699190611522565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610823576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f47617350726963654f7261636c653a206f6e6c7920746865206465706f73697460448201527f6f72206163636f756e742063616e20736574206973466a6f726420666c61670060648201526084016102fb565b60005460ff166108b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f47617350726963654f7261636c653a20466a6f72642063616e206f6e6c79206260448201527f65206163746976617465642061667465722045636f746f6e650000000000000060648201526084016102fb565b600054610100900460ff161561094c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f47617350726963654f7261636c653a20466a6f726420616c726561647920616360448201527f746976650000000000000000000000000000000000000000000000000000000060648201526084016102fb565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663c59859186040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106c2573d6000803e3d6000fd5b60008054610100900460ff1615610a1657610a0b6109f883610ced565b51610a04906044611587565b60476112d8565b6105dd9060106115c5565b6000610a21836112ef565b60005490915060ff1615610a355792915050565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab89190611509565b610ac29082611587565b9392505050565b60008054610100900460ff16610b61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f47617350726963654f7261636c653a206765744c314665655570706572426f7560448201527f6e64206f6e6c7920737570706f72747320466a6f72640000000000000000000060648201526084016102fb565b6000610b6e60ff84611602565b610b789084611587565b610b83906010611587565b610b8e906044611587565b9050610ac281611012565b6000805460ff1615610c2d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f47617350726963654f7261636c653a207363616c61722829206973206465707260448201527f656361746564000000000000000000000000000000000000000000000000000060648201526084016102fb565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663f82061406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b6060610e84565b818153600101919050565b600082840393505b83811015610ac25782810151828201511860001a1590930292600101610d07565b825b60208210610d74578251610d3f601f83610cf4565b52602092909201917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090910190602101610d2a565b8115610ac2578251610d896001840383610cf4565b520160010192915050565b60006001830392505b6101078210610dd557610dc78360ff16610dc260fd610dc28760081c60e00189610cf4565b610cf4565b935061010682039150610d9d565b60078210610e0257610dfb8360ff16610dc260078503610dc28760081c60e00189610cf4565b9050610ac2565b610e1b8360ff16610dc28560081c8560051b0187610cf4565b949350505050565b610e7c828203610e60610e5084600081518060001a8160011a60081b178160021a60101b17915050919050565b639e3779b90260131c611fff1690565b8060021b6040510182815160e01c1860e01b8151188152505050565b600101919050565b6180003860405139618000604051016020830180600d8551820103826002015b81811015610fb7576000805b50508051604051600082901a600183901a60081b1760029290921a60101b91909117639e3779b9810260111c617ffc16909101805160e081811c878603811890911b90911890915284019081830390848410610f0c5750610f47565b600184019350611fff8211610f41578251600081901a600182901a60081b1760029190911a60101b178103610f415750610f47565b50610eb0565b838310610f55575050610fb7565b60018303925085831115610f7357610f708787888603610d28565b96505b610f87600985016003850160038501610cff565b9150610f94878284610d94565b965050610fac84610fa786848601610e23565b610e23565b915050809350610ea4565b5050610fc98383848851850103610d28565b925050506040519150618000820180820391508183526020830160005b83811015610ffe578281015182820152602001610fe6565b506000920191825250602001604052919050565b600061101f8260476112d8565b9150600061102b610c8c565b611033610661565b63ffffffff1661104391906115c5565b61104b610600565b61105361097a565b61105e90601061163d565b63ffffffff1661106e91906115c5565b6110789190611587565b9050600061108984620cc3946115c5565b6110b3907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd763200611669565b90506110c1600660026115c5565b6110cc90600a6117fd565b6110d683836115c5565b610e1b9190611602565b6000806110ec836112ef565b905060006110f8610600565b61110061097a565b61110b90601061163d565b63ffffffff1661111b91906115c5565b90506000611127610c8c565b61112f610661565b63ffffffff1661113f91906115c5565b9050600061114d8284611587565b61115790856115c5565b90506111656006600a6117fd565b6111709060106115c5565b61117a9082611602565b9695505050505050565b600080611190836112ef565b9050600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112179190611509565b61121f610600565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa15801561127e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a29190611509565b6112ac9085611587565b6112b691906115c5565b6112c091906115c5565b90506112ce6006600a6117fd565b610e1b9082611602565b6000818310156112e85781610ac2565b5090919050565b80516000908190815b818110156113725784818151811061131257611312611809565b01602001517fff00000000000000000000000000000000000000000000000000000000000000166000036113525761134b600484611587565b9250611360565b61135d601084611587565b92505b8061136a81611838565b9150506112f8565b50610e1b82610440611587565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156113c057600080fd5b813567ffffffffffffffff808211156113d857600080fd5b818401915084601f8301126113ec57600080fd5b8135818111156113fe576113fe61137f565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156114445761144461137f565b8160405282815287602084870101111561145d57600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156114aa5785810183015185820160400152820161148e565b818111156114bc576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60006020828403121561150257600080fd5b5035919050565b60006020828403121561151b57600080fd5b5051919050565b60006020828403121561153457600080fd5b815173ffffffffffffffffffffffffffffffffffffffff81168114610ac257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561159a5761159a611558565b500190565b6000602082840312156115b157600080fd5b815163ffffffff81168114610ac257600080fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156115fd576115fd611558565b500290565b600082611638577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600063ffffffff8083168185168183048111821515161561166057611660611558565b02949350505050565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038413811516156116a3576116a3611558565b827f80000000000000000000000000000000000000000000000000000000000000000384128116156116d7576116d7611558565b50500190565b600181815b8085111561173657817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561171c5761171c611558565b8085161561172957918102915b93841c93908002906116e2565b509250929050565b60008261174d575060016105dd565b8161175a575060006105dd565b8160018114611770576002811461177a57611796565b60019150506105dd565b60ff84111561178b5761178b611558565b50506001821b6105dd565b5060208310610133831016604e8410600b84101617156117b9575081810a6105dd565b6117c383836116dd565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156117f5576117f5611558565b029392505050565b6000610ac2838361173e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361186957611869611558565b506001019056fea164736f6c634300080f000a")

	// Update GasPricePriceOracle Proxy Parameters
	updateFjordGasPriceOracleSource = UpgradeDepositSource{Intent: "Fjord: Gas Price Oracle Proxy Update"}
	fjordGasPriceOracleAddress      = common.HexToAddress("0xa919894851548179A0750865e7974DA599C0Fac7")

	// Enable Fjord Parameters
	enableFjordSource = UpgradeDepositSource{Intent: "Fjord: Gas Price Oracle Set Fjord"}
	enableFjordInput  = crypto.Keccak256([]byte("setFjord()"))[:4]
)

// FjordNetworkUpgradeTransactions returns the transactions required to upgrade the Fjord network.
func FjordNetworkUpgradeTransactions() ([]hexutil.Bytes, error) {
	upgradeTxns := make([]hexutil.Bytes, 0, 3)

	// Deploy Gas Price Oracle transaction
	deployGasPriceOracle, err := types.NewTx(&types.DepositTx{
		SourceHash:          deployFjordGasPriceOracleSource.SourceHash(),
		From:                GasPriceOracleFjordDeployerAddress,
		To:                  nil,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 1_450_000,
		IsSystemTransaction: false,
		Data:                gasPriceOracleFjordDeploymentBytecode,
	}).MarshalBinary()

	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, deployGasPriceOracle)

	updateGasPriceOracleProxy, err := types.NewTx(&types.DepositTx{
		SourceHash:          updateFjordGasPriceOracleSource.SourceHash(),
		From:                common.Address{},
		To:                  &predeploys.GasPriceOracleAddr,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 50_000,
		IsSystemTransaction: false,
		Data:                upgradeToCalldata(fjordGasPriceOracleAddress),
	}).MarshalBinary()

	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, updateGasPriceOracleProxy)

	enableFjord, err := types.NewTx(&types.DepositTx{
		SourceHash:          enableFjordSource.SourceHash(),
		From:                L1InfoDepositerAddress,
		To:                  &predeploys.GasPriceOracleAddr,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 90_000,
		IsSystemTransaction: false,
		Data:                enableFjordInput,
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}
	upgradeTxns = append(upgradeTxns, enableFjord)

	return upgradeTxns, nil
}
