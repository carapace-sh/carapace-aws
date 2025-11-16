package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateRuleCmd = &cobra.Command{
	Use:   "update-rule",
	Short: "Updates a rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateRuleCmd).Standalone()

	datazone_updateRuleCmd.Flags().String("description", "", "The description of the rule.")
	datazone_updateRuleCmd.Flags().String("detail", "", "The detail of the rule.")
	datazone_updateRuleCmd.Flags().String("domain-identifier", "", "The ID of the domain in which a rule is to be updated.")
	datazone_updateRuleCmd.Flags().String("identifier", "", "The ID of the rule that is to be updated")
	datazone_updateRuleCmd.Flags().Bool("include-child-domain-units", false, "Specifies whether to update this rule in the child domain units.")
	datazone_updateRuleCmd.Flags().String("name", "", "The name of the rule.")
	datazone_updateRuleCmd.Flags().Bool("no-include-child-domain-units", false, "Specifies whether to update this rule in the child domain units.")
	datazone_updateRuleCmd.Flags().String("scope", "", "The scrope of the rule.")
	datazone_updateRuleCmd.MarkFlagRequired("domain-identifier")
	datazone_updateRuleCmd.MarkFlagRequired("identifier")
	datazone_updateRuleCmd.Flag("no-include-child-domain-units").Hidden = true
	datazoneCmd.AddCommand(datazone_updateRuleCmd)
}
