package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifySamlPropertiesCmd = &cobra.Command{
	Use:   "modify-saml-properties",
	Short: "Modifies multiple properties related to SAML 2.0 authentication, including the enablement status, user access URL, and relay state parameter name that are used for configuring federation with an SAML 2.0 identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifySamlPropertiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_modifySamlPropertiesCmd).Standalone()

		workspaces_modifySamlPropertiesCmd.Flags().String("properties-to-delete", "", "The SAML properties to delete as part of your request.")
		workspaces_modifySamlPropertiesCmd.Flags().String("resource-id", "", "The directory identifier for which you want to configure SAML properties.")
		workspaces_modifySamlPropertiesCmd.Flags().String("saml-properties", "", "The properties for configuring SAML 2.0 authentication.")
		workspaces_modifySamlPropertiesCmd.MarkFlagRequired("resource-id")
	})
	workspacesCmd.AddCommand(workspaces_modifySamlPropertiesCmd)
}
