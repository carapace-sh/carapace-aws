package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_describeWorkspaceDirectoriesCmd = &cobra.Command{
	Use:   "describe-workspace-directories",
	Short: "Describes the available directories that are registered with Amazon WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_describeWorkspaceDirectoriesCmd).Standalone()

	workspaces_describeWorkspaceDirectoriesCmd.Flags().String("directory-ids", "", "The identifiers of the directories.")
	workspaces_describeWorkspaceDirectoriesCmd.Flags().String("filters", "", "The filter condition for the WorkSpaces.")
	workspaces_describeWorkspaceDirectoriesCmd.Flags().String("limit", "", "The maximum number of directories to return.")
	workspaces_describeWorkspaceDirectoriesCmd.Flags().String("next-token", "", "If you received a `NextToken` from a previous call that was paginated, provide this token to receive the next set of results.")
	workspaces_describeWorkspaceDirectoriesCmd.Flags().String("workspace-directory-names", "", "The names of the WorkSpace directories.")
	workspacesCmd.AddCommand(workspaces_describeWorkspaceDirectoriesCmd)
}
