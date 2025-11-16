package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listLicensesCmd = &cobra.Command{
	Use:   "list-licenses",
	Short: "Lists the licenses for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listLicensesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_listLicensesCmd).Standalone()

		licenseManager_listLicensesCmd.Flags().String("filters", "", "Filters to scope the results.")
		licenseManager_listLicensesCmd.Flags().String("license-arns", "", "Amazon Resource Names (ARNs) of the licenses.")
		licenseManager_listLicensesCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
		licenseManager_listLicensesCmd.Flags().String("next-token", "", "Token for the next set of results.")
	})
	licenseManagerCmd.AddCommand(licenseManager_listLicensesCmd)
}
