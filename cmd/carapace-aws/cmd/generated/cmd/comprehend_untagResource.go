package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a specific tag associated with an Amazon Comprehend resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_untagResourceCmd).Standalone()

		comprehend_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the given Amazon Comprehend resource from which you want to remove the tags.")
		comprehend_untagResourceCmd.Flags().String("tag-keys", "", "The initial part of a key-value pair that forms a tag being removed from a given resource.")
		comprehend_untagResourceCmd.MarkFlagRequired("resource-arn")
		comprehend_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	comprehendCmd.AddCommand(comprehend_untagResourceCmd)
}
