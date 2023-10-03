package tx

import (
	"fmt"

	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/cosmos/cosmos-proto/anyutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
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

func GetTxSigs(tx sdk.Tx) (*txv1beta1.AuthInfo, [][]byte, error) {
	verifiableTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return nil, nil, fmt.Errorf("invalid tx type; expected: %T, got: %T", (*authsigning.SigVerifiableTx)(nil), tx)
	}

	sigs, err := verifiableTx.GetSignaturesV2()
	if err != nil {
		return nil, nil, err
	}

	signers, err := verifiableTx.GetSigners()
	if err != nil {
		return nil, nil, err
	}

	// check that signer length and signature length are the same
	if len(sigs) != len(signers) {
		return nil, nil, fmt.Errorf("invalid number of signer;  expected: %d, got %d", len(signers), len(sigs))
	}

	// TODO

	return nil, nil, nil
}
