package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_generateDataKeyPairWithoutPlaintextCmd = &cobra.Command{
	Use:   "generate-data-key-pair-without-plaintext",
	Short: "Returns a unique asymmetric data key pair for use outside of KMS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_generateDataKeyPairWithoutPlaintextCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_generateDataKeyPairWithoutPlaintextCmd).Standalone()

		kms_generateDataKeyPairWithoutPlaintextCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
		kms_generateDataKeyPairWithoutPlaintextCmd.Flags().String("encryption-context", "", "Specifies the encryption context that will be used when encrypting the private key in the data key pair.")
		kms_generateDataKeyPairWithoutPlaintextCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
		kms_generateDataKeyPairWithoutPlaintextCmd.Flags().String("key-id", "", "Specifies the symmetric encryption KMS key that encrypts the private key in the data key pair.")
		kms_generateDataKeyPairWithoutPlaintextCmd.Flags().String("key-pair-spec", "", "Determines the type of data key pair that is generated.")
		kms_generateDataKeyPairWithoutPlaintextCmd.MarkFlagRequired("key-id")
		kms_generateDataKeyPairWithoutPlaintextCmd.MarkFlagRequired("key-pair-spec")
	})
	kmsCmd.AddCommand(kms_generateDataKeyPairWithoutPlaintextCmd)
}
