package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_untagOpenIdconnectProviderCmd = &cobra.Command{
	Use:   "untag-open-idconnect-provider",
	Short: "Removes the specified tags from the specified OpenID Connect (OIDC)-compatible identity provider in IAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_untagOpenIdconnectProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_untagOpenIdconnectProviderCmd).Standalone()

		iam_untagOpenIdconnectProviderCmd.Flags().String("open-idconnect-provider-arn", "", "The ARN of the OIDC provider in IAM from which you want to remove tags.")
		iam_untagOpenIdconnectProviderCmd.Flags().String("tag-keys", "", "A list of key names as a simple array of strings.")
		iam_untagOpenIdconnectProviderCmd.MarkFlagRequired("open-idconnect-provider-arn")
		iam_untagOpenIdconnectProviderCmd.MarkFlagRequired("tag-keys")
	})
	iamCmd.AddCommand(iam_untagOpenIdconnectProviderCmd)
}
