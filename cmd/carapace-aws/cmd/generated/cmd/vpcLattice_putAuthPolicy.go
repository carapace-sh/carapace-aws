package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_putAuthPolicyCmd = &cobra.Command{
	Use:   "put-auth-policy",
	Short: "Creates or updates the auth policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_putAuthPolicyCmd).Standalone()

	vpcLattice_putAuthPolicyCmd.Flags().String("policy", "", "The auth policy.")
	vpcLattice_putAuthPolicyCmd.Flags().String("resource-identifier", "", "The ID or ARN of the service network or service for which the policy is created.")
	vpcLattice_putAuthPolicyCmd.MarkFlagRequired("policy")
	vpcLattice_putAuthPolicyCmd.MarkFlagRequired("resource-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_putAuthPolicyCmd)
}
