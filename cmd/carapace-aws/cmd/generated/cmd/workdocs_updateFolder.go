package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_updateFolderCmd = &cobra.Command{
	Use:   "update-folder",
	Short: "Updates the specified attributes of the specified folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_updateFolderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_updateFolderCmd).Standalone()

		workdocs_updateFolderCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_updateFolderCmd.Flags().String("folder-id", "", "The ID of the folder.")
		workdocs_updateFolderCmd.Flags().String("name", "", "The name of the folder.")
		workdocs_updateFolderCmd.Flags().String("parent-folder-id", "", "The ID of the parent folder.")
		workdocs_updateFolderCmd.Flags().String("resource-state", "", "The resource state of the folder.")
		workdocs_updateFolderCmd.MarkFlagRequired("folder-id")
	})
	workdocsCmd.AddCommand(workdocs_updateFolderCmd)
}
