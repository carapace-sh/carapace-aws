package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteOpenIdconnectProviderCmd = &cobra.Command{
	Use:   "delete-open-idconnect-provider",
	Short: "Deletes an OpenID Connect identity provider (IdP) resource object in IAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteOpenIdconnectProviderCmd).Standalone()

	iam_deleteOpenIdconnectProviderCmd.Flags().String("open-idconnect-provider-arn", "", "The Amazon Resource Name (ARN) of the IAM OpenID Connect provider resource object to delete.")
	iam_deleteOpenIdconnectProviderCmd.MarkFlagRequired("open-idconnect-provider-arn")
	iamCmd.AddCommand(iam_deleteOpenIdconnectProviderCmd)
}
