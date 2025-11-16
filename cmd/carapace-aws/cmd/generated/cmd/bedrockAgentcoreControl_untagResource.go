package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_untagResourceCmd).Standalone()

		bedrockAgentcoreControl_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to untag.")
		bedrockAgentcoreControl_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys of the tags to remove from the resource.")
		bedrockAgentcoreControl_untagResourceCmd.MarkFlagRequired("resource-arn")
		bedrockAgentcoreControl_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_untagResourceCmd)
}
