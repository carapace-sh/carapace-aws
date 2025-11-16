package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_getLicenseConfigurationCmd = &cobra.Command{
	Use:   "get-license-configuration",
	Short: "Gets detailed information about the specified license configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_getLicenseConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_getLicenseConfigurationCmd).Standalone()

		licenseManager_getLicenseConfigurationCmd.Flags().String("license-configuration-arn", "", "Amazon Resource Name (ARN) of the license configuration.")
		licenseManager_getLicenseConfigurationCmd.MarkFlagRequired("license-configuration-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_getLicenseConfigurationCmd)
}
