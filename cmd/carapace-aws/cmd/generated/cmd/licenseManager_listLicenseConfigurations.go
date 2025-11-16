package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listLicenseConfigurationsCmd = &cobra.Command{
	Use:   "list-license-configurations",
	Short: "Lists the license configurations for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listLicenseConfigurationsCmd).Standalone()

	licenseManager_listLicenseConfigurationsCmd.Flags().String("filters", "", "Filters to scope the results.")
	licenseManager_listLicenseConfigurationsCmd.Flags().String("license-configuration-arns", "", "Amazon Resource Names (ARN) of the license configurations.")
	licenseManager_listLicenseConfigurationsCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
	licenseManager_listLicenseConfigurationsCmd.Flags().String("next-token", "", "Token for the next set of results.")
	licenseManagerCmd.AddCommand(licenseManager_listLicenseConfigurationsCmd)
}
