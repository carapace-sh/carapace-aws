package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listTokensCmd = &cobra.Command{
	Use:   "list-tokens",
	Short: "Lists your tokens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listTokensCmd).Standalone()

	licenseManager_listTokensCmd.Flags().String("filters", "", "Filters to scope the results.")
	licenseManager_listTokensCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
	licenseManager_listTokensCmd.Flags().String("next-token", "", "Token for the next set of results.")
	licenseManager_listTokensCmd.Flags().String("token-ids", "", "Token IDs.")
	licenseManagerCmd.AddCommand(licenseManager_listTokensCmd)
}
