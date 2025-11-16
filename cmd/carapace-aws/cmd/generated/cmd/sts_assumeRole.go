package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sts_assumeRoleCmd = &cobra.Command{
	Use:   "assume-role",
	Short: "Returns a set of temporary security credentials that you can use to access Amazon Web Services resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sts_assumeRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sts_assumeRoleCmd).Standalone()

		sts_assumeRoleCmd.Flags().String("duration-seconds", "", "The duration, in seconds, of the role session.")
		sts_assumeRoleCmd.Flags().String("external-id", "", "A unique identifier that might be required when you assume a role in another account.")
		sts_assumeRoleCmd.Flags().String("policy", "", "An IAM policy in JSON format that you want to use as an inline session policy.")
		sts_assumeRoleCmd.Flags().String("policy-arns", "", "The Amazon Resource Names (ARNs) of the IAM managed policies that you want to use as managed session policies.")
		sts_assumeRoleCmd.Flags().String("provided-contexts", "", "A list of previously acquired trusted context assertions in the format of a JSON array.")
		sts_assumeRoleCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the role to assume.")
		sts_assumeRoleCmd.Flags().String("role-session-name", "", "An identifier for the assumed role session.")
		sts_assumeRoleCmd.Flags().String("serial-number", "", "The identification number of the MFA device that is associated with the user who is making the `AssumeRole` call.")
		sts_assumeRoleCmd.Flags().String("source-identity", "", "The source identity specified by the principal that is calling the `AssumeRole` operation.")
		sts_assumeRoleCmd.Flags().String("tags", "", "A list of session tags that you want to pass.")
		sts_assumeRoleCmd.Flags().String("token-code", "", "The value provided by the MFA device, if the trust policy of the role being assumed requires MFA.")
		sts_assumeRoleCmd.Flags().String("transitive-tag-keys", "", "A list of keys for session tags that you want to set as transitive.")
		sts_assumeRoleCmd.MarkFlagRequired("role-arn")
		sts_assumeRoleCmd.MarkFlagRequired("role-session-name")
	})
	stsCmd.AddCommand(sts_assumeRoleCmd)
}
