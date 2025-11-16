package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_updateRulesetCmd = &cobra.Command{
	Use:   "update-ruleset",
	Short: "Updates specified ruleset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_updateRulesetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_updateRulesetCmd).Standalone()

		databrew_updateRulesetCmd.Flags().String("description", "", "The description of the ruleset.")
		databrew_updateRulesetCmd.Flags().String("name", "", "The name of the ruleset to be updated.")
		databrew_updateRulesetCmd.Flags().String("rules", "", "A list of rules that are defined with the ruleset.")
		databrew_updateRulesetCmd.MarkFlagRequired("name")
		databrew_updateRulesetCmd.MarkFlagRequired("rules")
	})
	databrewCmd.AddCommand(databrew_updateRulesetCmd)
}
