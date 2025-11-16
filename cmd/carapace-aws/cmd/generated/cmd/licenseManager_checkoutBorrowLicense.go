package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_checkoutBorrowLicenseCmd = &cobra.Command{
	Use:   "checkout-borrow-license",
	Short: "Checks out the specified license for offline use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_checkoutBorrowLicenseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_checkoutBorrowLicenseCmd).Standalone()

		licenseManager_checkoutBorrowLicenseCmd.Flags().String("checkout-metadata", "", "Information about constraints.")
		licenseManager_checkoutBorrowLicenseCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		licenseManager_checkoutBorrowLicenseCmd.Flags().String("digital-signature-method", "", "Digital signature method.")
		licenseManager_checkoutBorrowLicenseCmd.Flags().String("entitlements", "", "License entitlements.")
		licenseManager_checkoutBorrowLicenseCmd.Flags().String("license-arn", "", "Amazon Resource Name (ARN) of the license.")
		licenseManager_checkoutBorrowLicenseCmd.Flags().String("node-id", "", "Node ID.")
		licenseManager_checkoutBorrowLicenseCmd.MarkFlagRequired("client-token")
		licenseManager_checkoutBorrowLicenseCmd.MarkFlagRequired("digital-signature-method")
		licenseManager_checkoutBorrowLicenseCmd.MarkFlagRequired("entitlements")
		licenseManager_checkoutBorrowLicenseCmd.MarkFlagRequired("license-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_checkoutBorrowLicenseCmd)
}
