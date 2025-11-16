package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createServiceNetworkVpcAssociationCmd = &cobra.Command{
	Use:   "create-service-network-vpc-association",
	Short: "Associates a VPC with a service network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createServiceNetworkVpcAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_createServiceNetworkVpcAssociationCmd).Standalone()

		vpcLattice_createServiceNetworkVpcAssociationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		vpcLattice_createServiceNetworkVpcAssociationCmd.Flags().String("dns-options", "", "DNS options for the service network VPC association.")
		vpcLattice_createServiceNetworkVpcAssociationCmd.Flags().Bool("no-private-dns-enabled", false, "Indicates if private DNS is enabled for the VPC association.")
		vpcLattice_createServiceNetworkVpcAssociationCmd.Flags().Bool("private-dns-enabled", false, "Indicates if private DNS is enabled for the VPC association.")
		vpcLattice_createServiceNetworkVpcAssociationCmd.Flags().String("security-group-ids", "", "The IDs of the security groups.")
		vpcLattice_createServiceNetworkVpcAssociationCmd.Flags().String("service-network-identifier", "", "The ID or ARN of the service network.")
		vpcLattice_createServiceNetworkVpcAssociationCmd.Flags().String("tags", "", "The tags for the association.")
		vpcLattice_createServiceNetworkVpcAssociationCmd.Flags().String("vpc-identifier", "", "The ID of the VPC.")
		vpcLattice_createServiceNetworkVpcAssociationCmd.Flag("no-private-dns-enabled").Hidden = true
		vpcLattice_createServiceNetworkVpcAssociationCmd.MarkFlagRequired("service-network-identifier")
		vpcLattice_createServiceNetworkVpcAssociationCmd.MarkFlagRequired("vpc-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_createServiceNetworkVpcAssociationCmd)
}
