package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getIdentityProviderCmd = &cobra.Command{
	Use:   "get-identity-provider",
	Short: "Gets the identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getIdentityProviderCmd).Standalone()

	workspacesWeb_getIdentityProviderCmd.Flags().String("identity-provider-arn", "", "The ARN of the identity provider.")
	workspacesWeb_getIdentityProviderCmd.MarkFlagRequired("identity-provider-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_getIdentityProviderCmd)
}
