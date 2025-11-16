package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_startVirtualMachinesMetadataSyncCmd = &cobra.Command{
	Use:   "start-virtual-machines-metadata-sync",
	Short: "This action sends a request to sync metadata across the specified virtual machines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_startVirtualMachinesMetadataSyncCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_startVirtualMachinesMetadataSyncCmd).Standalone()

		backupGateway_startVirtualMachinesMetadataSyncCmd.Flags().String("hypervisor-arn", "", "The Amazon Resource Name (ARN) of the hypervisor.")
		backupGateway_startVirtualMachinesMetadataSyncCmd.MarkFlagRequired("hypervisor-arn")
	})
	backupGatewayCmd.AddCommand(backupGateway_startVirtualMachinesMetadataSyncCmd)
}
