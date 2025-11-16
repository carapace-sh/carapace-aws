package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_updateApplicationCmd).Standalone()

	applicationInsights_updateApplicationCmd.Flags().String("attach-missing-permission", "", "If set to true, the managed policies for SSM and CW will be attached to the instance roles if they are missing.")
	applicationInsights_updateApplicationCmd.Flags().String("auto-config-enabled", "", "Turns auto-configuration on or off.")
	applicationInsights_updateApplicationCmd.Flags().String("cwemonitor-enabled", "", "Indicates whether Application Insights can listen to CloudWatch events for the application resources, such as `instance terminated`, `failed deployment`, and others.")
	applicationInsights_updateApplicationCmd.Flags().String("ops-center-enabled", "", "When set to `true`, creates opsItems for any problems detected on an application.")
	applicationInsights_updateApplicationCmd.Flags().String("ops-item-snstopic-arn", "", "The SNS topic provided to Application Insights that is associated to the created opsItem.")
	applicationInsights_updateApplicationCmd.Flags().String("remove-snstopic", "", "Disassociates the SNS topic from the opsItem created for detected problems.")
	applicationInsights_updateApplicationCmd.Flags().String("resource-group-name", "", "The name of the resource group.")
	applicationInsights_updateApplicationCmd.Flags().String("snsnotification-arn", "", "The SNS topic ARN.")
	applicationInsights_updateApplicationCmd.MarkFlagRequired("resource-group-name")
	applicationInsightsCmd.AddCommand(applicationInsights_updateApplicationCmd)
}
