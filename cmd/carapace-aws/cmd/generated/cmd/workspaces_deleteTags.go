package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_deleteTagsCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Deletes the specified tags from the specified WorkSpaces resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_deleteTagsCmd).Standalone()

	workspaces_deleteTagsCmd.Flags().String("resource-id", "", "The identifier of the WorkSpaces resource.")
	workspaces_deleteTagsCmd.Flags().String("tag-keys", "", "The tag keys.")
	workspaces_deleteTagsCmd.MarkFlagRequired("resource-id")
	workspaces_deleteTagsCmd.MarkFlagRequired("tag-keys")
	workspacesCmd.AddCommand(workspaces_deleteTagsCmd)
}
