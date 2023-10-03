package tx

import (
	txv1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TxEncoder returns a default protobuf v2 TxEncoder
func TxEncoder() sdk.TxEncoder {
	return func(tx sdk.Tx) ([]byte, error) {
		body, err := GetTxBody(tx)
		if err != nil {
			return nil, err
		}

		raw := &txv1beta1.Tx{
			Body: body,
			// AuthInfo: txWrapper.getAuthInfoBytes(),
			// Signatures: txWrapper.tx.Signatures,
		}

		return proto.Marshal(raw)
	}
}

// TxJSONEncoder returns a default protobuf v2 JSON TxEncoder.
func TxJSONEncoder(cdc codec.Codec) sdk.TxEncoder {
	return func(tx sdk.Tx) ([]byte, error) {
		body, err := GetTxBody(tx)
		if err != nil {
			return nil, err
		}

		return cdc.MarshalJSON(&txv1beta1.Tx{
			Body: body,
			// AuthInfo: txWrapper.getAuthInfoBytes(),
			// Signatures: txWrapper.tx.Signatures,
		})
	}
}
