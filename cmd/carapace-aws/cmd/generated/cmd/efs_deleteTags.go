package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_deleteTagsCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "DEPRECATED - `DeleteTags` is deprecated and not maintained.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_deleteTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_deleteTagsCmd).Standalone()

		efs_deleteTagsCmd.Flags().String("file-system-id", "", "The ID of the file system whose tags you want to delete (String).")
		efs_deleteTagsCmd.Flags().String("tag-keys", "", "A list of tag keys to delete.")
		efs_deleteTagsCmd.MarkFlagRequired("file-system-id")
		efs_deleteTagsCmd.MarkFlagRequired("tag-keys")
	})
	efsCmd.AddCommand(efs_deleteTagsCmd)
}
