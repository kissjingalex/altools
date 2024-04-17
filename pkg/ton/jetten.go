package ton

import (
	"encoding/base64"
	"fmt"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton/jetton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"math/big"
)

type JettonData struct {
	TotalSupply *big.Int
	Mintable    bool
	AdminAddr   *address.Address
}

type JettonBalance struct {
	TokenWallet string
	Balance     *big.Int
}

func (svc *TonService) GetJettonBalance(sender *TxSender, contractAddr string) (*JettonBalance, error) {
	client := jetton.NewJettonMasterClient(svc.apiClient, address.MustParseAddr(contractAddr))

	w, err := svc.getRawWalletFromPrivateKey(sender.PrivateKey)
	if err != nil {
		fmt.Printf("sender is not valid. %v\n", err)
		return nil, err
	}

	tokenWallet, err := client.GetJettonWallet(svc.ctx, w.WalletAddress())
	if err != nil {
		fmt.Printf("fail to get sender token wallet. %s\n", sender.Address)
		return nil, err
	}

	b, err := tokenWallet.GetBalance(svc.ctx)
	if err != nil {
		fmt.Printf("fail to get balance for sender. %s\n", sender.Address)
		return nil, err
	}

	return &JettonBalance{
		TokenWallet: tokenWallet.Address().String(),
		Balance:     b,
	}, nil
}

func (svc *TonService) GetJettonData(contractAddr string) (*JettonData, error) {
	b, err := svc.apiClient.CurrentMasterchainInfo(svc.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get masterchain info: %w", err)
	}

	contractAddress := address.MustParseAddr(contractAddr)
	res, err := svc.apiClient.WaitForBlock(b.SeqNo).RunGetMethod(svc.ctx, b, contractAddress, "get_jetton_data")
	if err != nil {
		return nil, fmt.Errorf("failed to run get_jetton_data method: %w", err)
	}

	supply, err := res.Int(0)
	if err != nil {
		return nil, fmt.Errorf("supply get err: %w", err)
	}

	mintable, err := res.Int(1)
	if err != nil {
		return nil, fmt.Errorf("mintable get err: %w", err)
	}

	adminAddr, err := res.Slice(2)
	if err != nil {
		return nil, fmt.Errorf("admin addr get err: %w", err)
	}
	addr, err := adminAddr.LoadAddr()
	if err != nil {
		return nil, fmt.Errorf("failed to load address from adminAddr slice: %w", err)
	}

	return &JettonData{
		TotalSupply: supply,
		Mintable:    mintable.Cmp(big.NewInt(0)) != 0,
		AdminAddr:   addr,
	}, nil
}

func (svc *TonService) TransferJetton(sender *TxSender, contractAddr string, userAddr string, amount *big.Int) (*TransferResult, error) {
	client := jetton.NewJettonMasterClient(svc.apiClient, address.MustParseAddr(contractAddr))

	w, err := svc.getRawWalletFromPrivateKey(sender.PrivateKey)
	if err != nil {
		fmt.Printf("sender is not valid. %v\n", err)
		return nil, err
	}

	tokenWallet, err := client.GetJettonWallet(svc.ctx, w.WalletAddress())
	if err != nil {
		fmt.Printf("fail to get sender token wallet. %s\n", sender.Address)
		return nil, err
	}

	b, err := tokenWallet.GetBalance(svc.ctx)
	if err != nil {
		fmt.Printf("fail to get balance for sender. %s\n", sender.Address)
		return nil, err
	}

	fmt.Printf("balance=%s\n", b.String())

	fmt.Println("Transferring tokens...")

	amountTokens := tlb.MustFromNano(amount, 9)

	// address of receiver's wallet (not token wallet, just usual)
	to := address.MustParseAddr(userAddr)
	transferPayload, err := tokenWallet.BuildTransferPayloadV2(to, to, amountTokens, tlb.ZeroCoins, nil, nil)
	if err != nil {
		return nil, err
	}

	//TODO 暂时不知道为什么一定要0.05 （jetton wallet的转账费用）
	amountTon := tlb.MustFromTON("0.05") // tlb.ZeroCoins
	msg := wallet.SimpleMessage(tokenWallet.Address(), amountTon, transferPayload)

	fmt.Println("sending transaction...")
	tx, block, err := w.SendWaitTransaction(svc.ctx, msg)
	if err != nil {
		panic(err)
	}

	return &TransferResult{
		TxHash:     base64.StdEncoding.EncodeToString(tx.Hash),
		BlockSeqNo: block.SeqNo,
		Balance:    nil,
	}, nil
}
