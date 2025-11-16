package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_getFolderPathCmd = &cobra.Command{
	Use:   "get-folder-path",
	Short: "Retrieves the path information (the hierarchy from the root folder) for the specified folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_getFolderPathCmd).Standalone()

	workdocs_getFolderPathCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_getFolderPathCmd.Flags().String("fields", "", "A comma-separated list of values.")
	workdocs_getFolderPathCmd.Flags().String("folder-id", "", "The ID of the folder.")
	workdocs_getFolderPathCmd.Flags().String("limit", "", "The maximum number of levels in the hierarchy to return.")
	workdocs_getFolderPathCmd.Flags().String("marker", "", "This value is not supported.")
	workdocs_getFolderPathCmd.MarkFlagRequired("folder-id")
	workdocsCmd.AddCommand(workdocs_getFolderPathCmd)
}
