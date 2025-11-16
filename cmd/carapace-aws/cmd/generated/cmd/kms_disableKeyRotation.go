package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_disableKeyRotationCmd = &cobra.Command{
	Use:   "disable-key-rotation",
	Short: "Disables [automatic rotation of the key material](https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable-disable.html) of the specified symmetric encryption KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_disableKeyRotationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kms_disableKeyRotationCmd).Standalone()

		kms_disableKeyRotationCmd.Flags().String("key-id", "", "Identifies a symmetric encryption KMS key.")
		kms_disableKeyRotationCmd.MarkFlagRequired("key-id")
	})
	kmsCmd.AddCommand(kms_disableKeyRotationCmd)
}
