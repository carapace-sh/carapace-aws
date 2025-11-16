package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_enableKeyRotationCmd = &cobra.Command{
	Use:   "enable-key-rotation",
	Short: "Enables [automatic rotation of the key material](https://docs.aws.amazon.com/kms/latest/developerguide/rotating-keys-enable-disable.html) of the specified symmetric encryption KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_enableKeyRotationCmd).Standalone()

	kms_enableKeyRotationCmd.Flags().String("key-id", "", "Identifies a symmetric encryption KMS key.")
	kms_enableKeyRotationCmd.Flags().String("rotation-period-in-days", "", "Use this parameter to specify a custom period of time between each rotation date.")
	kms_enableKeyRotationCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_enableKeyRotationCmd)
}
