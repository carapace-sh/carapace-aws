package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the specified container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_untagResourceCmd).Standalone()

	mediastore_untagResourceCmd.Flags().String("resource", "", "The Amazon Resource Name (ARN) for the container.")
	mediastore_untagResourceCmd.Flags().String("tag-keys", "", "A comma-separated list of keys for tags that you want to remove from the container.")
	mediastore_untagResourceCmd.MarkFlagRequired("resource")
	mediastore_untagResourceCmd.MarkFlagRequired("tag-keys")
	mediastoreCmd.AddCommand(mediastore_untagResourceCmd)
}
