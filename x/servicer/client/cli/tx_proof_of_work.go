package cli

import (
	"strconv"

	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/lavanet/lava/x/servicer/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdProofOfWork() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proof-of-work [spec] [session] [client-request] [work-proof] [compute-units] [block-of-work]",
		Short: "Broadcast message proofOfWork",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSpec := &types.SpecName{Name: args[0]}
			num, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argSession := &types.SessionID{Num: num}
			argClientRequest := new(types.ClientRequest)
			err = json.Unmarshal([]byte(args[2]), argClientRequest)
			if err != nil {
				return err
			}
			argWorkProof := new(types.WorkProof)
			err = json.Unmarshal([]byte(args[3]), argWorkProof)
			if err != nil {
				return err
			}
			argComputeUnits, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}
			num, err = cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argBlockOfWork := &types.BlockNum{Num: num}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProofOfWork(
				clientCtx.GetFromAddress().String(),
				argSpec,
				argSession,
				argClientRequest,
				argWorkProof,
				argComputeUnits,
				argBlockOfWork,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}