package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_resetEncryptionKeyCmd = &cobra.Command{
	Use:   "reset-encryption-key",
	Short: "Resets an encryption key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_resetEncryptionKeyCmd).Standalone()

	inspector2_resetEncryptionKeyCmd.Flags().String("resource-type", "", "The resource type the key encrypts.")
	inspector2_resetEncryptionKeyCmd.Flags().String("scan-type", "", "The scan type the key encrypts.")
	inspector2_resetEncryptionKeyCmd.MarkFlagRequired("resource-type")
	inspector2_resetEncryptionKeyCmd.MarkFlagRequired("scan-type")
	inspector2Cmd.AddCommand(inspector2_resetEncryptionKeyCmd)
}
