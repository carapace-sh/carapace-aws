package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_deleteClientBrandingCmd = &cobra.Command{
	Use:   "delete-client-branding",
	Short: "Deletes customized client branding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_deleteClientBrandingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_deleteClientBrandingCmd).Standalone()

		workspaces_deleteClientBrandingCmd.Flags().String("platforms", "", "The device type for which you want to delete client branding.")
		workspaces_deleteClientBrandingCmd.Flags().String("resource-id", "", "The directory identifier of the WorkSpace for which you want to delete client branding.")
		workspaces_deleteClientBrandingCmd.MarkFlagRequired("platforms")
		workspaces_deleteClientBrandingCmd.MarkFlagRequired("resource-id")
	})
	workspacesCmd.AddCommand(workspaces_deleteClientBrandingCmd)
}
