package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_updateVpceconfigurationCmd = &cobra.Command{
	Use:   "update-vpceconfiguration",
	Short: "Updates information about an Amazon Virtual Private Cloud (VPC) endpoint configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_updateVpceconfigurationCmd).Standalone()

	devicefarm_updateVpceconfigurationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the VPC endpoint configuration you want to update.")
	devicefarm_updateVpceconfigurationCmd.Flags().String("service-dns-name", "", "The DNS (domain) name used to connect to your private service in your VPC.")
	devicefarm_updateVpceconfigurationCmd.Flags().String("vpce-configuration-description", "", "An optional description that provides details about your VPC endpoint configuration.")
	devicefarm_updateVpceconfigurationCmd.Flags().String("vpce-configuration-name", "", "The friendly name you give to your VPC endpoint configuration to manage your configurations more easily.")
	devicefarm_updateVpceconfigurationCmd.Flags().String("vpce-service-name", "", "The name of the VPC endpoint service running in your AWS account that you want Device Farm to test.")
	devicefarm_updateVpceconfigurationCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_updateVpceconfigurationCmd)
}
