package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_generateDataKeyWithoutPlaintextCmd = &cobra.Command{
	Use:   "generate-data-key-without-plaintext",
	Short: "Returns a unique symmetric data key for use outside of KMS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_generateDataKeyWithoutPlaintextCmd).Standalone()

	kms_generateDataKeyWithoutPlaintextCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
	kms_generateDataKeyWithoutPlaintextCmd.Flags().String("encryption-context", "", "Specifies the encryption context that will be used when encrypting the data key.")
	kms_generateDataKeyWithoutPlaintextCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
	kms_generateDataKeyWithoutPlaintextCmd.Flags().String("key-id", "", "Specifies the symmetric encryption KMS key that encrypts the data key.")
	kms_generateDataKeyWithoutPlaintextCmd.Flags().String("key-spec", "", "The length of the data key.")
	kms_generateDataKeyWithoutPlaintextCmd.Flags().String("number-of-bytes", "", "The length of the data key in bytes.")
	kms_generateDataKeyWithoutPlaintextCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_generateDataKeyWithoutPlaintextCmd)
}
