package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomationRuntime_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untag an Amazon Bedrock Data Automation resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomationRuntime_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockDataAutomationRuntime_untagResourceCmd).Standalone()

		bedrockDataAutomationRuntime_untagResourceCmd.Flags().String("resource-arn", "", "")
		bedrockDataAutomationRuntime_untagResourceCmd.Flags().String("tag-keys", "", "")
		bedrockDataAutomationRuntime_untagResourceCmd.MarkFlagRequired("resource-arn")
		bedrockDataAutomationRuntime_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	bedrockDataAutomationRuntimeCmd.AddCommand(bedrockDataAutomationRuntime_untagResourceCmd)
}
