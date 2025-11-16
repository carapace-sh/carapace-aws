package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_deleteWorkspaceInstanceCmd = &cobra.Command{
	Use:   "delete-workspace-instance",
	Short: "Deletes the specified WorkSpace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_deleteWorkspaceInstanceCmd).Standalone()

	workspacesInstances_deleteWorkspaceInstanceCmd.Flags().String("workspace-instance-id", "", "Unique identifier of the WorkSpaces Instance targeted for deletion.")
	workspacesInstances_deleteWorkspaceInstanceCmd.MarkFlagRequired("workspace-instance-id")
	workspacesInstancesCmd.AddCommand(workspacesInstances_deleteWorkspaceInstanceCmd)
}
