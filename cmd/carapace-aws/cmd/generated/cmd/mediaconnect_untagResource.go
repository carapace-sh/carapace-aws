package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_untagResourceCmd).Standalone()

		mediaconnect_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to untag.")
		mediaconnect_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to be removed.")
		mediaconnect_untagResourceCmd.MarkFlagRequired("resource-arn")
		mediaconnect_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	mediaconnectCmd.AddCommand(mediaconnect_untagResourceCmd)
}
