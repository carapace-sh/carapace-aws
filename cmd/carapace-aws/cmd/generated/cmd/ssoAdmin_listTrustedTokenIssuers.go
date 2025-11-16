package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listTrustedTokenIssuersCmd = &cobra.Command{
	Use:   "list-trusted-token-issuers",
	Short: "Lists all the trusted token issuers configured in an instance of IAM Identity Center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listTrustedTokenIssuersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_listTrustedTokenIssuersCmd).Standalone()

		ssoAdmin_listTrustedTokenIssuersCmd.Flags().String("instance-arn", "", "Specifies the ARN of the instance of IAM Identity Center with the trusted token issuer configurations that you want to list.")
		ssoAdmin_listTrustedTokenIssuersCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
		ssoAdmin_listTrustedTokenIssuersCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		ssoAdmin_listTrustedTokenIssuersCmd.MarkFlagRequired("instance-arn")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_listTrustedTokenIssuersCmd)
}
