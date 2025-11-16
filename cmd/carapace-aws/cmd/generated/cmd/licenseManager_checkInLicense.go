package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_checkInLicenseCmd = &cobra.Command{
	Use:   "check-in-license",
	Short: "Checks in the specified license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_checkInLicenseCmd).Standalone()

	licenseManager_checkInLicenseCmd.Flags().String("beneficiary", "", "License beneficiary.")
	licenseManager_checkInLicenseCmd.Flags().String("license-consumption-token", "", "License consumption token.")
	licenseManager_checkInLicenseCmd.MarkFlagRequired("license-consumption-token")
	licenseManagerCmd.AddCommand(licenseManager_checkInLicenseCmd)
}
