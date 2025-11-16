package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getDefaultEncryptionConfigurationCmd = &cobra.Command{
	Use:   "get-default-encryption-configuration",
	Short: "Retrieves information about the default encryption configuration for the Amazon Web Services account in the default or specified region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getDefaultEncryptionConfigurationCmd).Standalone()

	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getDefaultEncryptionConfigurationCmd)
}
