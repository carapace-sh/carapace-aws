package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_listVirtualMachinesCmd = &cobra.Command{
	Use:   "list-virtual-machines",
	Short: "Lists your virtual machines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_listVirtualMachinesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_listVirtualMachinesCmd).Standalone()

		backupGateway_listVirtualMachinesCmd.Flags().String("hypervisor-arn", "", "The Amazon Resource Name (ARN) of the hypervisor connected to your virtual machine.")
		backupGateway_listVirtualMachinesCmd.Flags().String("max-results", "", "The maximum number of virtual machines to list.")
		backupGateway_listVirtualMachinesCmd.Flags().String("next-token", "", "The next item following a partial list of returned resources.")
	})
	backupGatewayCmd.AddCommand(backupGateway_listVirtualMachinesCmd)
}
