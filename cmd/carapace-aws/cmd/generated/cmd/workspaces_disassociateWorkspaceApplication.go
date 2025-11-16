package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_disassociateWorkspaceApplicationCmd = &cobra.Command{
	Use:   "disassociate-workspace-application",
	Short: "Disassociates the specified application from a WorkSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_disassociateWorkspaceApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_disassociateWorkspaceApplicationCmd).Standalone()

		workspaces_disassociateWorkspaceApplicationCmd.Flags().String("application-id", "", "The identifier of the application.")
		workspaces_disassociateWorkspaceApplicationCmd.Flags().String("workspace-id", "", "The identifier of the WorkSpace.")
		workspaces_disassociateWorkspaceApplicationCmd.MarkFlagRequired("application-id")
		workspaces_disassociateWorkspaceApplicationCmd.MarkFlagRequired("workspace-id")
	})
	workspacesCmd.AddCommand(workspaces_disassociateWorkspaceApplicationCmd)
}
