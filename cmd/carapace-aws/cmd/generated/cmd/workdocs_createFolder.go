package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_createFolderCmd = &cobra.Command{
	Use:   "create-folder",
	Short: "Creates a folder with the specified name and parent folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_createFolderCmd).Standalone()

	workdocs_createFolderCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_createFolderCmd.Flags().String("name", "", "The name of the new folder.")
	workdocs_createFolderCmd.Flags().String("parent-folder-id", "", "The ID of the parent folder.")
	workdocs_createFolderCmd.MarkFlagRequired("parent-folder-id")
	workdocsCmd.AddCommand(workdocs_createFolderCmd)
}
