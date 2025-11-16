package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getDataQualityRulesetCmd = &cobra.Command{
	Use:   "get-data-quality-ruleset",
	Short: "Returns an existing ruleset by identifier or name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getDataQualityRulesetCmd).Standalone()

	glue_getDataQualityRulesetCmd.Flags().String("name", "", "The name of the ruleset.")
	glue_getDataQualityRulesetCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_getDataQualityRulesetCmd)
}
