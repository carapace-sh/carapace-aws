package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_createWorkspacesPoolCmd = &cobra.Command{
	Use:   "create-workspaces-pool",
	Short: "Creates a pool of WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_createWorkspacesPoolCmd).Standalone()

	workspaces_createWorkspacesPoolCmd.Flags().String("application-settings", "", "Indicates the application settings of the pool.")
	workspaces_createWorkspacesPoolCmd.Flags().String("bundle-id", "", "The identifier of the bundle for the pool.")
	workspaces_createWorkspacesPoolCmd.Flags().String("capacity", "", "The user capacity of the pool.")
	workspaces_createWorkspacesPoolCmd.Flags().String("description", "", "The pool description.")
	workspaces_createWorkspacesPoolCmd.Flags().String("directory-id", "", "The identifier of the directory for the pool.")
	workspaces_createWorkspacesPoolCmd.Flags().String("pool-name", "", "The name of the pool.")
	workspaces_createWorkspacesPoolCmd.Flags().String("running-mode", "", "The running mode for the pool.")
	workspaces_createWorkspacesPoolCmd.Flags().String("tags", "", "The tags for the pool.")
	workspaces_createWorkspacesPoolCmd.Flags().String("timeout-settings", "", "Indicates the timeout settings of the pool.")
	workspaces_createWorkspacesPoolCmd.MarkFlagRequired("bundle-id")
	workspaces_createWorkspacesPoolCmd.MarkFlagRequired("capacity")
	workspaces_createWorkspacesPoolCmd.MarkFlagRequired("description")
	workspaces_createWorkspacesPoolCmd.MarkFlagRequired("directory-id")
	workspaces_createWorkspacesPoolCmd.MarkFlagRequired("pool-name")
	workspacesCmd.AddCommand(workspaces_createWorkspacesPoolCmd)
}
