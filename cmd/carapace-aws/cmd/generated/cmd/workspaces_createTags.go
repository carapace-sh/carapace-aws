package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createTagsCmd = &cobra.Command{
	Use:   "create-tags",
	Short: "Creates the specified tags for the specified WorkSpaces resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_createTagsCmd).Standalone()

		workspaces_createTagsCmd.Flags().String("resource-id", "", "The identifier of the WorkSpaces resource.")
		workspaces_createTagsCmd.Flags().String("tags", "", "The tags.")
		workspaces_createTagsCmd.MarkFlagRequired("resource-id")
		workspaces_createTagsCmd.MarkFlagRequired("tags")
	})
	workspacesCmd.AddCommand(workspaces_createTagsCmd)
}
