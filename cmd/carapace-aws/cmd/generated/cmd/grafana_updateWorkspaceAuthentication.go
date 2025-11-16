package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var grafana_updateWorkspaceAuthenticationCmd = &cobra.Command{
	Use:   "update-workspace-authentication",
	Short: "Use this operation to define the identity provider (IdP) that this workspace authenticates users from, using SAML.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(grafana_updateWorkspaceAuthenticationCmd).Standalone()

	grafana_updateWorkspaceAuthenticationCmd.Flags().String("authentication-providers", "", "Specifies whether this workspace uses SAML 2.0, IAM Identity Center, or both to authenticate users for using the Grafana console within a workspace.")
	grafana_updateWorkspaceAuthenticationCmd.Flags().String("saml-configuration", "", "If the workspace uses SAML, use this structure to map SAML assertion attributes to workspace user information and define which groups in the assertion attribute are to have the `Admin` and `Editor` roles in the workspace.")
	grafana_updateWorkspaceAuthenticationCmd.Flags().String("workspace-id", "", "The ID of the workspace to update the authentication for.")
	grafana_updateWorkspaceAuthenticationCmd.MarkFlagRequired("authentication-providers")
	grafana_updateWorkspaceAuthenticationCmd.MarkFlagRequired("workspace-id")
	grafanaCmd.AddCommand(grafana_updateWorkspaceAuthenticationCmd)
}
