package ton

import (
	"github.com/davecgh/go-spew/spew"
	"math/big"
	"testing"
)

var testAlex = &TestAccount{
	Address:    "UQAEirkgAaB3ZBMBbpLCseMrx5CojWp_u-T4VrNlSAeUQrrv",
	PrivateKey: "0x856ed5fd7b2e89b75a360865c9fccc347ac38a83c2c0040565d66fee1085cb6e032fe72aa0f9ece8b2e80203285121954740dc03c7e106b167b4bb6257e64ca8",
	Mnemonic:   "icon group stamp shuffle quantum demise dose thing opera trap front make pattern satisfy ahead monkey fee census fetch concert wood top skull puppy",
}

var testMao = &TestAccount{
	Address:    "UQCFMdsNB2JLiTZszuaoz4DminjT4Cc9UtdaVHSWRc-Q17-q",
	PrivateKey: "",
	Mnemonic:   "",
}

const FishJettonContract = "EQATcUc69sGSCCMSadsVUKdGwM1BMKS-HKCWGPk60xZGgwsK"

func TestTonService_Jetten(t *testing.T) {
	contractAddr := "EQATcUc69sGSCCMSadsVUKdGwM1BMKS-HKCWGPk60xZGgwsK"

	svc := NewTonService(false)
	data, err := svc.GetJettonData(contractAddr)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(data)
}

func TestTonService_GetJettonBalance(t *testing.T) {
	svc := NewTonService(false)
	data, err := svc.GetJettonBalance(&TxSender{
		Address: testAlex.Address, PrivateKey: testAlex.PrivateKey,
	}, FishJettonContract)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(data)
}

func TestTonService_TrasferJetton(t *testing.T) {
	svc := NewTonService(false)
	rs, err := svc.TransferJetton(&TxSender{
		Address: testAlex.Address, PrivateKey: testAlex.PrivateKey,
	}, FishJettonContract, testMao.Address, big.NewInt(100000000000))
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(rs)
}
