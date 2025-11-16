package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listUsageForLicenseConfigurationCmd = &cobra.Command{
	Use:   "list-usage-for-license-configuration",
	Short: "Lists all license usage records for a license configuration, displaying license consumption details by resource at a selected point in time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listUsageForLicenseConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_listUsageForLicenseConfigurationCmd).Standalone()

		licenseManager_listUsageForLicenseConfigurationCmd.Flags().String("filters", "", "Filters to scope the results.")
		licenseManager_listUsageForLicenseConfigurationCmd.Flags().String("license-configuration-arn", "", "Amazon Resource Name (ARN) of the license configuration.")
		licenseManager_listUsageForLicenseConfigurationCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
		licenseManager_listUsageForLicenseConfigurationCmd.Flags().String("next-token", "", "Token for the next set of results.")
		licenseManager_listUsageForLicenseConfigurationCmd.MarkFlagRequired("license-configuration-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_listUsageForLicenseConfigurationCmd)
}
