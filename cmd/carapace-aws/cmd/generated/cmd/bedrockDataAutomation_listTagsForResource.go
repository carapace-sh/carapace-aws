package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockDataAutomation_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List tags for an Amazon Bedrock Data Automation resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockDataAutomation_listTagsForResourceCmd).Standalone()

	bedrockDataAutomation_listTagsForResourceCmd.Flags().String("resource-arn", "", "")
	bedrockDataAutomation_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	bedrockDataAutomationCmd.AddCommand(bedrockDataAutomation_listTagsForResourceCmd)
}
