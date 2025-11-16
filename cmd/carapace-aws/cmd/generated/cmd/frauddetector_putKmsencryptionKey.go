package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_putKmsencryptionKeyCmd = &cobra.Command{
	Use:   "put-kmsencryption-key",
	Short: "Specifies the KMS key to be used to encrypt content in Amazon Fraud Detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_putKmsencryptionKeyCmd).Standalone()

	frauddetector_putKmsencryptionKeyCmd.Flags().String("kms-encryption-key-arn", "", "The KMS encryption key ARN.")
	frauddetector_putKmsencryptionKeyCmd.MarkFlagRequired("kms-encryption-key-arn")
	frauddetectorCmd.AddCommand(frauddetector_putKmsencryptionKeyCmd)
}
