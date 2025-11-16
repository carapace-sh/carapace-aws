package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_updateLicenseSpecificationsForResourceCmd = &cobra.Command{
	Use:   "update-license-specifications-for-resource",
	Short: "Adds or removes the specified license configurations for the specified Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_updateLicenseSpecificationsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_updateLicenseSpecificationsForResourceCmd).Standalone()

		licenseManager_updateLicenseSpecificationsForResourceCmd.Flags().String("add-license-specifications", "", "ARNs of the license configurations to add.")
		licenseManager_updateLicenseSpecificationsForResourceCmd.Flags().String("remove-license-specifications", "", "ARNs of the license configurations to remove.")
		licenseManager_updateLicenseSpecificationsForResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the Amazon Web Services resource.")
		licenseManager_updateLicenseSpecificationsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_updateLicenseSpecificationsForResourceCmd)
}
