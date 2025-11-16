package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_getOrderCmd = &cobra.Command{
	Use:   "get-order",
	Short: "Gets information about the specified order.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_getOrderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_getOrderCmd).Standalone()

		outposts_getOrderCmd.Flags().String("order-id", "", "The ID of the order.")
		outposts_getOrderCmd.MarkFlagRequired("order-id")
	})
	outpostsCmd.AddCommand(outposts_getOrderCmd)
}
