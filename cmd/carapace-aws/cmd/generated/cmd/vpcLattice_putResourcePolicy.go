package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Attaches a resource-based permission policy to a service or service network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_putResourcePolicyCmd).Standalone()

	vpcLattice_putResourcePolicyCmd.Flags().String("policy", "", "An IAM policy.")
	vpcLattice_putResourcePolicyCmd.Flags().String("resource-arn", "", "The ID or ARN of the service network or service for which the policy is created.")
	vpcLattice_putResourcePolicyCmd.MarkFlagRequired("policy")
	vpcLattice_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	vpcLatticeCmd.AddCommand(vpcLattice_putResourcePolicyCmd)
}
