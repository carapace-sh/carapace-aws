package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a specific tag associated with an Amazon Translate resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(translate_untagResourceCmd).Standalone()

		translate_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the given Amazon Translate resource from which you want to remove the tags.")
		translate_untagResourceCmd.Flags().String("tag-keys", "", "The initial part of a key-value pair that forms a tag being removed from a given resource.")
		translate_untagResourceCmd.MarkFlagRequired("resource-arn")
		translate_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	translateCmd.AddCommand(translate_untagResourceCmd)
}
