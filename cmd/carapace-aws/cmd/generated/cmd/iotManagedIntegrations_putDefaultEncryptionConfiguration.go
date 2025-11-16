package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_putDefaultEncryptionConfigurationCmd = &cobra.Command{
	Use:   "put-default-encryption-configuration",
	Short: "Sets the default encryption configuration for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_putDefaultEncryptionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_putDefaultEncryptionConfigurationCmd).Standalone()

		iotManagedIntegrations_putDefaultEncryptionConfigurationCmd.Flags().String("encryption-type", "", "The type of encryption used for the encryption configuration.")
		iotManagedIntegrations_putDefaultEncryptionConfigurationCmd.Flags().String("kms-key-arn", "", "The Key Amazon Resource Name (ARN) of the AWS KMS key used for KMS encryption if you use `KMS_BASED_ENCRYPTION`.")
		iotManagedIntegrations_putDefaultEncryptionConfigurationCmd.MarkFlagRequired("encryption-type")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_putDefaultEncryptionConfigurationCmd)
}
