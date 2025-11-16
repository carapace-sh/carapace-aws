package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_untagResourceCmd).Standalone()

	mediapackageVod_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource.")
	mediapackageVod_untagResourceCmd.Flags().String("tag-keys", "", "A comma-separated list of the tag keys to remove from the resource.")
	mediapackageVod_untagResourceCmd.MarkFlagRequired("resource-arn")
	mediapackageVod_untagResourceCmd.MarkFlagRequired("tag-keys")
	mediapackageVodCmd.AddCommand(mediapackageVod_untagResourceCmd)
}
