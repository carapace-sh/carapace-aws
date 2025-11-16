package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_createLicenseCmd = &cobra.Command{
	Use:   "create-license",
	Short: "Creates a license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_createLicenseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_createLicenseCmd).Standalone()

		licenseManager_createLicenseCmd.Flags().String("beneficiary", "", "License beneficiary.")
		licenseManager_createLicenseCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		licenseManager_createLicenseCmd.Flags().String("consumption-configuration", "", "Configuration for consumption of the license.")
		licenseManager_createLicenseCmd.Flags().String("entitlements", "", "License entitlements.")
		licenseManager_createLicenseCmd.Flags().String("home-region", "", "Home Region for the license.")
		licenseManager_createLicenseCmd.Flags().String("issuer", "", "License issuer.")
		licenseManager_createLicenseCmd.Flags().String("license-metadata", "", "Information about the license.")
		licenseManager_createLicenseCmd.Flags().String("license-name", "", "License name.")
		licenseManager_createLicenseCmd.Flags().String("product-name", "", "Product name.")
		licenseManager_createLicenseCmd.Flags().String("product-sku", "", "Product SKU.")
		licenseManager_createLicenseCmd.Flags().String("tags", "", "Tags to add to the license.")
		licenseManager_createLicenseCmd.Flags().String("validity", "", "Date and time range during which the license is valid, in ISO8601-UTC format.")
		licenseManager_createLicenseCmd.MarkFlagRequired("beneficiary")
		licenseManager_createLicenseCmd.MarkFlagRequired("client-token")
		licenseManager_createLicenseCmd.MarkFlagRequired("consumption-configuration")
		licenseManager_createLicenseCmd.MarkFlagRequired("entitlements")
		licenseManager_createLicenseCmd.MarkFlagRequired("home-region")
		licenseManager_createLicenseCmd.MarkFlagRequired("issuer")
		licenseManager_createLicenseCmd.MarkFlagRequired("license-name")
		licenseManager_createLicenseCmd.MarkFlagRequired("product-name")
		licenseManager_createLicenseCmd.MarkFlagRequired("product-sku")
		licenseManager_createLicenseCmd.MarkFlagRequired("validity")
	})
	licenseManagerCmd.AddCommand(licenseManager_createLicenseCmd)
}
