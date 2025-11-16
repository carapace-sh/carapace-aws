package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_getLicenseCmd = &cobra.Command{
	Use:   "get-license",
	Short: "Gets detailed information about the specified license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_getLicenseCmd).Standalone()

	licenseManager_getLicenseCmd.Flags().String("license-arn", "", "Amazon Resource Name (ARN) of the license.")
	licenseManager_getLicenseCmd.Flags().String("version", "", "License version.")
	licenseManager_getLicenseCmd.MarkFlagRequired("license-arn")
	licenseManagerCmd.AddCommand(licenseManager_getLicenseCmd)
}
