package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listConnectorsCmd = &cobra.Command{
	Use:   "list-connectors",
	Short: "Lists the connectors for the specified Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listConnectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_listConnectorsCmd).Standalone()

		transfer_listConnectorsCmd.Flags().String("max-results", "", "The maximum number of items to return.")
		transfer_listConnectorsCmd.Flags().String("next-token", "", "When you can get additional results from the `ListConnectors` call, a `NextToken` parameter is returned in the output.")
	})
	transferCmd.AddCommand(transfer_listConnectorsCmd)
}
