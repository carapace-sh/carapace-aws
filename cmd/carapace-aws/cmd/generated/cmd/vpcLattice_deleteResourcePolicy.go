package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the specified resource policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_deleteResourcePolicyCmd).Standalone()

		vpcLattice_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		vpcLattice_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_deleteResourcePolicyCmd)
}
