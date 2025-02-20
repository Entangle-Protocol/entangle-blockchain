package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/Entangle-Protocol/entangle-blockchain/x/distributorsauth/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdAddDistributor(),
		CmdRemoveDistributor(),
		CmdAddAdmin(),
		CmdRemoveAdmin(),
		CmdSubmitDistributorsAuthProposal(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdAddDistributor() *cobra.Command {
	return &cobra.Command{
		Use:   "add-distributor [address] [endTime]",
		Short: "Add Distributor",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			//argsAddress := string(args[0])
			// argsAddress, err := sdk.AccAddressFromBech32(string(args[0]))
			// if err != nil {
			// 	return err
			// }
			argsAddress := string(args[0])
			argsAddress, err := types.EnsureBech32Address(argsAddress)
			if err != nil {
				return err
			}

			argsEndTime := string(args[1])
			argsEndTimeUint64, err := strconv.ParseUint(string(argsEndTime), 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddDistributor(clientCtx.GetFromAddress(), argsAddress, argsEndTimeUint64)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)

			// new := argsEndTimeUint64
			// new = new + 1
			// clientCtx = nil
			// return nil
		},
	}
}

func CmdRemoveDistributor() *cobra.Command {
	return &cobra.Command{
		Use:   "remove-distributor [address]",
		Short: "Remove Distributor",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAddress := string(args[0])
			argsAddress, err := types.EnsureBech32Address(argsAddress)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveDistributor(clientCtx.GetFromAddress(), argsAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

func CmdAddAdmin() *cobra.Command {
	return &cobra.Command{
		Use:   "add-admin [address]",
		Short: "Add Admin",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAddress := string(args[0])
			argsAddress, err := types.EnsureBech32Address(argsAddress)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddAdmin(clientCtx.GetFromAddress(), argsAddress, false)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

func CmdRemoveAdmin() *cobra.Command {
	return &cobra.Command{
		Use:   "remove-admin [address]",
		Short: "Remove Admin",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAddress := string(args[0])
			argsAddress, err := types.EnsureBech32Address(argsAddress)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveAdmin(clientCtx.GetFromAddress(), argsAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetCmdSubmitDistributorsAuthProposal implements the command to submit a add_distributor proposal
func CmdSubmitDistributorsAuthProposal() *cobra.Command {
	bech32PrefixAccAddr := sdk.GetConfig().GetBech32AccountAddrPrefix()

	cmd := &cobra.Command{
		Use:   "add-distributor-gov [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit add distributor proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit dd distributor proposal .
The proposal details must be supplied via a JSON file.

Example:
$ %s tx gov submit-legacy-proposal add-distributor-gov  <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:

{
  "title": "Add very new distributor",
  "description": "Very honorable protocol integration",
  "address": "%s1s5afhd6gxevu37mkqcvvsj8qeylhn0rz46zdlq",
  "end_date": "1701363600",
  "deposit": "10000000aNGL"
}
`,
				version.AppName, bech32PrefixAccAddr,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := ParseAddDistributorProposal(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}

			// amount, err := sdk.ParseCoinsNormalized(proposal.Title)
			// if err != nil {
			// 	return err
			// }

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			content := types.NewAddDistributorProposal(proposal.Title, proposal.Description, proposal.Address, proposal.Deposit, proposal.EndDate)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
