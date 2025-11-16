package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_getHypervisorCmd = &cobra.Command{
	Use:   "get-hypervisor",
	Short: "This action requests information about the specified hypervisor to which the gateway will connect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_getHypervisorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_getHypervisorCmd).Standalone()

		backupGateway_getHypervisorCmd.Flags().String("hypervisor-arn", "", "The Amazon Resource Name (ARN) of the hypervisor.")
		backupGateway_getHypervisorCmd.MarkFlagRequired("hypervisor-arn")
	})
	backupGatewayCmd.AddCommand(backupGateway_getHypervisorCmd)
}
