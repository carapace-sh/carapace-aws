package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Retrieves information about the specified resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_getResourcePolicyCmd).Standalone()

		vpcLattice_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the service network or service.")
		vpcLattice_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_getResourcePolicyCmd)
}
