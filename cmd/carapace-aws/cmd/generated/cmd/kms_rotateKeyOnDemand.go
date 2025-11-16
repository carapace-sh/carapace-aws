package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_rotateKeyOnDemandCmd = &cobra.Command{
	Use:   "rotate-key-on-demand",
	Short: "Immediately initiates rotation of the key material of the specified symmetric encryption KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_rotateKeyOnDemandCmd).Standalone()

	kms_rotateKeyOnDemandCmd.Flags().String("key-id", "", "Identifies a symmetric encryption KMS key.")
	kms_rotateKeyOnDemandCmd.MarkFlagRequired("key-id")
	kmsCmd.AddCommand(kms_rotateKeyOnDemandCmd)
}
