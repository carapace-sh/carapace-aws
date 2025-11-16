package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_generateDataKeyPairCmd = &cobra.Command{
	Use:   "generate-data-key-pair",
	Short: "Returns a unique asymmetric data key pair for use outside of KMS.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_generateDataKeyPairCmd).Standalone()

	kms_generateDataKeyPairCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
	kms_generateDataKeyPairCmd.Flags().String("encryption-context", "", "Specifies the encryption context that will be used when encrypting the private key in the data key pair.")
	kms_generateDataKeyPairCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
	kms_generateDataKeyPairCmd.Flags().String("key-id", "", "Specifies the symmetric encryption KMS key that encrypts the private key in the data key pair.")
	kms_generateDataKeyPairCmd.Flags().String("key-pair-spec", "", "Determines the type of data key pair that is generated.")
	kms_generateDataKeyPairCmd.Flags().String("recipient", "", "A signed [attestation document](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave-how.html#term-attestdoc) from an Amazon Web Services Nitro enclave or NitroTPM, and the encryption algorithm to use with the public key in the attestation document.")
	kms_generateDataKeyPairCmd.MarkFlagRequired("key-id")
	kms_generateDataKeyPairCmd.MarkFlagRequired("key-pair-spec")
	kmsCmd.AddCommand(kms_generateDataKeyPairCmd)
}
