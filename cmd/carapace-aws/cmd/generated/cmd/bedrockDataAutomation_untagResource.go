package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untag an Amazon Bedrock Data Automation resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockDataAutomation_untagResourceCmd).Standalone()

		bedrockDataAutomation_untagResourceCmd.Flags().String("resource-arn", "", "")
		bedrockDataAutomation_untagResourceCmd.Flags().String("tag-keys", "", "")
		bedrockDataAutomation_untagResourceCmd.MarkFlagRequired("resource-arn")
		bedrockDataAutomation_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_untagResourceCmd)
}
