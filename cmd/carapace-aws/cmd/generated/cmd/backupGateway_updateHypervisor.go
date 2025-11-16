package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_updateHypervisorCmd = &cobra.Command{
	Use:   "update-hypervisor",
	Short: "Updates a hypervisor metadata, including its host, username, and password.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_updateHypervisorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_updateHypervisorCmd).Standalone()

		backupGateway_updateHypervisorCmd.Flags().String("host", "", "The updated host of the hypervisor.")
		backupGateway_updateHypervisorCmd.Flags().String("hypervisor-arn", "", "The Amazon Resource Name (ARN) of the hypervisor to update.")
		backupGateway_updateHypervisorCmd.Flags().String("log-group-arn", "", "The Amazon Resource Name (ARN) of the group of gateways within the requested log.")
		backupGateway_updateHypervisorCmd.Flags().String("name", "", "The updated name for the hypervisor")
		backupGateway_updateHypervisorCmd.Flags().String("password", "", "The updated password for the hypervisor.")
		backupGateway_updateHypervisorCmd.Flags().String("username", "", "The updated username for the hypervisor.")
		backupGateway_updateHypervisorCmd.MarkFlagRequired("hypervisor-arn")
	})
	backupGatewayCmd.AddCommand(backupGateway_updateHypervisorCmd)
}
