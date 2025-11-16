package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listCisScansCmd = &cobra.Command{
	Use:   "list-cis-scans",
	Short: "Returns a CIS scan list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listCisScansCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listCisScansCmd).Standalone()

		inspector2_listCisScansCmd.Flags().String("detail-level", "", "The detail applied to the CIS scan.")
		inspector2_listCisScansCmd.Flags().String("filter-criteria", "", "The CIS scan filter criteria.")
		inspector2_listCisScansCmd.Flags().String("max-results", "", "The maximum number of results to be returned.")
		inspector2_listCisScansCmd.Flags().String("next-token", "", "The pagination token from a previous request that's used to retrieve the next page of results.")
		inspector2_listCisScansCmd.Flags().String("sort-by", "", "The CIS scans sort by order.")
		inspector2_listCisScansCmd.Flags().String("sort-order", "", "The CIS scans sort order.")
	})
	inspector2Cmd.AddCommand(inspector2_listCisScansCmd)
}
