package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listRulesCmd = &cobra.Command{
	Use:   "list-rules",
	Short: "Lists existing rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listRulesCmd).Standalone()

	datazone_listRulesCmd.Flags().String("action", "", "The action of the rule.")
	datazone_listRulesCmd.Flags().String("asset-types", "", "The asset types of the rule.")
	datazone_listRulesCmd.Flags().Bool("data-product", false, "The data product of the rule.")
	datazone_listRulesCmd.Flags().String("domain-identifier", "", "The ID of the domain in which the rules are to be listed.")
	datazone_listRulesCmd.Flags().Bool("include-cascaded", false, "Specifies whether to include cascading rules in the results.")
	datazone_listRulesCmd.Flags().String("max-results", "", "The maximum number of rules to return in a single call to `ListRules`.")
	datazone_listRulesCmd.Flags().String("next-token", "", "When the number of rules is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of rules, the response includes a pagination token named `NextToken`.")
	datazone_listRulesCmd.Flags().Bool("no-data-product", false, "The data product of the rule.")
	datazone_listRulesCmd.Flags().Bool("no-include-cascaded", false, "Specifies whether to include cascading rules in the results.")
	datazone_listRulesCmd.Flags().String("project-ids", "", "The IDs of projects in which rules are to be listed.")
	datazone_listRulesCmd.Flags().String("rule-type", "", "The type of the rule.")
	datazone_listRulesCmd.Flags().String("target-identifier", "", "The target ID of the rule.")
	datazone_listRulesCmd.Flags().String("target-type", "", "The target type of the rule.")
	datazone_listRulesCmd.MarkFlagRequired("domain-identifier")
	datazone_listRulesCmd.Flag("no-data-product").Hidden = true
	datazone_listRulesCmd.Flag("no-include-cascaded").Hidden = true
	datazone_listRulesCmd.MarkFlagRequired("target-identifier")
	datazone_listRulesCmd.MarkFlagRequired("target-type")
	datazoneCmd.AddCommand(datazone_listRulesCmd)
}
