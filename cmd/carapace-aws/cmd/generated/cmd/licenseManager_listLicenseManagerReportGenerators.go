package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listLicenseManagerReportGeneratorsCmd = &cobra.Command{
	Use:   "list-license-manager-report-generators",
	Short: "Lists the report generators for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listLicenseManagerReportGeneratorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_listLicenseManagerReportGeneratorsCmd).Standalone()

		licenseManager_listLicenseManagerReportGeneratorsCmd.Flags().String("filters", "", "Filters to scope the results.")
		licenseManager_listLicenseManagerReportGeneratorsCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
		licenseManager_listLicenseManagerReportGeneratorsCmd.Flags().String("next-token", "", "Token for the next set of results.")
	})
	licenseManagerCmd.AddCommand(licenseManager_listLicenseManagerReportGeneratorsCmd)
}
