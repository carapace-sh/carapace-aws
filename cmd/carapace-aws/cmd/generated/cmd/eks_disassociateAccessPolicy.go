package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_disassociateAccessPolicyCmd = &cobra.Command{
	Use:   "disassociate-access-policy",
	Short: "Disassociates an access policy from an access entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_disassociateAccessPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_disassociateAccessPolicyCmd).Standalone()

		eks_disassociateAccessPolicyCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_disassociateAccessPolicyCmd.Flags().String("policy-arn", "", "The ARN of the policy to disassociate from the access entry.")
		eks_disassociateAccessPolicyCmd.Flags().String("principal-arn", "", "The ARN of the IAM principal for the `AccessEntry`.")
		eks_disassociateAccessPolicyCmd.MarkFlagRequired("cluster-name")
		eks_disassociateAccessPolicyCmd.MarkFlagRequired("policy-arn")
		eks_disassociateAccessPolicyCmd.MarkFlagRequired("principal-arn")
	})
	eksCmd.AddCommand(eks_disassociateAccessPolicyCmd)
}
