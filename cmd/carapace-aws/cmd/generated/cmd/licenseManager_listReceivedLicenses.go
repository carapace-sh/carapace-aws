package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listReceivedLicensesCmd = &cobra.Command{
	Use:   "list-received-licenses",
	Short: "Lists received licenses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listReceivedLicensesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_listReceivedLicensesCmd).Standalone()

		licenseManager_listReceivedLicensesCmd.Flags().String("filters", "", "Filters to scope the results.")
		licenseManager_listReceivedLicensesCmd.Flags().String("license-arns", "", "Amazon Resource Names (ARNs) of the licenses.")
		licenseManager_listReceivedLicensesCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
		licenseManager_listReceivedLicensesCmd.Flags().String("next-token", "", "Token for the next set of results.")
	})
	licenseManagerCmd.AddCommand(licenseManager_listReceivedLicensesCmd)
}
