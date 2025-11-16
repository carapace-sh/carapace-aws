package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sts_assumeRoleWithWebIdentityCmd = &cobra.Command{
	Use:   "assume-role-with-web-identity",
	Short: "Returns a set of temporary security credentials for users who have been authenticated in a mobile or web application with a web identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sts_assumeRoleWithWebIdentityCmd).Standalone()

	sts_assumeRoleWithWebIdentityCmd.Flags().String("duration-seconds", "", "The duration, in seconds, of the role session.")
	sts_assumeRoleWithWebIdentityCmd.Flags().String("policy", "", "An IAM policy in JSON format that you want to use as an inline session policy.")
	sts_assumeRoleWithWebIdentityCmd.Flags().String("policy-arns", "", "The Amazon Resource Names (ARNs) of the IAM managed policies that you want to use as managed session policies.")
	sts_assumeRoleWithWebIdentityCmd.Flags().String("provider-id", "", "The fully qualified host component of the domain name of the OAuth 2.0 identity provider.")
	sts_assumeRoleWithWebIdentityCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the role that the caller is assuming.")
	sts_assumeRoleWithWebIdentityCmd.Flags().String("role-session-name", "", "An identifier for the assumed role session.")
	sts_assumeRoleWithWebIdentityCmd.Flags().String("web-identity-token", "", "The OAuth 2.0 access token or OpenID Connect ID token that is provided by the identity provider.")
	sts_assumeRoleWithWebIdentityCmd.MarkFlagRequired("role-arn")
	sts_assumeRoleWithWebIdentityCmd.MarkFlagRequired("role-session-name")
	sts_assumeRoleWithWebIdentityCmd.MarkFlagRequired("web-identity-token")
	stsCmd.AddCommand(sts_assumeRoleWithWebIdentityCmd)
}
