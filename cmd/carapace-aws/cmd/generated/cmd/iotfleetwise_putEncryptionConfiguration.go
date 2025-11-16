package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_putEncryptionConfigurationCmd = &cobra.Command{
	Use:   "put-encryption-configuration",
	Short: "Creates or updates the encryption configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_putEncryptionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_putEncryptionConfigurationCmd).Standalone()

		iotfleetwise_putEncryptionConfigurationCmd.Flags().String("encryption-type", "", "The type of encryption.")
		iotfleetwise_putEncryptionConfigurationCmd.Flags().String("kms-key-id", "", "The ID of the KMS key that is used for encryption.")
		iotfleetwise_putEncryptionConfigurationCmd.MarkFlagRequired("encryption-type")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_putEncryptionConfigurationCmd)
}
