package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateDashboardCmd = &cobra.Command{
	Use:   "update-dashboard",
	Short: "Updates a dashboard in an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateDashboardCmd).Standalone()

	quicksight_updateDashboardCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard that you're updating.")
	quicksight_updateDashboardCmd.Flags().String("dashboard-id", "", "The ID for the dashboard.")
	quicksight_updateDashboardCmd.Flags().String("dashboard-publish-options", "", "Options for publishing the dashboard when you create it:")
	quicksight_updateDashboardCmd.Flags().String("definition", "", "The definition of a dashboard.")
	quicksight_updateDashboardCmd.Flags().String("name", "", "The display name of the dashboard.")
	quicksight_updateDashboardCmd.Flags().String("parameters", "", "A structure that contains the parameters of the dashboard.")
	quicksight_updateDashboardCmd.Flags().String("source-entity", "", "The entity that you are using as a source when you update the dashboard.")
	quicksight_updateDashboardCmd.Flags().String("theme-arn", "", "The Amazon Resource Name (ARN) of the theme that is being used for this dashboard.")
	quicksight_updateDashboardCmd.Flags().String("validation-strategy", "", "The option to relax the validation needed to update a dashboard with definition objects.")
	quicksight_updateDashboardCmd.Flags().String("version-description", "", "A description for the first version of the dashboard being created.")
	quicksight_updateDashboardCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateDashboardCmd.MarkFlagRequired("dashboard-id")
	quicksight_updateDashboardCmd.MarkFlagRequired("name")
	quicksightCmd.AddCommand(quicksight_updateDashboardCmd)
}
