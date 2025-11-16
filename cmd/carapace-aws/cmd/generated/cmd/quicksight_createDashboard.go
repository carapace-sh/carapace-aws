package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createDashboardCmd = &cobra.Command{
	Use:   "create-dashboard",
	Short: "Creates a dashboard from either a template or directly with a `DashboardDefinition`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createDashboardCmd).Standalone()

		quicksight_createDashboardCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account where you want to create the dashboard.")
		quicksight_createDashboardCmd.Flags().String("dashboard-id", "", "The ID for the dashboard, also added to the IAM policy.")
		quicksight_createDashboardCmd.Flags().String("dashboard-publish-options", "", "Options for publishing the dashboard when you create it:")
		quicksight_createDashboardCmd.Flags().String("definition", "", "The definition of a dashboard.")
		quicksight_createDashboardCmd.Flags().String("folder-arns", "", "When you create the dashboard, Amazon Quick Sight adds the dashboard to these folders.")
		quicksight_createDashboardCmd.Flags().String("link-entities", "", "A list of analysis Amazon Resource Names (ARNs) to be linked to the dashboard.")
		quicksight_createDashboardCmd.Flags().String("link-sharing-configuration", "", "A structure that contains the permissions of a shareable link to the dashboard.")
		quicksight_createDashboardCmd.Flags().String("name", "", "The display name of the dashboard.")
		quicksight_createDashboardCmd.Flags().String("parameters", "", "The parameters for the creation of the dashboard, which you want to use to override the default settings.")
		quicksight_createDashboardCmd.Flags().String("permissions", "", "A structure that contains the permissions of the dashboard.")
		quicksight_createDashboardCmd.Flags().String("source-entity", "", "The entity that you are using as a source when you create the dashboard.")
		quicksight_createDashboardCmd.Flags().String("tags", "", "Contains a map of the key-value pairs for the resource tag or tags assigned to the dashboard.")
		quicksight_createDashboardCmd.Flags().String("theme-arn", "", "The Amazon Resource Name (ARN) of the theme that is being used for this dashboard.")
		quicksight_createDashboardCmd.Flags().String("validation-strategy", "", "The option to relax the validation needed to create a dashboard with definition objects.")
		quicksight_createDashboardCmd.Flags().String("version-description", "", "A description for the first version of the dashboard being created.")
		quicksight_createDashboardCmd.MarkFlagRequired("aws-account-id")
		quicksight_createDashboardCmd.MarkFlagRequired("dashboard-id")
		quicksight_createDashboardCmd.MarkFlagRequired("name")
	})
	quicksightCmd.AddCommand(quicksight_createDashboardCmd)
}
