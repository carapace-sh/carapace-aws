package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifyCertificateBasedAuthPropertiesCmd = &cobra.Command{
	Use:   "modify-certificate-based-auth-properties",
	Short: "Modifies the properties of the certificate-based authentication you want to use with your WorkSpaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifyCertificateBasedAuthPropertiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_modifyCertificateBasedAuthPropertiesCmd).Standalone()

		workspaces_modifyCertificateBasedAuthPropertiesCmd.Flags().String("certificate-based-auth-properties", "", "The properties of the certificate-based authentication.")
		workspaces_modifyCertificateBasedAuthPropertiesCmd.Flags().String("properties-to-delete", "", "The properties of the certificate-based authentication you want to delete.")
		workspaces_modifyCertificateBasedAuthPropertiesCmd.Flags().String("resource-id", "", "The resource identifiers, in the form of directory IDs.")
		workspaces_modifyCertificateBasedAuthPropertiesCmd.MarkFlagRequired("resource-id")
	})
	workspacesCmd.AddCommand(workspaces_modifyCertificateBasedAuthPropertiesCmd)
}
