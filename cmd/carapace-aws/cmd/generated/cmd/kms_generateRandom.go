package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_generateRandomCmd = &cobra.Command{
	Use:   "generate-random",
	Short: "Returns a random byte string that is cryptographically secure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_generateRandomCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_generateRandomCmd).Standalone()

		kms_generateRandomCmd.Flags().String("custom-key-store-id", "", "Generates the random byte string in the CloudHSM cluster that is associated with the specified CloudHSM key store.")
		kms_generateRandomCmd.Flags().String("number-of-bytes", "", "The length of the random byte string.")
		kms_generateRandomCmd.Flags().String("recipient", "", "A signed [attestation document](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitro-enclave-how.html#term-attestdoc) from an Amazon Web Services Nitro enclave or NitroTPM, and the encryption algorithm to use with the public key in the attestation document.")
	})
	kmsCmd.AddCommand(kms_generateRandomCmd)
}
