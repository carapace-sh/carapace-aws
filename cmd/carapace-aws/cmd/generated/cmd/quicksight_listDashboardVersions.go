package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listDashboardVersionsCmd = &cobra.Command{
	Use:   "list-dashboard-versions",
	Short: "Lists all the versions of the dashboards in the Amazon Quick Sight subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listDashboardVersionsCmd).Standalone()

	quicksight_listDashboardVersionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard that you're listing versions for.")
	quicksight_listDashboardVersionsCmd.Flags().String("dashboard-id", "", "The ID for the dashboard.")
	quicksight_listDashboardVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_listDashboardVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no more results.")
	quicksight_listDashboardVersionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_listDashboardVersionsCmd.MarkFlagRequired("dashboard-id")
	quicksightCmd.AddCommand(quicksight_listDashboardVersionsCmd)
}
