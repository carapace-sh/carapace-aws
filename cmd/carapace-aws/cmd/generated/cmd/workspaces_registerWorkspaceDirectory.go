package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_registerWorkspaceDirectoryCmd = &cobra.Command{
	Use:   "register-workspace-directory",
	Short: "Registers the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_registerWorkspaceDirectoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_registerWorkspaceDirectoryCmd).Standalone()

		workspaces_registerWorkspaceDirectoryCmd.Flags().String("active-directory-config", "", "The active directory config of the directory.")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("directory-id", "", "The identifier of the directory.")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("enable-self-service", "", "Indicates whether self-service capabilities are enabled or disabled.")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("idc-instance-arn", "", "The Amazon Resource Name (ARN) of the identity center instance.")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("microsoft-entra-config", "", "The details about Microsoft Entra config.")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("subnet-ids", "", "The identifiers of the subnets for your virtual private cloud (VPC).")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("tags", "", "The tags associated with the directory.")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("tenancy", "", "Indicates whether your WorkSpace directory is dedicated or shared.")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("user-identity-type", "", "The type of identity management the user is using.")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("workspace-directory-description", "", "Description of the directory to register.")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("workspace-directory-name", "", "The name of the directory to register.")
		workspaces_registerWorkspaceDirectoryCmd.Flags().String("workspace-type", "", "Indicates whether the directory's WorkSpace type is personal or pools.")
	})
	workspacesCmd.AddCommand(workspaces_registerWorkspaceDirectoryCmd)
}
