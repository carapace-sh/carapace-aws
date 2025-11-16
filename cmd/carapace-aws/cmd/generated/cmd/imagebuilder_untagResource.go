package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_untagResourceCmd).Standalone()

		imagebuilder_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to untag.")
		imagebuilder_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to remove from the resource.")
		imagebuilder_untagResourceCmd.MarkFlagRequired("resource-arn")
		imagebuilder_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	imagebuilderCmd.AddCommand(imagebuilder_untagResourceCmd)
}
