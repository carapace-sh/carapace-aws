package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getResourceConfigurationCmd = &cobra.Command{
	Use:   "get-resource-configuration",
	Short: "Retrieves information about the specified resource configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getResourceConfigurationCmd).Standalone()

	vpcLattice_getResourceConfigurationCmd.Flags().String("resource-configuration-identifier", "", "The ID of the resource configuration.")
	vpcLattice_getResourceConfigurationCmd.MarkFlagRequired("resource-configuration-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_getResourceConfigurationCmd)
}
