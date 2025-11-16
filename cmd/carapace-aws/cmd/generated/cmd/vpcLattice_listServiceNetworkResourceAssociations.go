package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listServiceNetworkResourceAssociationsCmd = &cobra.Command{
	Use:   "list-service-network-resource-associations",
	Short: "Lists the associations between a service network and a resource configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listServiceNetworkResourceAssociationsCmd).Standalone()

	vpcLattice_listServiceNetworkResourceAssociationsCmd.Flags().Bool("include-children", false, "Include service network resource associations of the child resource configuration with the grouped resource configuration.")
	vpcLattice_listServiceNetworkResourceAssociationsCmd.Flags().String("max-results", "", "The maximum page size.")
	vpcLattice_listServiceNetworkResourceAssociationsCmd.Flags().String("next-token", "", "If there are additional results, a pagination token for the next page of results.")
	vpcLattice_listServiceNetworkResourceAssociationsCmd.Flags().Bool("no-include-children", false, "Include service network resource associations of the child resource configuration with the grouped resource configuration.")
	vpcLattice_listServiceNetworkResourceAssociationsCmd.Flags().String("resource-configuration-identifier", "", "The ID of the resource configuration.")
	vpcLattice_listServiceNetworkResourceAssociationsCmd.Flags().String("service-network-identifier", "", "The ID of the service network.")
	vpcLattice_listServiceNetworkResourceAssociationsCmd.Flag("no-include-children").Hidden = true
	vpcLatticeCmd.AddCommand(vpcLattice_listServiceNetworkResourceAssociationsCmd)
}
