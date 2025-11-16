package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDataQualityRulesetEvaluationRunCmd = &cobra.Command{
	Use:   "get-data-quality-ruleset-evaluation-run",
	Short: "Retrieves a specific run where a ruleset is evaluated against a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDataQualityRulesetEvaluationRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getDataQualityRulesetEvaluationRunCmd).Standalone()

		glue_getDataQualityRulesetEvaluationRunCmd.Flags().String("run-id", "", "The unique run identifier associated with this run.")
		glue_getDataQualityRulesetEvaluationRunCmd.MarkFlagRequired("run-id")
	})
	glueCmd.AddCommand(glue_getDataQualityRulesetEvaluationRunCmd)
}
