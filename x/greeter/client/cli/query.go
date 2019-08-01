package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"

	types "github.com/cosmos/hellochain/x/greeter/types"
)

// TODO comment
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {

	greeterQueryCmd := &cobra.Command{
		Use:                        "greeter",
		Short:                      "Querying commands for the greeter module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	greeterQueryCmd.AddCommand(client.GetCommands(
		GetCmdListGreetings(storeKey, cdc),
	)...)
	return greeterQueryCmd
}

// GetCmdResolveGreetings queries all greetings
func GetCmdListGreetings(queryRoute string, cdc *codec.Codec) *cobra.Command {

	return &cobra.Command{
		Use:   "list [addr]",
		Short: "list greetings for address. Usage list [address]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr := args[0]

			//res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", queryRoute, name), nil)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list/%s", queryRoute, addr), nil)
			if err != nil {

				return nil
			}

			out := types.NewQueryResGreetings()
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
