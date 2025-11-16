package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tag an Amazon Bedrock Data Automation resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_tagResourceCmd).Standalone()

	bedrockDataAutomation_tagResourceCmd.Flags().String("resource-arn", "", "")
	bedrockDataAutomation_tagResourceCmd.Flags().String("tags", "", "")
	bedrockDataAutomation_tagResourceCmd.MarkFlagRequired("resource-arn")
	bedrockDataAutomation_tagResourceCmd.MarkFlagRequired("tags")
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_tagResourceCmd)
}
