package mycustommodulecli

import (
	"Autheo/x/mycustommoduletypes" // Replace with your module's import path

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/spf13/cobra"
)

func NewSubmitProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-my-custom-proposal [title] [description]",
		Short: "Submit a custom governance proposal",
		Long:  `Submit a custom governance proposal with a title, description, and optionally additional parameters.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title := args[0]
			description := args[1]

			// Optional: Parse additional flags/parameters if your proposal needs more fields.

			// Create the custom proposal
			content := mycustommoduletypes.MyCustomProposal{
				Title:       title,
				Description: description,
				// Set additional fields here...
			}

			// Validate the content
			if err := content.ValidateBasic(); err != nil {
				return err
			}

			// Read deposit flag
			depositStr, err := cmd.Flags().GetString("Deposit")
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			proposer := clientCtx.GetFromAddress()

			// Create the message
			msg, err := govtypes.NewMsgSubmitProposal(
				content,  // Proposal content (govtypes.Content)
				deposit,  // Initial deposit
				proposer, // Proposer’s address (AccAddress)
			)
			if err != nil {
				return err
			}

			// Generate or broadcast the transaction
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	// Add flags for deposit and transaction configuration
	cmd.Flags().String("Deposit", "", "Deposit for the proposal")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
