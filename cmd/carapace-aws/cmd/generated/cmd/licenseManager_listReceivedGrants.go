package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listReceivedGrantsCmd = &cobra.Command{
	Use:   "list-received-grants",
	Short: "Lists grants that are received.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listReceivedGrantsCmd).Standalone()

	licenseManager_listReceivedGrantsCmd.Flags().String("filters", "", "Filters to scope the results.")
	licenseManager_listReceivedGrantsCmd.Flags().String("grant-arns", "", "Amazon Resource Names (ARNs) of the grants.")
	licenseManager_listReceivedGrantsCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
	licenseManager_listReceivedGrantsCmd.Flags().String("next-token", "", "Token for the next set of results.")
	licenseManagerCmd.AddCommand(licenseManager_listReceivedGrantsCmd)
}
