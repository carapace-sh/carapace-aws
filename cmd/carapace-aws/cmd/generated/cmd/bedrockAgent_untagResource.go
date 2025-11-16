package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_untagResourceCmd).Standalone()

		bedrockAgent_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which to remove tags.")
		bedrockAgent_untagResourceCmd.Flags().String("tag-keys", "", "A list of keys of the tags to remove from the resource.")
		bedrockAgent_untagResourceCmd.MarkFlagRequired("resource-arn")
		bedrockAgent_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_untagResourceCmd)
}
