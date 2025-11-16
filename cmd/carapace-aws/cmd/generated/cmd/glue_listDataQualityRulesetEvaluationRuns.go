package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listDataQualityRulesetEvaluationRunsCmd = &cobra.Command{
	Use:   "list-data-quality-ruleset-evaluation-runs",
	Short: "Lists all the runs meeting the filter criteria, where a ruleset is evaluated against a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listDataQualityRulesetEvaluationRunsCmd).Standalone()

	glue_listDataQualityRulesetEvaluationRunsCmd.Flags().String("filter", "", "The filter criteria.")
	glue_listDataQualityRulesetEvaluationRunsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	glue_listDataQualityRulesetEvaluationRunsCmd.Flags().String("next-token", "", "A paginated token to offset the results.")
	glueCmd.AddCommand(glue_listDataQualityRulesetEvaluationRunsCmd)
}
