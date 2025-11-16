package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_cancelOrderCmd = &cobra.Command{
	Use:   "cancel-order",
	Short: "Cancels the specified order for an Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_cancelOrderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_cancelOrderCmd).Standalone()

		outposts_cancelOrderCmd.Flags().String("order-id", "", "The ID of the order.")
		outposts_cancelOrderCmd.MarkFlagRequired("order-id")
	})
	outpostsCmd.AddCommand(outposts_cancelOrderCmd)
}
