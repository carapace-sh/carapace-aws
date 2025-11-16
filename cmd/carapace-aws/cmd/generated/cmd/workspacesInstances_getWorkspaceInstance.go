package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_getWorkspaceInstanceCmd = &cobra.Command{
	Use:   "get-workspace-instance",
	Short: "Retrieves detailed information about a specific WorkSpace Instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_getWorkspaceInstanceCmd).Standalone()

	workspacesInstances_getWorkspaceInstanceCmd.Flags().String("workspace-instance-id", "", "Unique identifier of the WorkSpace Instance to retrieve.")
	workspacesInstances_getWorkspaceInstanceCmd.MarkFlagRequired("workspace-instance-id")
	workspacesInstancesCmd.AddCommand(workspacesInstances_getWorkspaceInstanceCmd)
}
