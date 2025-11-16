package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_deleteFolderCmd = &cobra.Command{
	Use:   "delete-folder",
	Short: "Permanently deletes the specified folder and its contents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_deleteFolderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_deleteFolderCmd).Standalone()

		workdocs_deleteFolderCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_deleteFolderCmd.Flags().String("folder-id", "", "The ID of the folder.")
		workdocs_deleteFolderCmd.MarkFlagRequired("folder-id")
	})
	workdocsCmd.AddCommand(workdocs_deleteFolderCmd)
}
