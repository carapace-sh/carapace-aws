package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_untagResourceCmd).Standalone()

	panorama_untagResourceCmd.Flags().String("resource-arn", "", "The resource's ARN.")
	panorama_untagResourceCmd.Flags().String("tag-keys", "", "Tag keys to remove.")
	panorama_untagResourceCmd.MarkFlagRequired("resource-arn")
	panorama_untagResourceCmd.MarkFlagRequired("tag-keys")
	panoramaCmd.AddCommand(panorama_untagResourceCmd)
}
