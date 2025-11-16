package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listLicenseConversionTasksCmd = &cobra.Command{
	Use:   "list-license-conversion-tasks",
	Short: "Lists the license type conversion tasks for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listLicenseConversionTasksCmd).Standalone()

	licenseManager_listLicenseConversionTasksCmd.Flags().String("filters", "", "Filters to scope the results.")
	licenseManager_listLicenseConversionTasksCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
	licenseManager_listLicenseConversionTasksCmd.Flags().String("next-token", "", "Token for the next set of results.")
	licenseManagerCmd.AddCommand(licenseManager_listLicenseConversionTasksCmd)
}
