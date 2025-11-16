package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listCoverageStatisticsCmd = &cobra.Command{
	Use:   "list-coverage-statistics",
	Short: "Lists Amazon Inspector coverage statistics for your environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listCoverageStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listCoverageStatisticsCmd).Standalone()

		inspector2_listCoverageStatisticsCmd.Flags().String("filter-criteria", "", "An object that contains details on the filters to apply to the coverage data for your environment.")
		inspector2_listCoverageStatisticsCmd.Flags().String("group-by", "", "The value to group the results by.")
		inspector2_listCoverageStatisticsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	})
	inspector2Cmd.AddCommand(inspector2_listCoverageStatisticsCmd)
}
