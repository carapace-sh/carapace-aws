package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_createOrderCmd = &cobra.Command{
	Use:   "create-order",
	Short: "Creates an order for an Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_createOrderCmd).Standalone()

	outposts_createOrderCmd.Flags().String("line-items", "", "The line items that make up the order.")
	outposts_createOrderCmd.Flags().String("outpost-identifier", "", "The ID or the Amazon Resource Name (ARN) of the Outpost.")
	outposts_createOrderCmd.Flags().String("payment-option", "", "The payment option.")
	outposts_createOrderCmd.Flags().String("payment-term", "", "The payment terms.")
	outposts_createOrderCmd.MarkFlagRequired("outpost-identifier")
	outposts_createOrderCmd.MarkFlagRequired("payment-option")
	outpostsCmd.AddCommand(outposts_createOrderCmd)
}
