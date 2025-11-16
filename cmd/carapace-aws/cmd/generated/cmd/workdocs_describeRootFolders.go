package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_describeRootFoldersCmd = &cobra.Command{
	Use:   "describe-root-folders",
	Short: "Describes the current user's special folders; the `RootFolder` and the `RecycleBin`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_describeRootFoldersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_describeRootFoldersCmd).Standalone()

		workdocs_describeRootFoldersCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_describeRootFoldersCmd.Flags().String("limit", "", "The maximum number of items to return.")
		workdocs_describeRootFoldersCmd.Flags().String("marker", "", "The marker for the next set of results.")
		workdocs_describeRootFoldersCmd.MarkFlagRequired("authentication-token")
	})
	workdocsCmd.AddCommand(workdocs_describeRootFoldersCmd)
}
