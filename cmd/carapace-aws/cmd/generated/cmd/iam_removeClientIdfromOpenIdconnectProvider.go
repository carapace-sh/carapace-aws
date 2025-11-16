package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_removeClientIdfromOpenIdconnectProviderCmd = &cobra.Command{
	Use:   "remove-client-idfrom-open-idconnect-provider",
	Short: "Removes the specified client ID (also known as audience) from the list of client IDs registered for the specified IAM OpenID Connect (OIDC) provider resource object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_removeClientIdfromOpenIdconnectProviderCmd).Standalone()

	iam_removeClientIdfromOpenIdconnectProviderCmd.Flags().String("client-id", "", "The client ID (also known as audience) to remove from the IAM OIDC provider resource.")
	iam_removeClientIdfromOpenIdconnectProviderCmd.Flags().String("open-idconnect-provider-arn", "", "The Amazon Resource Name (ARN) of the IAM OIDC provider resource to remove the client ID from.")
	iam_removeClientIdfromOpenIdconnectProviderCmd.MarkFlagRequired("client-id")
	iam_removeClientIdfromOpenIdconnectProviderCmd.MarkFlagRequired("open-idconnect-provider-arn")
	iamCmd.AddCommand(iam_removeClientIdfromOpenIdconnectProviderCmd)
}
