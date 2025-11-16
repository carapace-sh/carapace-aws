package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listLicenseSpecificationsForResourceCmd = &cobra.Command{
	Use:   "list-license-specifications-for-resource",
	Short: "Describes the license configurations for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listLicenseSpecificationsForResourceCmd).Standalone()

	licenseManager_listLicenseSpecificationsForResourceCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
	licenseManager_listLicenseSpecificationsForResourceCmd.Flags().String("next-token", "", "Token for the next set of results.")
	licenseManager_listLicenseSpecificationsForResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of a resource that has an associated license configuration.")
	licenseManager_listLicenseSpecificationsForResourceCmd.MarkFlagRequired("resource-arn")
	licenseManagerCmd.AddCommand(licenseManager_listLicenseSpecificationsForResourceCmd)
}
