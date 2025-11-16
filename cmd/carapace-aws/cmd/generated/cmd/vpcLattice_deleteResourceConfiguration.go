package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteResourceConfigurationCmd = &cobra.Command{
	Use:   "delete-resource-configuration",
	Short: "Deletes the specified resource configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteResourceConfigurationCmd).Standalone()

	vpcLattice_deleteResourceConfigurationCmd.Flags().String("resource-configuration-identifier", "", "The ID or ARN of the resource configuration.")
	vpcLattice_deleteResourceConfigurationCmd.MarkFlagRequired("resource-configuration-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_deleteResourceConfigurationCmd)
}
