package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_checkoutLicenseCmd = &cobra.Command{
	Use:   "checkout-license",
	Short: "Checks out the specified license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_checkoutLicenseCmd).Standalone()

	licenseManager_checkoutLicenseCmd.Flags().String("beneficiary", "", "License beneficiary.")
	licenseManager_checkoutLicenseCmd.Flags().String("checkout-type", "", "Checkout type.")
	licenseManager_checkoutLicenseCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	licenseManager_checkoutLicenseCmd.Flags().String("entitlements", "", "License entitlements.")
	licenseManager_checkoutLicenseCmd.Flags().String("key-fingerprint", "", "Key fingerprint identifying the license.")
	licenseManager_checkoutLicenseCmd.Flags().String("node-id", "", "Node ID.")
	licenseManager_checkoutLicenseCmd.Flags().String("product-sku", "", "Product SKU.")
	licenseManager_checkoutLicenseCmd.MarkFlagRequired("checkout-type")
	licenseManager_checkoutLicenseCmd.MarkFlagRequired("client-token")
	licenseManager_checkoutLicenseCmd.MarkFlagRequired("entitlements")
	licenseManager_checkoutLicenseCmd.MarkFlagRequired("key-fingerprint")
	licenseManager_checkoutLicenseCmd.MarkFlagRequired("product-sku")
	licenseManagerCmd.AddCommand(licenseManager_checkoutLicenseCmd)
}
