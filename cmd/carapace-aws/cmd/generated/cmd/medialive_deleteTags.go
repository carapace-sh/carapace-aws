package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteTagsCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Removes tags for a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_deleteTagsCmd).Standalone()

		medialive_deleteTagsCmd.Flags().String("resource-arn", "", "")
		medialive_deleteTagsCmd.Flags().String("tag-keys", "", "An array of tag keys to delete")
		medialive_deleteTagsCmd.MarkFlagRequired("resource-arn")
		medialive_deleteTagsCmd.MarkFlagRequired("tag-keys")
	})
	medialiveCmd.AddCommand(medialive_deleteTagsCmd)
}
