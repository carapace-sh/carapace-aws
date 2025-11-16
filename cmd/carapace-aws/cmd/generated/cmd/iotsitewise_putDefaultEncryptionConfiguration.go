package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_putDefaultEncryptionConfigurationCmd = &cobra.Command{
	Use:   "put-default-encryption-configuration",
	Short: "Sets the default encryption configuration for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_putDefaultEncryptionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_putDefaultEncryptionConfigurationCmd).Standalone()

		iotsitewise_putDefaultEncryptionConfigurationCmd.Flags().String("encryption-type", "", "The type of encryption used for the encryption configuration.")
		iotsitewise_putDefaultEncryptionConfigurationCmd.Flags().String("kms-key-id", "", "The Key ID of the customer managed key used for KMS encryption.")
		iotsitewise_putDefaultEncryptionConfigurationCmd.MarkFlagRequired("encryption-type")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_putDefaultEncryptionConfigurationCmd)
}
