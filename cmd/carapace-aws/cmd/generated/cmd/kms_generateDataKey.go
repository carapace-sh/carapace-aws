package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_generateDataKeyCmd = &cobra.Command{
	Use:   "generate-data-key",
	Short: "Returns a unique symmetric data key for use outside of KMS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_generateDataKeyCmd).Standalone()

	kms_generateDataKeyCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
	kms_generateDataKeyCmd.Flags().String("encryption-context", "", "Specifies the encryption context that will be used when encrypting the data key.")
	kms_generateDataKeyCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
	kms_generateDataKeyCmd.Flags().String("key-id", "", "Specifies the symmetric encryption KMS key that encrypts the data key.")
	kms_generateDataKeyCmd.Flags().String("key-spec", "", "Specifies the length of the data key.")
	kms_generateDataKeyCmd.Flags().String("number-of-bytes", "", "Specifies the length of the data key in bytes.")
	kms_generateDataKeyCmd.Flags().String("recipient", "", "A signed [attestation document](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave-how.html#term-attestdoc) from an Amazon Web Services Nitro enclave or NitroTPM, and the encryption algorithm to use with the public key in the attestation document.")
	kms_generateDataKeyCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_generateDataKeyCmd)
}
