package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getVpceconfigurationCmd = &cobra.Command{
	Use:   "get-vpceconfiguration",
	Short: "Returns information about the configuration settings for your Amazon Virtual Private Cloud (VPC) endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getVpceconfigurationCmd).Standalone()

	devicefarm_getVpceconfigurationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the VPC endpoint configuration you want to describe.")
	devicefarm_getVpceconfigurationCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_getVpceconfigurationCmd)
}
