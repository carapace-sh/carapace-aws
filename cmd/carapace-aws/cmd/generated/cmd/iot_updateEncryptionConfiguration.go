package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateEncryptionConfigurationCmd = &cobra.Command{
	Use:   "update-encryption-configuration",
	Short: "Updates the encryption configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateEncryptionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateEncryptionConfigurationCmd).Standalone()

		iot_updateEncryptionConfigurationCmd.Flags().String("encryption-type", "", "The type of the Amazon Web Services Key Management Service (KMS) key.")
		iot_updateEncryptionConfigurationCmd.Flags().String("kms-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role assumed by Amazon Web Services IoT Core to call KMS on behalf of the customer.")
		iot_updateEncryptionConfigurationCmd.Flags().String("kms-key-arn", "", "The ARN of the customer-managed KMS key.")
		iot_updateEncryptionConfigurationCmd.MarkFlagRequired("encryption-type")
	})
	iotCmd.AddCommand(iot_updateEncryptionConfigurationCmd)
}
