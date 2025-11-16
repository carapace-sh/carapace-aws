package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_createTagsCmd = &cobra.Command{
	Use:   "create-tags",
	Short: "DEPRECATED - `CreateTags` is deprecated and not maintained.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_createTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_createTagsCmd).Standalone()

		efs_createTagsCmd.Flags().String("file-system-id", "", "The ID of the file system whose tags you want to modify (String).")
		efs_createTagsCmd.Flags().String("tags", "", "An array of `Tag` objects to add.")
		efs_createTagsCmd.MarkFlagRequired("file-system-id")
		efs_createTagsCmd.MarkFlagRequired("tags")
	})
	efsCmd.AddCommand(efs_createTagsCmd)
}
