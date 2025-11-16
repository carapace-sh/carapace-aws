package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_deleteVpceconfigurationCmd = &cobra.Command{
	Use:   "delete-vpceconfiguration",
	Short: "Deletes a configuration for your Amazon Virtual Private Cloud (VPC) endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_deleteVpceconfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_deleteVpceconfigurationCmd).Standalone()

		devicefarm_deleteVpceconfigurationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the VPC endpoint configuration you want to delete.")
		devicefarm_deleteVpceconfigurationCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_deleteVpceconfigurationCmd)
}
