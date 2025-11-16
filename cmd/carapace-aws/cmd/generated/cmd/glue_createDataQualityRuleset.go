package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createDataQualityRulesetCmd = &cobra.Command{
	Use:   "create-data-quality-ruleset",
	Short: "Creates a data quality ruleset with DQDL rules applied to a specified Glue table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createDataQualityRulesetCmd).Standalone()

	glue_createDataQualityRulesetCmd.Flags().String("client-token", "", "Used for idempotency and is recommended to be set to a random ID (such as a UUID) to avoid creating or starting multiple instances of the same resource.")
	glue_createDataQualityRulesetCmd.Flags().String("data-quality-security-configuration", "", "The name of the security configuration created with the data quality encryption option.")
	glue_createDataQualityRulesetCmd.Flags().String("description", "", "A description of the data quality ruleset.")
	glue_createDataQualityRulesetCmd.Flags().String("name", "", "A unique name for the data quality ruleset.")
	glue_createDataQualityRulesetCmd.Flags().String("ruleset", "", "A Data Quality Definition Language (DQDL) ruleset.")
	glue_createDataQualityRulesetCmd.Flags().String("tags", "", "A list of tags applied to the data quality ruleset.")
	glue_createDataQualityRulesetCmd.Flags().String("target-table", "", "A target table associated with the data quality ruleset.")
	glue_createDataQualityRulesetCmd.MarkFlagRequired("name")
	glue_createDataQualityRulesetCmd.MarkFlagRequired("ruleset")
	glueCmd.AddCommand(glue_createDataQualityRulesetCmd)
}
