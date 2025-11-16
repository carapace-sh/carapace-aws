package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_testHypervisorConfigurationCmd = &cobra.Command{
	Use:   "test-hypervisor-configuration",
	Short: "Tests your hypervisor configuration to validate that backup gateway can connect with the hypervisor and its resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_testHypervisorConfigurationCmd).Standalone()

	backupGateway_testHypervisorConfigurationCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway to the hypervisor to test.")
	backupGateway_testHypervisorConfigurationCmd.Flags().String("host", "", "The server host of the hypervisor.")
	backupGateway_testHypervisorConfigurationCmd.Flags().String("password", "", "The password for the hypervisor.")
	backupGateway_testHypervisorConfigurationCmd.Flags().String("username", "", "The username for the hypervisor.")
	backupGateway_testHypervisorConfigurationCmd.MarkFlagRequired("gateway-arn")
	backupGateway_testHypervisorConfigurationCmd.MarkFlagRequired("host")
	backupGatewayCmd.AddCommand(backupGateway_testHypervisorConfigurationCmd)
}
