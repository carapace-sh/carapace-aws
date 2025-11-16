package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_cancelDataQualityRulesetEvaluationRunCmd = &cobra.Command{
	Use:   "cancel-data-quality-ruleset-evaluation-run",
	Short: "Cancels a run where a ruleset is being evaluated against a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_cancelDataQualityRulesetEvaluationRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_cancelDataQualityRulesetEvaluationRunCmd).Standalone()

		glue_cancelDataQualityRulesetEvaluationRunCmd.Flags().String("run-id", "", "The unique run identifier associated with this run.")
		glue_cancelDataQualityRulesetEvaluationRunCmd.MarkFlagRequired("run-id")
	})
	glueCmd.AddCommand(glue_cancelDataQualityRulesetEvaluationRunCmd)
}
