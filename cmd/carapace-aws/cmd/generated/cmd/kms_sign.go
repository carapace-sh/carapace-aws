package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Creates a [digital signature](https://en.wikipedia.org/wiki/Digital_signature) for a message or message digest by using the private key in an asymmetric signing KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_signCmd).Standalone()

	kms_signCmd.Flags().String("dry-run", "", "Checks if your request will succeed.")
	kms_signCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
	kms_signCmd.Flags().String("key-id", "", "Identifies an asymmetric KMS key.")
	kms_signCmd.Flags().String("message", "", "Specifies the message or message digest to sign.")
	kms_signCmd.Flags().String("message-type", "", "Tells KMS whether the value of the `Message` parameter should be hashed as part of the signing algorithm.")
	kms_signCmd.Flags().String("signing-algorithm", "", "Specifies the signing algorithm to use when signing the message.")
	kms_signCmd.MarkFlagRequired("key-id")
	kms_signCmd.MarkFlagRequired("message")
	kms_signCmd.MarkFlagRequired("signing-algorithm")
	kmsCmd.AddCommand(kms_signCmd)
}
