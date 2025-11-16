package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_addClientIdtoOpenIdconnectProviderCmd = &cobra.Command{
	Use:   "add-client-idto-open-idconnect-provider",
	Short: "Adds a new client ID (also known as audience) to the list of client IDs already registered for the specified IAM OpenID Connect (OIDC) provider resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_addClientIdtoOpenIdconnectProviderCmd).Standalone()

	iam_addClientIdtoOpenIdconnectProviderCmd.Flags().String("client-id", "", "The client ID (also known as audience) to add to the IAM OpenID Connect provider resource.")
	iam_addClientIdtoOpenIdconnectProviderCmd.Flags().String("open-idconnect-provider-arn", "", "The Amazon Resource Name (ARN) of the IAM OpenID Connect (OIDC) provider resource to add the client ID to.")
	iam_addClientIdtoOpenIdconnectProviderCmd.MarkFlagRequired("client-id")
	iam_addClientIdtoOpenIdconnectProviderCmd.MarkFlagRequired("open-idconnect-provider-arn")
	iamCmd.AddCommand(iam_addClientIdtoOpenIdconnectProviderCmd)
}
