package tx

import (
	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/cosmos/cosmos-proto/anyutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetTxBody(tx sdk.Tx) (*txv1beta1.TxBody, error) {
	msgsV2, err := tx.GetMsgsV2()
	if err != nil {
		return nil, err
	}

	anyMsgs := make([]*anypb.Any, len(msgsV2))
	for i, msg := range msgsV2 {
		anyMsgs[i], err = anyutil.New(msg)
		if err != nil {
			return nil, err
		}
	}

	return &txv1beta1.TxBody{
		Messages:      anyMsgs,
		Memo:          getTxMemo(tx),
		TimeoutHeight: getTxWithTimeoutHeight(tx),
	}, nil
}

func getTxMemo(tx sdk.Tx) string {
	if tx, ok := tx.(sdk.TxWithMemo); ok {
		return tx.GetMemo()
	}

	return ""
}

func getTxWithTimeoutHeight(tx sdk.Tx) uint64 {
	if tx, ok := tx.(sdk.TxWithTimeoutHeight); ok {
		return tx.GetTimeoutHeight()
	}

	return 0
}
