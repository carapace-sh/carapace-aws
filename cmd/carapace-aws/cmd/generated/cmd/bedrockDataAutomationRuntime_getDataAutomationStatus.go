package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomationRuntime_getDataAutomationStatusCmd = &cobra.Command{
	Use:   "get-data-automation-status",
	Short: "API used to get data automation status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomationRuntime_getDataAutomationStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockDataAutomationRuntime_getDataAutomationStatusCmd).Standalone()

		bedrockDataAutomationRuntime_getDataAutomationStatusCmd.Flags().String("invocation-arn", "", "Invocation arn.")
		bedrockDataAutomationRuntime_getDataAutomationStatusCmd.MarkFlagRequired("invocation-arn")
	})
	bedrockDataAutomationRuntimeCmd.AddCommand(bedrockDataAutomationRuntime_getDataAutomationStatusCmd)
}
