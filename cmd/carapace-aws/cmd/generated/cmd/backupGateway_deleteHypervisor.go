package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_deleteHypervisorCmd = &cobra.Command{
	Use:   "delete-hypervisor",
	Short: "Deletes a hypervisor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_deleteHypervisorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_deleteHypervisorCmd).Standalone()

		backupGateway_deleteHypervisorCmd.Flags().String("hypervisor-arn", "", "The Amazon Resource Name (ARN) of the hypervisor to delete.")
		backupGateway_deleteHypervisorCmd.MarkFlagRequired("hypervisor-arn")
	})
	backupGatewayCmd.AddCommand(backupGateway_deleteHypervisorCmd)
}
