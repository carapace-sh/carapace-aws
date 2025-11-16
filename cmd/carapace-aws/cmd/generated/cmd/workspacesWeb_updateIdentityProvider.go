package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_updateIdentityProviderCmd = &cobra.Command{
	Use:   "update-identity-provider",
	Short: "Updates the identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_updateIdentityProviderCmd).Standalone()

	workspacesWeb_updateIdentityProviderCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_updateIdentityProviderCmd.Flags().String("identity-provider-arn", "", "The ARN of the identity provider.")
	workspacesWeb_updateIdentityProviderCmd.Flags().String("identity-provider-details", "", "The details of the identity provider.")
	workspacesWeb_updateIdentityProviderCmd.Flags().String("identity-provider-name", "", "The name of the identity provider.")
	workspacesWeb_updateIdentityProviderCmd.Flags().String("identity-provider-type", "", "The type of the identity provider.")
	workspacesWeb_updateIdentityProviderCmd.MarkFlagRequired("identity-provider-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_updateIdentityProviderCmd)
}
