package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_createLicenseVersionCmd = &cobra.Command{
	Use:   "create-license-version",
	Short: "Creates a new version of the specified license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_createLicenseVersionCmd).Standalone()

	licenseManager_createLicenseVersionCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	licenseManager_createLicenseVersionCmd.Flags().String("consumption-configuration", "", "Configuration for consumption of the license.")
	licenseManager_createLicenseVersionCmd.Flags().String("entitlements", "", "License entitlements.")
	licenseManager_createLicenseVersionCmd.Flags().String("home-region", "", "Home Region of the license.")
	licenseManager_createLicenseVersionCmd.Flags().String("issuer", "", "License issuer.")
	licenseManager_createLicenseVersionCmd.Flags().String("license-arn", "", "Amazon Resource Name (ARN) of the license.")
	licenseManager_createLicenseVersionCmd.Flags().String("license-metadata", "", "Information about the license.")
	licenseManager_createLicenseVersionCmd.Flags().String("license-name", "", "License name.")
	licenseManager_createLicenseVersionCmd.Flags().String("product-name", "", "Product name.")
	licenseManager_createLicenseVersionCmd.Flags().String("source-version", "", "Current version of the license.")
	licenseManager_createLicenseVersionCmd.Flags().String("status", "", "License status.")
	licenseManager_createLicenseVersionCmd.Flags().String("validity", "", "Date and time range during which the license is valid, in ISO8601-UTC format.")
	licenseManager_createLicenseVersionCmd.MarkFlagRequired("client-token")
	licenseManager_createLicenseVersionCmd.MarkFlagRequired("consumption-configuration")
	licenseManager_createLicenseVersionCmd.MarkFlagRequired("entitlements")
	licenseManager_createLicenseVersionCmd.MarkFlagRequired("home-region")
	licenseManager_createLicenseVersionCmd.MarkFlagRequired("issuer")
	licenseManager_createLicenseVersionCmd.MarkFlagRequired("license-arn")
	licenseManager_createLicenseVersionCmd.MarkFlagRequired("license-name")
	licenseManager_createLicenseVersionCmd.MarkFlagRequired("product-name")
	licenseManager_createLicenseVersionCmd.MarkFlagRequired("status")
	licenseManager_createLicenseVersionCmd.MarkFlagRequired("validity")
	licenseManagerCmd.AddCommand(licenseManager_createLicenseVersionCmd)
}
