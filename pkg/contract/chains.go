package contract

import "math/big"

var LocalChainConfig = &ChainConfig{
	GateWay:      "http://127.0.0.1:8545",
	WsGateWay:    "",
	ContractAddr: "",
	UserAddr:     "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
	PrivateKey:   "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	ChainId:      big.NewInt(31337),
}

var GoerliChainConfig = &ChainConfig{
	GateWay:      "https://rpc.ankr.com/eth_goerli/b0c6c54f3dc9cc51677e475fe355253c097e527a62d0cb188f1c882561ae1417",
	WsGateWay:    "",
	ContractAddr: "",
	UserAddr:     "0x9c52262e33C66a347C063D02B1D5A5F7897523f7",
	PrivateKey:   "c08a70581ef8f3b3e12beec950facea6c19bee0df46a27eac0777520d76e1063",
	ChainId:      big.NewInt(5),
}

var BaseGoerliChainConfig = &ChainConfig{
	GateWay:      "https://rpc.ankr.com/base_goerli/b0c6c54f3dc9cc51677e475fe355253c097e527a62d0cb188f1c882561ae1417",
	WsGateWay:    "",
	ContractAddr: "",
	UserAddr:     "0x9c52262e33C66a347C063D02B1D5A5F7897523f7",
	PrivateKey:   "c08a70581ef8f3b3e12beec950facea6c19bee0df46a27eac0777520d76e1063",
	ChainId:      big.NewInt(84531),
}

type ChainUser struct {
	Address    string
	PrivateKey string
}

type ChainUserGroup struct {
	Admin  ChainUser
	Seller ChainUser
	Buyer  ChainUser
}

var LocalUserGroup = &ChainUserGroup{
	Admin: ChainUser{
		Address:    "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		PrivateKey: "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	},
	Seller: ChainUser{
		Address:    "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
		PrivateKey: "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
	},
	Buyer: ChainUser{
		Address:    "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
		PrivateKey: "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
	},
}

var BaseGoerliUserGroup = &ChainUserGroup{
	Admin: ChainUser{
		Address:    "0x9c52262e33C66a347C063D02B1D5A5F7897523f7",
		PrivateKey: "c08a70581ef8f3b3e12beec950facea6c19bee0df46a27eac0777520d76e1063",
	},
	Seller: ChainUser{
		Address:    "0xbeCCAdb67acBe8377AE1C526F237e63Cd37305e0",
		PrivateKey: "56c9b431dbee05d98aeed5e165976df7a67289ae71b88dbea3d33a3922b282c5",
	},
	Buyer: ChainUser{
		Address:    "0xDb4e8bf873A83A593745bbF9e1aB12Cd6BcC3D8A",
		PrivateKey: "3dcaa21dd925c18bea8f91de7e6ecbd71a12c3ef80ed96822aeb910030401271",
	},
}
