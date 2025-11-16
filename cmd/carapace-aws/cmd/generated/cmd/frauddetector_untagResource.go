package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_untagResourceCmd).Standalone()

	frauddetector_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource from which to remove the tag.")
	frauddetector_untagResourceCmd.Flags().String("tag-keys", "", "The resource ARN.")
	frauddetector_untagResourceCmd.MarkFlagRequired("resource-arn")
	frauddetector_untagResourceCmd.MarkFlagRequired("tag-keys")
	frauddetectorCmd.AddCommand(frauddetector_untagResourceCmd)
}
