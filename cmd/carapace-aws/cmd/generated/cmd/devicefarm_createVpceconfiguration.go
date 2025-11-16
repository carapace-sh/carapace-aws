package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_createVpceconfigurationCmd = &cobra.Command{
	Use:   "create-vpceconfiguration",
	Short: "Creates a configuration record in Device Farm for your Amazon Virtual Private Cloud (VPC) endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_createVpceconfigurationCmd).Standalone()

	devicefarm_createVpceconfigurationCmd.Flags().String("service-dns-name", "", "The DNS name of the service running in your VPC that you want Device Farm to test.")
	devicefarm_createVpceconfigurationCmd.Flags().String("vpce-configuration-description", "", "An optional description that provides details about your VPC endpoint configuration.")
	devicefarm_createVpceconfigurationCmd.Flags().String("vpce-configuration-name", "", "The friendly name you give to your VPC endpoint configuration, to manage your configurations more easily.")
	devicefarm_createVpceconfigurationCmd.Flags().String("vpce-service-name", "", "The name of the VPC endpoint service running in your AWS account that you want Device Farm to test.")
	devicefarm_createVpceconfigurationCmd.MarkFlagRequired("service-dns-name")
	devicefarm_createVpceconfigurationCmd.MarkFlagRequired("vpce-configuration-name")
	devicefarm_createVpceconfigurationCmd.MarkFlagRequired("vpce-service-name")
	devicefarmCmd.AddCommand(devicefarm_createVpceconfigurationCmd)
}
