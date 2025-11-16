package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_untagResourceCmd).Standalone()

	mediapackagev2_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the MediaPackage resource that you're removing tags from.")
	mediapackagev2_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
	mediapackagev2_untagResourceCmd.MarkFlagRequired("resource-arn")
	mediapackagev2_untagResourceCmd.MarkFlagRequired("tag-keys")
	mediapackagev2Cmd.AddCommand(mediapackagev2_untagResourceCmd)
}
