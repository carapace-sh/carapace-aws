package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startDataQualityRuleRecommendationRunCmd = &cobra.Command{
	Use:   "start-data-quality-rule-recommendation-run",
	Short: "Starts a recommendation run that is used to generate rules when you don't know what rules to write.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startDataQualityRuleRecommendationRunCmd).Standalone()

	glue_startDataQualityRuleRecommendationRunCmd.Flags().String("client-token", "", "Used for idempotency and is recommended to be set to a random ID (such as a UUID) to avoid creating or starting multiple instances of the same resource.")
	glue_startDataQualityRuleRecommendationRunCmd.Flags().String("created-ruleset-name", "", "A name for the ruleset.")
	glue_startDataQualityRuleRecommendationRunCmd.Flags().String("data-quality-security-configuration", "", "The name of the security configuration created with the data quality encryption option.")
	glue_startDataQualityRuleRecommendationRunCmd.Flags().String("data-source", "", "The data source (Glue table) associated with this run.")
	glue_startDataQualityRuleRecommendationRunCmd.Flags().String("number-of-workers", "", "The number of `G.1X` workers to be used in the run.")
	glue_startDataQualityRuleRecommendationRunCmd.Flags().String("role", "", "An IAM role supplied to encrypt the results of the run.")
	glue_startDataQualityRuleRecommendationRunCmd.Flags().String("timeout", "", "The timeout for a run in minutes.")
	glue_startDataQualityRuleRecommendationRunCmd.MarkFlagRequired("data-source")
	glue_startDataQualityRuleRecommendationRunCmd.MarkFlagRequired("role")
	glueCmd.AddCommand(glue_startDataQualityRuleRecommendationRunCmd)
}
