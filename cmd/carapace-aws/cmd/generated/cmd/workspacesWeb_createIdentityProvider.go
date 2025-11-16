package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_createIdentityProviderCmd = &cobra.Command{
	Use:   "create-identity-provider",
	Short: "Creates an identity provider resource that is then associated with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_createIdentityProviderCmd).Standalone()

	workspacesWeb_createIdentityProviderCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_createIdentityProviderCmd.Flags().String("identity-provider-details", "", "The identity provider details.")
	workspacesWeb_createIdentityProviderCmd.Flags().String("identity-provider-name", "", "The identity provider name.")
	workspacesWeb_createIdentityProviderCmd.Flags().String("identity-provider-type", "", "The identity provider type.")
	workspacesWeb_createIdentityProviderCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
	workspacesWeb_createIdentityProviderCmd.Flags().String("tags", "", "The tags to add to the identity provider resource.")
	workspacesWeb_createIdentityProviderCmd.MarkFlagRequired("identity-provider-details")
	workspacesWeb_createIdentityProviderCmd.MarkFlagRequired("identity-provider-name")
	workspacesWeb_createIdentityProviderCmd.MarkFlagRequired("identity-provider-type")
	workspacesWeb_createIdentityProviderCmd.MarkFlagRequired("portal-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_createIdentityProviderCmd)
}
