package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_tagOpenIdconnectProviderCmd = &cobra.Command{
	Use:   "tag-open-idconnect-provider",
	Short: "Adds one or more tags to an OpenID Connect (OIDC)-compatible identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_tagOpenIdconnectProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_tagOpenIdconnectProviderCmd).Standalone()

		iam_tagOpenIdconnectProviderCmd.Flags().String("open-idconnect-provider-arn", "", "The ARN of the OIDC identity provider in IAM to which you want to add tags.")
		iam_tagOpenIdconnectProviderCmd.Flags().String("tags", "", "The list of tags that you want to attach to the OIDC identity provider in IAM.")
		iam_tagOpenIdconnectProviderCmd.MarkFlagRequired("open-idconnect-provider-arn")
		iam_tagOpenIdconnectProviderCmd.MarkFlagRequired("tags")
	})
	iamCmd.AddCommand(iam_tagOpenIdconnectProviderCmd)
}
