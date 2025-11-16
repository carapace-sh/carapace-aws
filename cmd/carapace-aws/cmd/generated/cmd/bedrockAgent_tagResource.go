package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associate tags with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_tagResourceCmd).Standalone()

	bedrockAgent_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to tag.")
	bedrockAgent_tagResourceCmd.Flags().String("tags", "", "An object containing key-value pairs that define the tags to attach to the resource.")
	bedrockAgent_tagResourceCmd.MarkFlagRequired("resource-arn")
	bedrockAgent_tagResourceCmd.MarkFlagRequired("tags")
	bedrockAgentCmd.AddCommand(bedrockAgent_tagResourceCmd)
}
