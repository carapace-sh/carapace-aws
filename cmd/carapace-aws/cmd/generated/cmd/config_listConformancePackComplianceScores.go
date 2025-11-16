package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listConformancePackComplianceScoresCmd = &cobra.Command{
	Use:   "list-conformance-pack-compliance-scores",
	Short: "Returns a list of conformance pack compliance scores.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listConformancePackComplianceScoresCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_listConformancePackComplianceScoresCmd).Standalone()

		config_listConformancePackComplianceScoresCmd.Flags().String("filters", "", "Filters the results based on the `ConformancePackComplianceScoresFilters`.")
		config_listConformancePackComplianceScoresCmd.Flags().String("limit", "", "The maximum number of conformance pack compliance scores returned on each page.")
		config_listConformancePackComplianceScoresCmd.Flags().String("next-token", "", "The `nextToken` string in a prior request that you can use to get the paginated response for the next set of conformance pack compliance scores.")
		config_listConformancePackComplianceScoresCmd.Flags().String("sort-by", "", "Sorts your conformance pack compliance scores in either ascending or descending order, depending on `SortOrder`.")
		config_listConformancePackComplianceScoresCmd.Flags().String("sort-order", "", "Determines the order in which conformance pack compliance scores are sorted.")
	})
	configCmd.AddCommand(config_listConformancePackComplianceScoresCmd)
}
