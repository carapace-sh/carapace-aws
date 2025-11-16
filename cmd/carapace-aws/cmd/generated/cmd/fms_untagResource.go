package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from an Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_untagResourceCmd).Standalone()

		fms_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to return tags for.")
		fms_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to remove from the resource.")
		fms_untagResourceCmd.MarkFlagRequired("resource-arn")
		fms_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	fmsCmd.AddCommand(fms_untagResourceCmd)
}
