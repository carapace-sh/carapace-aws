package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_describeFolderContentsCmd = &cobra.Command{
	Use:   "describe-folder-contents",
	Short: "Describes the contents of the specified folder, including its documents and subfolders.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_describeFolderContentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_describeFolderContentsCmd).Standalone()

		workdocs_describeFolderContentsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_describeFolderContentsCmd.Flags().String("folder-id", "", "The ID of the folder.")
		workdocs_describeFolderContentsCmd.Flags().String("include", "", "The contents to include.")
		workdocs_describeFolderContentsCmd.Flags().String("limit", "", "The maximum number of items to return with this call.")
		workdocs_describeFolderContentsCmd.Flags().String("marker", "", "The marker for the next set of results.")
		workdocs_describeFolderContentsCmd.Flags().String("order", "", "The order for the contents of the folder.")
		workdocs_describeFolderContentsCmd.Flags().String("sort", "", "The sorting criteria.")
		workdocs_describeFolderContentsCmd.Flags().String("type", "", "The type of items.")
		workdocs_describeFolderContentsCmd.MarkFlagRequired("folder-id")
	})
	workdocsCmd.AddCommand(workdocs_describeFolderContentsCmd)
}
