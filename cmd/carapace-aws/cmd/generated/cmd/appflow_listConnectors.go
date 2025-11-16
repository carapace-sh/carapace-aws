package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_listConnectorsCmd = &cobra.Command{
	Use:   "list-connectors",
	Short: "Returns the list of all registered custom connectors in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_listConnectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_listConnectorsCmd).Standalone()

		appflow_listConnectorsCmd.Flags().String("max-results", "", "Specifies the maximum number of items that should be returned in the result set.")
		appflow_listConnectorsCmd.Flags().String("next-token", "", "The pagination token for the next page of data.")
	})
	appflowCmd.AddCommand(appflow_listConnectorsCmd)
}
