package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createResourceConfigurationCmd = &cobra.Command{
	Use:   "create-resource-configuration",
	Short: "Creates a resource configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createResourceConfigurationCmd).Standalone()

	vpcLattice_createResourceConfigurationCmd.Flags().Bool("allow-association-to-shareable-service-network", false, "(SINGLE, GROUP, ARN) Specifies whether the resource configuration can be associated with a sharable service network.")
	vpcLattice_createResourceConfigurationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	vpcLattice_createResourceConfigurationCmd.Flags().String("custom-domain-name", "", "A custom domain name for your resource configuration.")
	vpcLattice_createResourceConfigurationCmd.Flags().String("domain-verification-identifier", "", "The domain verification ID of your verified custom domain name.")
	vpcLattice_createResourceConfigurationCmd.Flags().String("group-domain", "", "(GROUP) The group domain for a group resource configuration.")
	vpcLattice_createResourceConfigurationCmd.Flags().String("name", "", "The name of the resource configuration.")
	vpcLattice_createResourceConfigurationCmd.Flags().Bool("no-allow-association-to-shareable-service-network", false, "(SINGLE, GROUP, ARN) Specifies whether the resource configuration can be associated with a sharable service network.")
	vpcLattice_createResourceConfigurationCmd.Flags().String("port-ranges", "", "(SINGLE, GROUP, CHILD) The TCP port ranges that a consumer can use to access a resource configuration (for example: 1-65535).")
	vpcLattice_createResourceConfigurationCmd.Flags().String("protocol", "", "(SINGLE, GROUP) The protocol accepted by the resource configuration.")
	vpcLattice_createResourceConfigurationCmd.Flags().String("resource-configuration-definition", "", "Identifies the resource configuration in one of the following ways:")
	vpcLattice_createResourceConfigurationCmd.Flags().String("resource-configuration-group-identifier", "", "(CHILD) The ID or ARN of the parent resource configuration of type `GROUP`.")
	vpcLattice_createResourceConfigurationCmd.Flags().String("resource-gateway-identifier", "", "(SINGLE, GROUP, ARN) The ID or ARN of the resource gateway used to connect to the resource configuration.")
	vpcLattice_createResourceConfigurationCmd.Flags().String("tags", "", "The tags for the resource configuration.")
	vpcLattice_createResourceConfigurationCmd.Flags().String("type", "", "The type of resource configuration.")
	vpcLattice_createResourceConfigurationCmd.MarkFlagRequired("name")
	vpcLattice_createResourceConfigurationCmd.Flag("no-allow-association-to-shareable-service-network").Hidden = true
	vpcLattice_createResourceConfigurationCmd.MarkFlagRequired("type")
	vpcLatticeCmd.AddCommand(vpcLattice_createResourceConfigurationCmd)
}
