package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to a resource with the specified resourceArn.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_tagResourceCmd).Standalone()

	bedrockAgentcoreControl_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
	bedrockAgentcoreControl_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
	bedrockAgentcoreControl_tagResourceCmd.MarkFlagRequired("resource-arn")
	bedrockAgentcoreControl_tagResourceCmd.MarkFlagRequired("tags")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_tagResourceCmd)
}
