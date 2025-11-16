package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_associateWorkspaceApplicationCmd = &cobra.Command{
	Use:   "associate-workspace-application",
	Short: "Associates the specified application to the specified WorkSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_associateWorkspaceApplicationCmd).Standalone()

	workspaces_associateWorkspaceApplicationCmd.Flags().String("application-id", "", "The identifier of the application.")
	workspaces_associateWorkspaceApplicationCmd.Flags().String("workspace-id", "", "The identifier of the WorkSpace.")
	workspaces_associateWorkspaceApplicationCmd.MarkFlagRequired("application-id")
	workspaces_associateWorkspaceApplicationCmd.MarkFlagRequired("workspace-id")
	workspacesCmd.AddCommand(workspaces_associateWorkspaceApplicationCmd)
}
