package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_updateResourceConfigurationCmd = &cobra.Command{
	Use:   "update-resource-configuration",
	Short: "Updates the specified resource configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_updateResourceConfigurationCmd).Standalone()

	vpcLattice_updateResourceConfigurationCmd.Flags().Bool("allow-association-to-shareable-service-network", false, "Indicates whether to add the resource configuration to service networks that are shared with other accounts.")
	vpcLattice_updateResourceConfigurationCmd.Flags().Bool("no-allow-association-to-shareable-service-network", false, "Indicates whether to add the resource configuration to service networks that are shared with other accounts.")
	vpcLattice_updateResourceConfigurationCmd.Flags().String("port-ranges", "", "The TCP port ranges that a consumer can use to access a resource configuration.")
	vpcLattice_updateResourceConfigurationCmd.Flags().String("resource-configuration-definition", "", "Identifies the resource configuration in one of the following ways:")
	vpcLattice_updateResourceConfigurationCmd.Flags().String("resource-configuration-identifier", "", "The ID of the resource configuration.")
	vpcLattice_updateResourceConfigurationCmd.Flag("no-allow-association-to-shareable-service-network").Hidden = true
	vpcLattice_updateResourceConfigurationCmd.MarkFlagRequired("resource-configuration-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_updateResourceConfigurationCmd)
}
