package tron

import "github.com/fbsobreira/gotron-sdk/pkg/proto/core"

// https://developers.tron.network/reference/delegateresource-1
func (svc *TronService) DelegateEnergy(
	sender *TxSender,
	userAddr string,
	delegateBalance int64,
	lock bool,
	lockPeriod int64) (*TransferResult, error) {
	tx, err := svc.client.DelegateResource(sender.Address, userAddr, core.ResourceCode_ENERGY, delegateBalance, lock, lockPeriod)
	if err != nil {
		return nil, err
	}

	controller := NewTronController(svc.client, sender, tx.Transaction)
	if err = controller.ExecuteTransaction(); err != nil {
		return nil, err
	}

	rs := assembleTransferResult(controller, tx, "")

	return rs, nil

}
