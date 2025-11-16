package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createServiceNetworkResourceAssociationCmd = &cobra.Command{
	Use:   "create-service-network-resource-association",
	Short: "Associates the specified service network with the specified resource configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createServiceNetworkResourceAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_createServiceNetworkResourceAssociationCmd).Standalone()

		vpcLattice_createServiceNetworkResourceAssociationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		vpcLattice_createServiceNetworkResourceAssociationCmd.Flags().Bool("no-private-dns-enabled", false, "Indicates if private DNS is enabled for the service network resource association.")
		vpcLattice_createServiceNetworkResourceAssociationCmd.Flags().Bool("private-dns-enabled", false, "Indicates if private DNS is enabled for the service network resource association.")
		vpcLattice_createServiceNetworkResourceAssociationCmd.Flags().String("resource-configuration-identifier", "", "The ID of the resource configuration to associate with the service network.")
		vpcLattice_createServiceNetworkResourceAssociationCmd.Flags().String("service-network-identifier", "", "The ID of the service network to associate with the resource configuration.")
		vpcLattice_createServiceNetworkResourceAssociationCmd.Flags().String("tags", "", "A key-value pair to associate with a resource.")
		vpcLattice_createServiceNetworkResourceAssociationCmd.Flag("no-private-dns-enabled").Hidden = true
		vpcLattice_createServiceNetworkResourceAssociationCmd.MarkFlagRequired("resource-configuration-identifier")
		vpcLattice_createServiceNetworkResourceAssociationCmd.MarkFlagRequired("service-network-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_createServiceNetworkResourceAssociationCmd)
}
