package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomationRuntime_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tag an Amazon Bedrock Data Automation resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomationRuntime_tagResourceCmd).Standalone()

	bedrockDataAutomationRuntime_tagResourceCmd.Flags().String("resource-arn", "", "")
	bedrockDataAutomationRuntime_tagResourceCmd.Flags().String("tags", "", "")
	bedrockDataAutomationRuntime_tagResourceCmd.MarkFlagRequired("resource-arn")
	bedrockDataAutomationRuntime_tagResourceCmd.MarkFlagRequired("tags")
	bedrockDataAutomationRuntimeCmd.AddCommand(bedrockDataAutomationRuntime_tagResourceCmd)
}
