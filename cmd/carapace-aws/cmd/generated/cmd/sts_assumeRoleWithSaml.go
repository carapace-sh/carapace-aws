package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sts_assumeRoleWithSamlCmd = &cobra.Command{
	Use:   "assume-role-with-saml",
	Short: "Returns a set of temporary security credentials for users who have been authenticated via a SAML authentication response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sts_assumeRoleWithSamlCmd).Standalone()

	sts_assumeRoleWithSamlCmd.Flags().String("duration-seconds", "", "The duration, in seconds, of the role session.")
	sts_assumeRoleWithSamlCmd.Flags().String("policy", "", "An IAM policy in JSON format that you want to use as an inline session policy.")
	sts_assumeRoleWithSamlCmd.Flags().String("policy-arns", "", "The Amazon Resource Names (ARNs) of the IAM managed policies that you want to use as managed session policies.")
	sts_assumeRoleWithSamlCmd.Flags().String("principal-arn", "", "The Amazon Resource Name (ARN) of the SAML provider in IAM that describes the IdP.")
	sts_assumeRoleWithSamlCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the role that the caller is assuming.")
	sts_assumeRoleWithSamlCmd.Flags().String("samlassertion", "", "The base64 encoded SAML authentication response provided by the IdP.")
	sts_assumeRoleWithSamlCmd.MarkFlagRequired("principal-arn")
	sts_assumeRoleWithSamlCmd.MarkFlagRequired("role-arn")
	sts_assumeRoleWithSamlCmd.MarkFlagRequired("samlassertion")
	stsCmd.AddCommand(sts_assumeRoleWithSamlCmd)
}
