package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_updateEncryptionKeyCmd = &cobra.Command{
	Use:   "update-encryption-key",
	Short: "Updates an encryption key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_updateEncryptionKeyCmd).Standalone()

	inspector2_updateEncryptionKeyCmd.Flags().String("kms-key-id", "", "A KMS key ID for the encryption key.")
	inspector2_updateEncryptionKeyCmd.Flags().String("resource-type", "", "The resource type for the encryption key.")
	inspector2_updateEncryptionKeyCmd.Flags().String("scan-type", "", "The scan type for the encryption key.")
	inspector2_updateEncryptionKeyCmd.MarkFlagRequired("kms-key-id")
	inspector2_updateEncryptionKeyCmd.MarkFlagRequired("resource-type")
	inspector2_updateEncryptionKeyCmd.MarkFlagRequired("scan-type")
	inspector2Cmd.AddCommand(inspector2_updateEncryptionKeyCmd)
}
