package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_associateAccessPolicyCmd = &cobra.Command{
	Use:   "associate-access-policy",
	Short: "Associates an access policy and its scope to an access entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_associateAccessPolicyCmd).Standalone()

	eks_associateAccessPolicyCmd.Flags().String("access-scope", "", "The scope for the `AccessPolicy`.")
	eks_associateAccessPolicyCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_associateAccessPolicyCmd.Flags().String("policy-arn", "", "The ARN of the `AccessPolicy` that you're associating.")
	eks_associateAccessPolicyCmd.Flags().String("principal-arn", "", "The Amazon Resource Name (ARN) of the IAM user or role for the `AccessEntry` that you're associating the access policy to.")
	eks_associateAccessPolicyCmd.MarkFlagRequired("access-scope")
	eks_associateAccessPolicyCmd.MarkFlagRequired("cluster-name")
	eks_associateAccessPolicyCmd.MarkFlagRequired("policy-arn")
	eks_associateAccessPolicyCmd.MarkFlagRequired("principal-arn")
	eksCmd.AddCommand(eks_associateAccessPolicyCmd)
}
