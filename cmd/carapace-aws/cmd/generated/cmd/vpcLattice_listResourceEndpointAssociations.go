package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listResourceEndpointAssociationsCmd = &cobra.Command{
	Use:   "list-resource-endpoint-associations",
	Short: "Lists the associations for the specified VPC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listResourceEndpointAssociationsCmd).Standalone()

	vpcLattice_listResourceEndpointAssociationsCmd.Flags().String("max-results", "", "The maximum page size.")
	vpcLattice_listResourceEndpointAssociationsCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
	vpcLattice_listResourceEndpointAssociationsCmd.Flags().String("resource-configuration-identifier", "", "The ID for the resource configuration associated with the VPC endpoint.")
	vpcLattice_listResourceEndpointAssociationsCmd.Flags().String("resource-endpoint-association-identifier", "", "The ID of the association.")
	vpcLattice_listResourceEndpointAssociationsCmd.Flags().String("vpc-endpoint-id", "", "The ID of the VPC endpoint in the association.")
	vpcLattice_listResourceEndpointAssociationsCmd.Flags().String("vpc-endpoint-owner", "", "The owner of the VPC endpoint in the association.")
	vpcLattice_listResourceEndpointAssociationsCmd.MarkFlagRequired("resource-configuration-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_listResourceEndpointAssociationsCmd)
}
