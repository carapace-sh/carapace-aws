package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listDataQualityRulesetsCmd = &cobra.Command{
	Use:   "list-data-quality-rulesets",
	Short: "Returns a paginated list of rulesets for the specified list of Glue tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listDataQualityRulesetsCmd).Standalone()

	glue_listDataQualityRulesetsCmd.Flags().String("filter", "", "The filter criteria.")
	glue_listDataQualityRulesetsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	glue_listDataQualityRulesetsCmd.Flags().String("next-token", "", "A paginated token to offset the results.")
	glue_listDataQualityRulesetsCmd.Flags().String("tags", "", "A list of key-value pair tags.")
	glueCmd.AddCommand(glue_listDataQualityRulesetsCmd)
}
