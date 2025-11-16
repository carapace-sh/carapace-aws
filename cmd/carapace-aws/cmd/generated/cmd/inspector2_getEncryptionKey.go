package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getEncryptionKeyCmd = &cobra.Command{
	Use:   "get-encryption-key",
	Short: "Gets an encryption key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getEncryptionKeyCmd).Standalone()

	inspector2_getEncryptionKeyCmd.Flags().String("resource-type", "", "The resource type the key encrypts.")
	inspector2_getEncryptionKeyCmd.Flags().String("scan-type", "", "The scan type the key encrypts.")
	inspector2_getEncryptionKeyCmd.MarkFlagRequired("resource-type")
	inspector2_getEncryptionKeyCmd.MarkFlagRequired("scan-type")
	inspector2Cmd.AddCommand(inspector2_getEncryptionKeyCmd)
}
