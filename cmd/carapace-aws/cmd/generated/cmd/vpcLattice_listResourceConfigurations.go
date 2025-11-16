package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listResourceConfigurationsCmd = &cobra.Command{
	Use:   "list-resource-configurations",
	Short: "Lists the resource configurations owned by or shared with this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listResourceConfigurationsCmd).Standalone()

	vpcLattice_listResourceConfigurationsCmd.Flags().String("domain-verification-identifier", "", "The domain verification ID.")
	vpcLattice_listResourceConfigurationsCmd.Flags().String("max-results", "", "The maximum page size.")
	vpcLattice_listResourceConfigurationsCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
	vpcLattice_listResourceConfigurationsCmd.Flags().String("resource-configuration-group-identifier", "", "The ID of the resource configuration of type `Group`.")
	vpcLattice_listResourceConfigurationsCmd.Flags().String("resource-gateway-identifier", "", "The ID of the resource gateway for the resource configuration.")
	vpcLatticeCmd.AddCommand(vpcLattice_listResourceConfigurationsCmd)
}
