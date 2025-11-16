package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getKmsencryptionKeyCmd = &cobra.Command{
	Use:   "get-kmsencryption-key",
	Short: "Gets the encryption key if a KMS key has been specified to be used to encrypt content in Amazon Fraud Detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getKmsencryptionKeyCmd).Standalone()

	frauddetectorCmd.AddCommand(frauddetector_getKmsencryptionKeyCmd)
}
