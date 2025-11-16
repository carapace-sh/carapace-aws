package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteAuthPolicyCmd = &cobra.Command{
	Use:   "delete-auth-policy",
	Short: "Deletes the specified auth policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteAuthPolicyCmd).Standalone()

	vpcLattice_deleteAuthPolicyCmd.Flags().String("resource-identifier", "", "The ID or ARN of the resource.")
	vpcLattice_deleteAuthPolicyCmd.MarkFlagRequired("resource-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_deleteAuthPolicyCmd)
}
