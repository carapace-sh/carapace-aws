package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Adds an application that is created from a resource group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_createApplicationCmd).Standalone()

		applicationInsights_createApplicationCmd.Flags().String("attach-missing-permission", "", "If set to true, the managed policies for SSM and CW will be attached to the instance roles if they are missing.")
		applicationInsights_createApplicationCmd.Flags().String("auto-config-enabled", "", "Indicates whether Application Insights automatically configures unmonitored resources in the resource group.")
		applicationInsights_createApplicationCmd.Flags().String("auto-create", "", "Configures all of the resources in the resource group by applying the recommended configurations.")
		applicationInsights_createApplicationCmd.Flags().String("cwemonitor-enabled", "", "Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as `instance terminated`, `failed deployment`, and others.")
		applicationInsights_createApplicationCmd.Flags().String("grouping-type", "", "Application Insights can create applications based on a resource group or on an account.")
		applicationInsights_createApplicationCmd.Flags().String("ops-center-enabled", "", "When set to `true`, creates opsItems for any problems detected on an application.")
		applicationInsights_createApplicationCmd.Flags().String("ops-item-snstopic-arn", "", "The SNS topic provided to Application Insights that is associated to the created opsItem.")
		applicationInsights_createApplicationCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
		applicationInsights_createApplicationCmd.Flags().String("snsnotification-arn", "", "The SNS notification topic ARN.")
		applicationInsights_createApplicationCmd.Flags().String("tags", "", "List of tags to add to the application.")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_createApplicationCmd)
}
