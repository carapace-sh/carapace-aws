package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listDashboardsCmd = &cobra.Command{
	Use:   "list-dashboards",
	Short: "Lists dashboards in an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listDashboardsCmd).Standalone()

	quicksight_listDashboardsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboards that you're listing.")
	quicksight_listDashboardsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_listDashboardsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listDashboardsCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_listDashboardsCmd)
}
