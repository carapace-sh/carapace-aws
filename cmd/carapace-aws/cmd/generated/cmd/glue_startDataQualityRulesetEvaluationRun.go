package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startDataQualityRulesetEvaluationRunCmd = &cobra.Command{
	Use:   "start-data-quality-ruleset-evaluation-run",
	Short: "Once you have a ruleset definition (either recommended or your own), you call this operation to evaluate the ruleset against a data source (Glue table).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startDataQualityRulesetEvaluationRunCmd).Standalone()

	glue_startDataQualityRulesetEvaluationRunCmd.Flags().String("additional-data-sources", "", "A map of reference strings to additional data sources you can specify for an evaluation run.")
	glue_startDataQualityRulesetEvaluationRunCmd.Flags().String("additional-run-options", "", "Additional run options you can specify for an evaluation run.")
	glue_startDataQualityRulesetEvaluationRunCmd.Flags().String("client-token", "", "Used for idempotency and is recommended to be set to a random ID (such as a UUID) to avoid creating or starting multiple instances of the same resource.")
	glue_startDataQualityRulesetEvaluationRunCmd.Flags().String("data-source", "", "The data source (Glue table) associated with this run.")
	glue_startDataQualityRulesetEvaluationRunCmd.Flags().String("number-of-workers", "", "The number of `G.1X` workers to be used in the run.")
	glue_startDataQualityRulesetEvaluationRunCmd.Flags().String("role", "", "An IAM role supplied to encrypt the results of the run.")
	glue_startDataQualityRulesetEvaluationRunCmd.Flags().String("ruleset-names", "", "A list of ruleset names.")
	glue_startDataQualityRulesetEvaluationRunCmd.Flags().String("timeout", "", "The timeout for a run in minutes.")
	glue_startDataQualityRulesetEvaluationRunCmd.MarkFlagRequired("data-source")
	glue_startDataQualityRulesetEvaluationRunCmd.MarkFlagRequired("role")
	glue_startDataQualityRulesetEvaluationRunCmd.MarkFlagRequired("ruleset-names")
	glueCmd.AddCommand(glue_startDataQualityRulesetEvaluationRunCmd)
}
