package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_getKeyRotationStatusCmd = &cobra.Command{
	Use:   "get-key-rotation-status",
	Short: "Provides detailed information about the rotation status for a KMS key, including whether [automatic rotation of the key material](https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable-disable.html) is enabled for the specified KMS key, the [rotation period](https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#rotation-period), and the next scheduled rotation date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_getKeyRotationStatusCmd).Standalone()

	kms_getKeyRotationStatusCmd.Flags().String("key-id", "", "Gets the rotation status for the specified KMS key.")
	kms_getKeyRotationStatusCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_getKeyRotationStatusCmd)
}
