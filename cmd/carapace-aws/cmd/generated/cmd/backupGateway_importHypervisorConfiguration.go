package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_importHypervisorConfigurationCmd = &cobra.Command{
	Use:   "import-hypervisor-configuration",
	Short: "Connect to a hypervisor by importing its configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_importHypervisorConfigurationCmd).Standalone()

	backupGateway_importHypervisorConfigurationCmd.Flags().String("host", "", "The server host of the hypervisor.")
	backupGateway_importHypervisorConfigurationCmd.Flags().String("kms-key-arn", "", "The Key Management Service for the hypervisor.")
	backupGateway_importHypervisorConfigurationCmd.Flags().String("name", "", "The name of the hypervisor.")
	backupGateway_importHypervisorConfigurationCmd.Flags().String("password", "", "The password for the hypervisor.")
	backupGateway_importHypervisorConfigurationCmd.Flags().String("tags", "", "The tags of the hypervisor configuration to import.")
	backupGateway_importHypervisorConfigurationCmd.Flags().String("username", "", "The username for the hypervisor.")
	backupGateway_importHypervisorConfigurationCmd.MarkFlagRequired("host")
	backupGateway_importHypervisorConfigurationCmd.MarkFlagRequired("name")
	backupGatewayCmd.AddCommand(backupGateway_importHypervisorConfigurationCmd)
}
