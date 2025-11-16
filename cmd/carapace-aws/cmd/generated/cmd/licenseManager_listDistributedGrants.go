package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listDistributedGrantsCmd = &cobra.Command{
	Use:   "list-distributed-grants",
	Short: "Lists the grants distributed for the specified license.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listDistributedGrantsCmd).Standalone()

	licenseManager_listDistributedGrantsCmd.Flags().String("filters", "", "Filters to scope the results.")
	licenseManager_listDistributedGrantsCmd.Flags().String("grant-arns", "", "Amazon Resource Names (ARNs) of the grants.")
	licenseManager_listDistributedGrantsCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
	licenseManager_listDistributedGrantsCmd.Flags().String("next-token", "", "Token for the next set of results.")
	licenseManagerCmd.AddCommand(licenseManager_listDistributedGrantsCmd)
}
