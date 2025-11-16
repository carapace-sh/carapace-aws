package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listCisScanConfigurationsCmd = &cobra.Command{
	Use:   "list-cis-scan-configurations",
	Short: "Lists CIS scan configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listCisScanConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listCisScanConfigurationsCmd).Standalone()

		inspector2_listCisScanConfigurationsCmd.Flags().String("filter-criteria", "", "The CIS scan configuration filter criteria.")
		inspector2_listCisScanConfigurationsCmd.Flags().String("max-results", "", "The maximum number of CIS scan configurations to be returned in a single page of results.")
		inspector2_listCisScanConfigurationsCmd.Flags().String("next-token", "", "The pagination token from a previous request that's used to retrieve the next page of results.")
		inspector2_listCisScanConfigurationsCmd.Flags().String("sort-by", "", "The CIS scan configuration sort by order.")
		inspector2_listCisScanConfigurationsCmd.Flags().String("sort-order", "", "The CIS scan configuration sort order order.")
	})
	inspector2Cmd.AddCommand(inspector2_listCisScanConfigurationsCmd)
}
