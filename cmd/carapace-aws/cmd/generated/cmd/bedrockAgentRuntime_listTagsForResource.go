package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all the tags for the resource you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_listTagsForResourceCmd).Standalone()

		bedrockAgentRuntime_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which to list tags.")
		bedrockAgentRuntime_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_listTagsForResourceCmd)
}
