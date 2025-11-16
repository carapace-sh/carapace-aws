package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associate tags with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_tagResourceCmd).Standalone()

		bedrockAgentRuntime_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to tag.")
		bedrockAgentRuntime_tagResourceCmd.Flags().String("tags", "", "An object containing key-value pairs that define the tags to attach to the resource.")
		bedrockAgentRuntime_tagResourceCmd.MarkFlagRequired("resource-arn")
		bedrockAgentRuntime_tagResourceCmd.MarkFlagRequired("tags")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_tagResourceCmd)
}
