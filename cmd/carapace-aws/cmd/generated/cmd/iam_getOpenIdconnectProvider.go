package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getOpenIdconnectProviderCmd = &cobra.Command{
	Use:   "get-open-idconnect-provider",
	Short: "Returns information about the specified OpenID Connect (OIDC) provider resource object in IAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getOpenIdconnectProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_getOpenIdconnectProviderCmd).Standalone()

		iam_getOpenIdconnectProviderCmd.Flags().String("open-idconnect-provider-arn", "", "The Amazon Resource Name (ARN) of the OIDC provider resource object in IAM to get information for.")
		iam_getOpenIdconnectProviderCmd.MarkFlagRequired("open-idconnect-provider-arn")
	})
	iamCmd.AddCommand(iam_getOpenIdconnectProviderCmd)
}
