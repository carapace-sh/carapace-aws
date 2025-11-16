package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_deleteIdentityProviderCmd = &cobra.Command{
	Use:   "delete-identity-provider",
	Short: "Deletes the identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_deleteIdentityProviderCmd).Standalone()

	workspacesWeb_deleteIdentityProviderCmd.Flags().String("identity-provider-arn", "", "The ARN of the identity provider.")
	workspacesWeb_deleteIdentityProviderCmd.MarkFlagRequired("identity-provider-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_deleteIdentityProviderCmd)
}
