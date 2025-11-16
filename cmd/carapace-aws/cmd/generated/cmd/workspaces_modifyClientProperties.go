package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifyClientPropertiesCmd = &cobra.Command{
	Use:   "modify-client-properties",
	Short: "Modifies the properties of the specified Amazon WorkSpaces clients.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifyClientPropertiesCmd).Standalone()

	workspaces_modifyClientPropertiesCmd.Flags().String("client-properties", "", "Information about the Amazon WorkSpaces client.")
	workspaces_modifyClientPropertiesCmd.Flags().String("resource-id", "", "The resource identifiers, in the form of directory IDs.")
	workspaces_modifyClientPropertiesCmd.MarkFlagRequired("client-properties")
	workspaces_modifyClientPropertiesCmd.MarkFlagRequired("resource-id")
	workspacesCmd.AddCommand(workspaces_modifyClientPropertiesCmd)
}
