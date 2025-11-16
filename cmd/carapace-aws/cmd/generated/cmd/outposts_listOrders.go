package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_listOrdersCmd = &cobra.Command{
	Use:   "list-orders",
	Short: "Lists the Outpost orders for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_listOrdersCmd).Standalone()

	outposts_listOrdersCmd.Flags().String("max-results", "", "")
	outposts_listOrdersCmd.Flags().String("next-token", "", "")
	outposts_listOrdersCmd.Flags().String("outpost-identifier-filter", "", "The ID or the Amazon Resource Name (ARN) of the Outpost.")
	outpostsCmd.AddCommand(outposts_listOrdersCmd)
}
