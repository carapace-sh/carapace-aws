package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_getLicenseUsageCmd = &cobra.Command{
	Use:   "get-license-usage",
	Short: "Gets detailed information about the usage of the specified license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_getLicenseUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_getLicenseUsageCmd).Standalone()

		licenseManager_getLicenseUsageCmd.Flags().String("license-arn", "", "Amazon Resource Name (ARN) of the license.")
		licenseManager_getLicenseUsageCmd.MarkFlagRequired("license-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_getLicenseUsageCmd)
}
