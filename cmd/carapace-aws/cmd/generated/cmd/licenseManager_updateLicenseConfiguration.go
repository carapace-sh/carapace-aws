package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_updateLicenseConfigurationCmd = &cobra.Command{
	Use:   "update-license-configuration",
	Short: "Modifies the attributes of an existing license configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_updateLicenseConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_updateLicenseConfigurationCmd).Standalone()

		licenseManager_updateLicenseConfigurationCmd.Flags().String("description", "", "New description of the license configuration.")
		licenseManager_updateLicenseConfigurationCmd.Flags().String("disassociate-when-not-found", "", "When true, disassociates a resource when software is uninstalled.")
		licenseManager_updateLicenseConfigurationCmd.Flags().String("license-configuration-arn", "", "Amazon Resource Name (ARN) of the license configuration.")
		licenseManager_updateLicenseConfigurationCmd.Flags().String("license-configuration-status", "", "New status of the license configuration.")
		licenseManager_updateLicenseConfigurationCmd.Flags().String("license-count", "", "New number of licenses managed by the license configuration.")
		licenseManager_updateLicenseConfigurationCmd.Flags().String("license-count-hard-limit", "", "New hard limit of the number of available licenses.")
		licenseManager_updateLicenseConfigurationCmd.Flags().String("license-rules", "", "New license rule.")
		licenseManager_updateLicenseConfigurationCmd.Flags().String("name", "", "New name of the license configuration.")
		licenseManager_updateLicenseConfigurationCmd.Flags().String("product-information-list", "", "New product information.")
		licenseManager_updateLicenseConfigurationCmd.MarkFlagRequired("license-configuration-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_updateLicenseConfigurationCmd)
}
