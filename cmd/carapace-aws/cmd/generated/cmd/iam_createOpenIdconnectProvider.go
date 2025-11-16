package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createOpenIdconnectProviderCmd = &cobra.Command{
	Use:   "create-open-idconnect-provider",
	Short: "Creates an IAM entity to describe an identity provider (IdP) that supports [OpenID Connect (OIDC)](http://openid.net/connect/).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createOpenIdconnectProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_createOpenIdconnectProviderCmd).Standalone()

		iam_createOpenIdconnectProviderCmd.Flags().String("client-idlist", "", "Provides a list of client IDs, also known as audiences.")
		iam_createOpenIdconnectProviderCmd.Flags().String("tags", "", "A list of tags that you want to attach to the new IAM OpenID Connect (OIDC) provider.")
		iam_createOpenIdconnectProviderCmd.Flags().String("thumbprint-list", "", "A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificates.")
		iam_createOpenIdconnectProviderCmd.Flags().String("url", "", "The URL of the identity provider.")
		iam_createOpenIdconnectProviderCmd.MarkFlagRequired("url")
	})
	iamCmd.AddCommand(iam_createOpenIdconnectProviderCmd)
}
