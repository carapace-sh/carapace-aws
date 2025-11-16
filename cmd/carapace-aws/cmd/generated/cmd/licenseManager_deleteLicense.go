package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_deleteLicenseCmd = &cobra.Command{
	Use:   "delete-license",
	Short: "Deletes the specified license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_deleteLicenseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_deleteLicenseCmd).Standalone()

		licenseManager_deleteLicenseCmd.Flags().String("license-arn", "", "Amazon Resource Name (ARN) of the license.")
		licenseManager_deleteLicenseCmd.Flags().String("source-version", "", "Current version of the license.")
		licenseManager_deleteLicenseCmd.MarkFlagRequired("license-arn")
		licenseManager_deleteLicenseCmd.MarkFlagRequired("source-version")
	})
	licenseManagerCmd.AddCommand(licenseManager_deleteLicenseCmd)
}
