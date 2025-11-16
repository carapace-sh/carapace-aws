package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getAuthPolicyCmd = &cobra.Command{
	Use:   "get-auth-policy",
	Short: "Retrieves information about the auth policy for the specified service or service network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getAuthPolicyCmd).Standalone()

	vpcLattice_getAuthPolicyCmd.Flags().String("resource-identifier", "", "The ID or ARN of the service network or service.")
	vpcLattice_getAuthPolicyCmd.MarkFlagRequired("resource-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_getAuthPolicyCmd)
}
