package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "The resource to untag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_untagResourceCmd).Standalone()

		mediatailor_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to untag.")
		mediatailor_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys associated with the resource.")
		mediatailor_untagResourceCmd.MarkFlagRequired("resource-arn")
		mediatailor_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	mediatailorCmd.AddCommand(mediatailor_untagResourceCmd)
}
