package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_createWorkspaceInstanceCmd = &cobra.Command{
	Use:   "create-workspace-instance",
	Short: "Launches a new WorkSpace Instance with specified configuration parameters, enabling programmatic workspace deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_createWorkspaceInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesInstances_createWorkspaceInstanceCmd).Standalone()

		workspacesInstances_createWorkspaceInstanceCmd.Flags().String("client-token", "", "Unique token to ensure idempotent instance creation, preventing duplicate workspace launches.")
		workspacesInstances_createWorkspaceInstanceCmd.Flags().String("managed-instance", "", "Comprehensive configuration settings for the WorkSpaces Instance, including network, compute, and storage parameters.")
		workspacesInstances_createWorkspaceInstanceCmd.Flags().String("tags", "", "Optional metadata tags for categorizing and managing WorkSpaces Instances.")
		workspacesInstances_createWorkspaceInstanceCmd.MarkFlagRequired("managed-instance")
	})
	workspacesInstancesCmd.AddCommand(workspacesInstances_createWorkspaceInstanceCmd)
}
