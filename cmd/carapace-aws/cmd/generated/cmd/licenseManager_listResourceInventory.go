package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_listResourceInventoryCmd = &cobra.Command{
	Use:   "list-resource-inventory",
	Short: "Lists resources managed using Systems Manager inventory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_listResourceInventoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_listResourceInventoryCmd).Standalone()

		licenseManager_listResourceInventoryCmd.Flags().String("filters", "", "Filters to scope the results.")
		licenseManager_listResourceInventoryCmd.Flags().String("max-results", "", "Maximum number of results to return in a single call.")
		licenseManager_listResourceInventoryCmd.Flags().String("next-token", "", "Token for the next set of results.")
	})
	licenseManagerCmd.AddCommand(licenseManager_listResourceInventoryCmd)
}
