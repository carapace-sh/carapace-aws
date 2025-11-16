package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateOpenIdconnectProviderThumbprintCmd = &cobra.Command{
	Use:   "update-open-idconnect-provider-thumbprint",
	Short: "Replaces the existing list of server certificate thumbprints associated with an OpenID Connect (OIDC) provider resource object with a new list of thumbprints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateOpenIdconnectProviderThumbprintCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_updateOpenIdconnectProviderThumbprintCmd).Standalone()

		iam_updateOpenIdconnectProviderThumbprintCmd.Flags().String("open-idconnect-provider-arn", "", "The Amazon Resource Name (ARN) of the IAM OIDC provider resource object for which you want to update the thumbprint.")
		iam_updateOpenIdconnectProviderThumbprintCmd.Flags().String("thumbprint-list", "", "A list of certificate thumbprints that are associated with the specified IAM OpenID Connect provider.")
		iam_updateOpenIdconnectProviderThumbprintCmd.MarkFlagRequired("open-idconnect-provider-arn")
		iam_updateOpenIdconnectProviderThumbprintCmd.MarkFlagRequired("thumbprint-list")
	})
	iamCmd.AddCommand(iam_updateOpenIdconnectProviderThumbprintCmd)
}
