package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_deleteTokenCmd = &cobra.Command{
	Use:   "delete-token",
	Short: "Deletes the specified token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_deleteTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_deleteTokenCmd).Standalone()

		licenseManager_deleteTokenCmd.Flags().String("token-id", "", "Token ID.")
		licenseManager_deleteTokenCmd.MarkFlagRequired("token-id")
	})
	licenseManagerCmd.AddCommand(licenseManager_deleteTokenCmd)
}
