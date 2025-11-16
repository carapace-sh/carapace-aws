package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getEncryptionConfigCmd = &cobra.Command{
	Use:   "get-encryption-config",
	Short: "Retrieves the current encryption configuration for X-Ray data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getEncryptionConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_getEncryptionConfigCmd).Standalone()

	})
	xrayCmd.AddCommand(xray_getEncryptionConfigCmd)
}
