package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listReceivedGrantsForOrganizationCmd = &cobra.Command{
	Use:   "list-received-grants-for-organization",
	Short: "Lists the grants received for all accounts in the organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listReceivedGrantsForOrganizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_listReceivedGrantsForOrganizationCmd).Standalone()

		licenseManager_listReceivedGrantsForOrganizationCmd.Flags().String("filters", "", "Filters to scope the results.")
		licenseManager_listReceivedGrantsForOrganizationCmd.Flags().String("license-arn", "", "The Amazon Resource Name (ARN) of the received license.")
		licenseManager_listReceivedGrantsForOrganizationCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
		licenseManager_listReceivedGrantsForOrganizationCmd.Flags().String("next-token", "", "Token for the next set of results.")
		licenseManager_listReceivedGrantsForOrganizationCmd.MarkFlagRequired("license-arn")
	})
	licenseManagerCmd.AddCommand(licenseManager_listReceivedGrantsForOrganizationCmd)
}
