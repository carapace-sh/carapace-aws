package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_listWorkspaceInstancesCmd = &cobra.Command{
	Use:   "list-workspace-instances",
	Short: "Retrieves a collection of WorkSpaces Instances based on specified filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_listWorkspaceInstancesCmd).Standalone()

	workspacesInstances_listWorkspaceInstancesCmd.Flags().String("max-results", "", "Maximum number of WorkSpaces Instances to return in a single response.")
	workspacesInstances_listWorkspaceInstancesCmd.Flags().String("next-token", "", "Pagination token for retrieving subsequent pages of WorkSpaces Instances.")
	workspacesInstances_listWorkspaceInstancesCmd.Flags().String("provision-states", "", "Filter WorkSpaces Instances by their current provisioning states.")
	workspacesInstancesCmd.AddCommand(workspacesInstances_listWorkspaceInstancesCmd)
}
