package types

import (
	//distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	//"github.com/cosmos/cosmos-sdk/x/distribution/types"
	//"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	//
	//"github.com/soohoio/stayking/utils"
)

const TypeMsgSetWithdrawalAddress = "set_withdrawal_address"

var _ sdk.Msg = &MsgSetWithdrawalAddress{}

func NewMsgSetWithdrawalAddress(delegatorAddress string, withdrawAddress string) *MsgSetWithdrawalAddress {
	return &MsgSetWithdrawalAddress {
		DelegatorAddress: delegatorAddress,
		WithdrawAddress:  withdrawAddress,
	}
}

func (msg *MsgSetWithdrawalAddress) Route() string {
	return RouterKey
}

func (msg *MsgSetWithdrawalAddress) Type() string {
	return TypeMsgSetWithdrawalAddress
}

func (msg *MsgSetWithdrawalAddress) GetSigners() []sdk.AccAddress {
	//creator, err := sdk.AccAddressFromBech32(msg.Creator)
	//if err != nil {
	//	panic(err)
	//}
	//return []sdk.AccAddress{creator}
	return nil;
}

func (msg *MsgSetWithdrawalAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetWithdrawalAddress) ValidateBasic() error {


	return nil
}