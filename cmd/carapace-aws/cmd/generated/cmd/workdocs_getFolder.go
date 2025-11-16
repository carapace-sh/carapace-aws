package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_getFolderCmd = &cobra.Command{
	Use:   "get-folder",
	Short: "Retrieves the metadata of the specified folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_getFolderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_getFolderCmd).Standalone()

		workdocs_getFolderCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_getFolderCmd.Flags().String("folder-id", "", "The ID of the folder.")
		workdocs_getFolderCmd.Flags().String("include-custom-metadata", "", "Set to TRUE to include custom metadata in the response.")
		workdocs_getFolderCmd.MarkFlagRequired("folder-id")
	})
	workdocsCmd.AddCommand(workdocs_getFolderCmd)
}
