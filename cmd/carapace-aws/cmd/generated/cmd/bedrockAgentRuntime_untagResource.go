package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_untagResourceCmd).Standalone()

		bedrockAgentRuntime_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which to remove tags.")
		bedrockAgentRuntime_untagResourceCmd.Flags().String("tag-keys", "", "A list of keys of the tags to remove from the resource.")
		bedrockAgentRuntime_untagResourceCmd.MarkFlagRequired("resource-arn")
		bedrockAgentRuntime_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_untagResourceCmd)
}
