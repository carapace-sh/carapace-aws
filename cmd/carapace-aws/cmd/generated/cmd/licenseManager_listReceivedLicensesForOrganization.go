package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listReceivedLicensesForOrganizationCmd = &cobra.Command{
	Use:   "list-received-licenses-for-organization",
	Short: "Lists the licenses received for all accounts in the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listReceivedLicensesForOrganizationCmd).Standalone()

	licenseManager_listReceivedLicensesForOrganizationCmd.Flags().String("filters", "", "Filters to scope the results.")
	licenseManager_listReceivedLicensesForOrganizationCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
	licenseManager_listReceivedLicensesForOrganizationCmd.Flags().String("next-token", "", "Token for the next set of results.")
	licenseManagerCmd.AddCommand(licenseManager_listReceivedLicensesForOrganizationCmd)
}
