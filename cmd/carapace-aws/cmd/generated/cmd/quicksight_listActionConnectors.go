package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listActionConnectorsCmd = &cobra.Command{
	Use:   "list-action-connectors",
	Short: "Lists all action connectors in the specified Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listActionConnectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_listActionConnectorsCmd).Standalone()

		quicksight_listActionConnectorsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID for which to list action connectors.")
		quicksight_listActionConnectorsCmd.Flags().String("max-results", "", "The maximum number of action connectors to return in a single response.")
		quicksight_listActionConnectorsCmd.Flags().String("next-token", "", "A pagination token to retrieve the next set of results.")
		quicksight_listActionConnectorsCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_listActionConnectorsCmd)
}
