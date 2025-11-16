package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd = &cobra.Command{
	Use:   "invoke-data-automation-async",
	Short: "Async API: Invoke data automation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd).Standalone()

	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.Flags().String("blueprints", "", "Blueprint list.")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.Flags().String("client-token", "", "Idempotency token.")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.Flags().String("data-automation-configuration", "", "Data automation configuration.")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.Flags().String("data-automation-profile-arn", "", "Data automation profile ARN")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.Flags().String("encryption-configuration", "", "Encryption configuration.")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.Flags().String("input-configuration", "", "Input configuration.")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.Flags().String("notification-configuration", "", "Notification configuration.")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.Flags().String("output-configuration", "", "Output configuration.")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.Flags().String("tags", "", "List of tags.")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.MarkFlagRequired("data-automation-profile-arn")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.MarkFlagRequired("input-configuration")
	bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd.MarkFlagRequired("output-configuration")
	bedrockDataAutomationRuntimeCmd.AddCommand(bedrockDataAutomationRuntime_invokeDataAutomationAsyncCmd)
}
