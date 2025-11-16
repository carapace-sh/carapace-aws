package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_deleteFolderContentsCmd = &cobra.Command{
	Use:   "delete-folder-contents",
	Short: "Deletes the contents of the specified folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_deleteFolderContentsCmd).Standalone()

	workdocs_deleteFolderContentsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_deleteFolderContentsCmd.Flags().String("folder-id", "", "The ID of the folder.")
	workdocs_deleteFolderContentsCmd.MarkFlagRequired("folder-id")
	workdocsCmd.AddCommand(workdocs_deleteFolderContentsCmd)
}
