package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listOpenIdconnectProviderTagsCmd = &cobra.Command{
	Use:   "list-open-idconnect-provider-tags",
	Short: "Lists the tags that are attached to the specified OpenID Connect (OIDC)-compatible identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listOpenIdconnectProviderTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listOpenIdconnectProviderTagsCmd).Standalone()

		iam_listOpenIdconnectProviderTagsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listOpenIdconnectProviderTagsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listOpenIdconnectProviderTagsCmd.Flags().String("open-idconnect-provider-arn", "", "The ARN of the OpenID Connect (OIDC) identity provider whose tags you want to see.")
		iam_listOpenIdconnectProviderTagsCmd.MarkFlagRequired("open-idconnect-provider-arn")
	})
	iamCmd.AddCommand(iam_listOpenIdconnectProviderTagsCmd)
}
