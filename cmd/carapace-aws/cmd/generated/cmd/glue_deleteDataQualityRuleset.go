package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteDataQualityRulesetCmd = &cobra.Command{
	Use:   "delete-data-quality-ruleset",
	Short: "Deletes a data quality ruleset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteDataQualityRulesetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteDataQualityRulesetCmd).Standalone()

		glue_deleteDataQualityRulesetCmd.Flags().String("name", "", "A name for the data quality ruleset.")
		glue_deleteDataQualityRulesetCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_deleteDataQualityRulesetCmd)
}
