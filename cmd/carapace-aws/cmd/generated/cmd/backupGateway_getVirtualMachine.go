package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_getVirtualMachineCmd = &cobra.Command{
	Use:   "get-virtual-machine",
	Short: "By providing the ARN (Amazon Resource Name), this API returns the virtual machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_getVirtualMachineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_getVirtualMachineCmd).Standalone()

		backupGateway_getVirtualMachineCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the virtual machine.")
		backupGateway_getVirtualMachineCmd.MarkFlagRequired("resource-arn")
	})
	backupGatewayCmd.AddCommand(backupGateway_getVirtualMachineCmd)
}
