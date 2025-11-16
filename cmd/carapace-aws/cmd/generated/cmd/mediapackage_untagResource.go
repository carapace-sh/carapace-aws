package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_untagResourceCmd).Standalone()

	mediapackage_untagResourceCmd.Flags().String("resource-arn", "", "")
	mediapackage_untagResourceCmd.Flags().String("tag-keys", "", "The key(s) of tag to be deleted")
	mediapackage_untagResourceCmd.MarkFlagRequired("resource-arn")
	mediapackage_untagResourceCmd.MarkFlagRequired("tag-keys")
	mediapackageCmd.AddCommand(mediapackage_untagResourceCmd)
}
