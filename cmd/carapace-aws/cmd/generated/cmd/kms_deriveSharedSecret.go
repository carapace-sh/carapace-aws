package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_deriveSharedSecretCmd = &cobra.Command{
	Use:   "derive-shared-secret",
	Short: "Derives a shared secret using a key agreement algorithm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_deriveSharedSecretCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_deriveSharedSecretCmd).Standalone()

		kms_deriveSharedSecretCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
		kms_deriveSharedSecretCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
		kms_deriveSharedSecretCmd.Flags().String("key-agreement-algorithm", "", "Specifies the key agreement algorithm used to derive the shared secret.")
		kms_deriveSharedSecretCmd.Flags().String("key-id", "", "Identifies an asymmetric NIST-standard ECC or SM2 (China Regions only) KMS key.")
		kms_deriveSharedSecretCmd.Flags().String("public-key", "", "Specifies the public key in your peer's NIST-standard elliptic curve (ECC) or SM2 (China Regions only) key pair.")
		kms_deriveSharedSecretCmd.Flags().String("recipient", "", "A signed [attestation document](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave-how.html#term-attestdoc) from an Amazon Web Services Nitro enclave or NitroTPM, and the encryption algorithm to use with the public key in the attestation document.")
		kms_deriveSharedSecretCmd.MarkFlagRequired("key-agreement-algorithm")
		kms_deriveSharedSecretCmd.MarkFlagRequired("key-id")
		kms_deriveSharedSecretCmd.MarkFlagRequired("public-key")
	})
	kmsCmd.AddCommand(kms_deriveSharedSecretCmd)
}
