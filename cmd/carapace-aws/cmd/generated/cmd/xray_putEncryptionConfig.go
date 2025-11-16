package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_putEncryptionConfigCmd = &cobra.Command{
	Use:   "put-encryption-config",
	Short: "Updates the encryption configuration for X-Ray data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_putEncryptionConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_putEncryptionConfigCmd).Standalone()

		xray_putEncryptionConfigCmd.Flags().String("key-id", "", "An Amazon Web Services KMS key in one of the following formats:")
		xray_putEncryptionConfigCmd.Flags().String("type", "", "The type of encryption.")
		xray_putEncryptionConfigCmd.MarkFlagRequired("type")
	})
	xrayCmd.AddCommand(xray_putEncryptionConfigCmd)
}
