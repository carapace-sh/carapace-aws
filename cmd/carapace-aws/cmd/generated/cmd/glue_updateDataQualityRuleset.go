package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateDataQualityRulesetCmd = &cobra.Command{
	Use:   "update-data-quality-ruleset",
	Short: "Updates the specified data quality ruleset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateDataQualityRulesetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_updateDataQualityRulesetCmd).Standalone()

		glue_updateDataQualityRulesetCmd.Flags().String("description", "", "A description of the ruleset.")
		glue_updateDataQualityRulesetCmd.Flags().String("name", "", "The name of the data quality ruleset.")
		glue_updateDataQualityRulesetCmd.Flags().String("ruleset", "", "A Data Quality Definition Language (DQDL) ruleset.")
		glue_updateDataQualityRulesetCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_updateDataQualityRulesetCmd)
}
