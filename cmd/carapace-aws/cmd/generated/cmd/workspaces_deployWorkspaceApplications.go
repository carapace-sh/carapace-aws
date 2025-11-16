package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_deployWorkspaceApplicationsCmd = &cobra.Command{
	Use:   "deploy-workspace-applications",
	Short: "Deploys associated applications to the specified WorkSpace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_deployWorkspaceApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_deployWorkspaceApplicationsCmd).Standalone()

		workspaces_deployWorkspaceApplicationsCmd.Flags().String("force", "", "Indicates whether the force flag is applied for the specified WorkSpace.")
		workspaces_deployWorkspaceApplicationsCmd.Flags().String("workspace-id", "", "The identifier of the WorkSpace.")
		workspaces_deployWorkspaceApplicationsCmd.MarkFlagRequired("workspace-id")
	})
	workspacesCmd.AddCommand(workspaces_deployWorkspaceApplicationsCmd)
}
