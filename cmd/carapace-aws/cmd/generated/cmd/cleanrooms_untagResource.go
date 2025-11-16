package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag or list of tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_untagResourceCmd).Standalone()

		cleanrooms_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) associated with the resource you want to remove the tag from.")
		cleanrooms_untagResourceCmd.Flags().String("tag-keys", "", "A list of key names of tags to be removed.")
		cleanrooms_untagResourceCmd.MarkFlagRequired("resource-arn")
		cleanrooms_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	cleanroomsCmd.AddCommand(cleanrooms_untagResourceCmd)
}
