package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_createLicenseConfigurationCmd = &cobra.Command{
	Use:   "create-license-configuration",
	Short: "Creates a license configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_createLicenseConfigurationCmd).Standalone()

	licenseManager_createLicenseConfigurationCmd.Flags().String("description", "", "Description of the license configuration.")
	licenseManager_createLicenseConfigurationCmd.Flags().String("disassociate-when-not-found", "", "When true, disassociates a resource when software is uninstalled.")
	licenseManager_createLicenseConfigurationCmd.Flags().String("license-count", "", "Number of licenses managed by the license configuration.")
	licenseManager_createLicenseConfigurationCmd.Flags().String("license-count-hard-limit", "", "Indicates whether hard or soft license enforcement is used.")
	licenseManager_createLicenseConfigurationCmd.Flags().String("license-counting-type", "", "Dimension used to track the license inventory.")
	licenseManager_createLicenseConfigurationCmd.Flags().String("license-rules", "", "License rules.")
	licenseManager_createLicenseConfigurationCmd.Flags().String("name", "", "Name of the license configuration.")
	licenseManager_createLicenseConfigurationCmd.Flags().String("product-information-list", "", "Product information.")
	licenseManager_createLicenseConfigurationCmd.Flags().String("tags", "", "Tags to add to the license configuration.")
	licenseManager_createLicenseConfigurationCmd.MarkFlagRequired("license-counting-type")
	licenseManager_createLicenseConfigurationCmd.MarkFlagRequired("name")
	licenseManagerCmd.AddCommand(licenseManager_createLicenseConfigurationCmd)
}
