package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_getPublicKeyCmd = &cobra.Command{
	Use:   "get-public-key",
	Short: "Returns the public key of an asymmetric KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_getPublicKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_getPublicKeyCmd).Standalone()

		kms_getPublicKeyCmd.Flags().String("grant-tokens", "", "A list of grant tokens.")
		kms_getPublicKeyCmd.Flags().String("key-id", "", "Identifies the asymmetric KMS key that includes the public key.")
		kms_getPublicKeyCmd.MarkFlagRequired("key-id")
	})
	kmsCmd.AddCommand(kms_getPublicKeyCmd)
}
